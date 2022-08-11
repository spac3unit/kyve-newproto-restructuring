package keeper

import (
	"context"
	"strings"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SubmitBundleProposal handles the logic of an SDK message that allows protocol nodes to submit a new bundle proposal.
func (k msgServer) SubmitBundleProposal(
	goCtx context.Context, msg *types.MsgSubmitBundleProposal,
) (*types.MsgSubmitBundleProposalResponse, error) {

	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check min stake+delegation
	if err := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId); err != nil {
		return nil, err
	}

	if err := k.stakerKeeper.AssertValaccountAuthorized(ctx, msg.PoolId, msg.Staker, msg.Creator); err != nil {
		return nil, err
	}

	// TODO BEGIN BUNDLE LOGIC
	pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)
	bundleProposal, found := k.GetBundleProposal(ctx, msg.PoolId)

	if !found {
		return nil, sdkErrors.ErrNotFound
	}

	// Validate submit bundle args.
	if err := k.validateSubmitBundleArgs(ctx, &bundleProposal, msg); err != nil {
		return nil, err
	}

	// reset points of uploader
	k.stakerKeeper.ResetPoints(ctx, msg.PoolId, msg.Staker)

	// If bundle was dropped or is of type KYVE_NO_DATA_BUNDLE just register new bundle.
	if bundleProposal.StorageId == "" || strings.HasPrefix(bundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		nextUploader := k.chooseNextUploaderFromAllStakers(ctx, msg.PoolId)

		if err := k.registerBundleProposalFromUploader(ctx, pool, bundleProposal, msg, nextUploader); err != nil {
			return nil, err
		}

		return &types.MsgSubmitBundleProposalResponse{}, nil
	}

	// increase points of stakers who did not vote at all
	k.handleNonVoters(ctx, msg.PoolId)

	// Get next uploader
	voters := append(bundleProposal.VotersValid, bundleProposal.VotersInvalid...)
	nextUploader := ""

	if len(voters) > 0 {
		nextUploader = k.chooseNextUploaderFromSelectedStakers(ctx, msg.PoolId, voters)
	} else {
		nextUploader = k.chooseNextUploaderFromAllStakers(ctx, msg.PoolId)
	}

	// check if the quorum was actually reached
	voteDistribution := k.getVoteDistribution(ctx, msg.PoolId)

	// handle valid proposal
	if voteDistribution.Status == types.BUNDLE_STATUS_VALID {
		// Calculate the total reward for the bundle, and individual payouts.
		bundleReward := k.calculatePayouts(ctx, msg.PoolId)

		if err := k.poolKeeper.ChargeFundersOfPool(ctx, msg.PoolId, bundleReward.Total); err != nil {
			return nil, err
		}

		pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)
		bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

		if len(pool.Funders) == 0 {
			// drop bundle because pool ran out of funds
			bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())
			k.SetBundleProposal(ctx, bundleProposal)
			// TODO: emit event

			return &types.MsgSubmitBundleProposalResponse{}, nil
		}

		// send network fee to treasury
		if err := util.TransferFromModuleToTreasury(k.accountKeeper, k.distrkeeper, ctx, pooltypes.ModuleName, bundleReward.Treasury); err != nil {
			return nil, err
		}

		// send commission to uploader
		if err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, pooltypes.ModuleName, bundleProposal.Uploader, bundleReward.Uploader); err != nil {
			return nil, err
		}

		// send delegation rewards to delegators
		// TODO: double check if delegation module receives assets
		k.delegationKeeper.AddAmountToDelegationRewards(ctx, bundleProposal.Uploader, bundleReward.Delegation)

		// slash stakers who voted incorrectly
		for _, voter := range bundleProposal.VotersInvalid {
			k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakertypes.SLASH_TYPE_VOTE)
		}

		if err := k.finalizeCurrentBundleProposal(ctx, pool, bundleProposal, voteDistribution, bundleReward); err != nil {
			return nil, err
		}

		if err := k.registerBundleProposalFromUploader(ctx, pool, bundleProposal, msg, nextUploader); err != nil {
			return nil, err
		}

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else if voteDistribution.Status == types.BUNDLE_STATUS_INVALID {
		// slash stakers who voted incorrectly - uploader receives upload slash
		for _, voter := range bundleProposal.VotersValid {
			if voter == bundleProposal.Uploader {
				k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakertypes.SLASH_TYPE_UPLOAD)
			} else {
				k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakertypes.SLASH_TYPE_VOTE)
			}
		}

		if err := k.dropCurrentBundleProposal(ctx, pool, bundleProposal, voteDistribution); err != nil {
			return nil, err
		}

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else {
		return nil, types.ErrQuorumNotReached
	}
}
