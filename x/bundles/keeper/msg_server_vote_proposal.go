package keeper

import (
	"context"
	"strings"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/bundles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// VoteProposal handles the logic of an SDK message that allows protocol nodes to vote on a pool's bundle proposal.
func (k msgServer) VoteProposal(
	goCtx context.Context, msg *types.MsgVoteProposal,
) (*types.MsgVoteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check min stake+delegation
	if err := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId); err != nil {
		return nil, err
	}

	if err := k.stakerKeeper.AssertValaccountAuthorized(ctx, msg.PoolId, msg.Staker, msg.Creator, ); err != nil {
		return nil, err
	}

	bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

	// Check if the sender is also the bundle's uploader.
	if bundleProposal.Uploader == msg.Staker {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrVoterIsUploader.Error())
	}

	// Check if bundle is not dropped or NO_DATA_BUNDLE
	if bundleProposal.StorageId == "" || strings.HasPrefix(bundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		return nil, sdkErrors.Wrapf(
			sdkErrors.ErrNotFound, types.ErrInvalidStorageId.Error(), bundleProposal.StorageId,
		)
	}

	// Check if the sender is voting on the same bundle.
	if msg.StorageId != bundleProposal.StorageId {
		return nil, sdkErrors.Wrapf(
			sdkErrors.ErrNotFound, types.ErrInvalidStorageId.Error(), bundleProposal.StorageId,
		)
	}

	// Check if the sender has already voted on the bundle.
	hasVotedValid := util.ContainsString(bundleProposal.VotersValid, msg.Staker)
	hasVotedInvalid := util.ContainsString(bundleProposal.VotersInvalid, msg.Staker)
	hasVotedAbstain := util.ContainsString(bundleProposal.VotersAbstain, msg.Staker)

	if hasVotedValid || hasVotedInvalid {
		return nil, sdkErrors.Wrapf(
			sdkErrors.ErrUnauthorized, types.ErrAlreadyVoted.Error(), bundleProposal.StorageId,
		)
	}

	if hasVotedAbstain {
		if msg.Vote == types.VOTE_TYPE_ABSTAIN {
			return nil, sdkErrors.Wrapf(
				sdkErrors.ErrUnauthorized, types.ErrAlreadyVoted.Error(), bundleProposal.StorageId,
			)
		}

		// remove voter from abstain votes
		bundleProposal.VotersAbstain, _ = util.RemoveFromStringArrayStable(bundleProposal.VotersAbstain, msg.Staker)
	}

	// Update and return.
	if msg.Vote == types.VOTE_TYPE_YES {
		bundleProposal.VotersValid = append(bundleProposal.VotersValid, msg.Staker)
	} else if msg.Vote == types.VOTE_TYPE_NO {
		bundleProposal.VotersInvalid = append(bundleProposal.VotersInvalid, msg.Staker)
	} else if msg.Vote == types.VOTE_TYPE_ABSTAIN {
		bundleProposal.VotersAbstain = append(bundleProposal.VotersAbstain, msg.Staker)
	} else {
		return nil, sdkErrors.Wrapf(
			sdkErrors.ErrUnauthorized, types.ErrInvalidVote.Error(), msg.Vote,
		)
	}

	k.SetBundleProposal(ctx, bundleProposal)

	// reset points
	k.stakerKeeper.ResetPoints(ctx, msg.PoolId, msg.Staker)

	// Emit a vote event.
	if err := ctx.EventManager().EmitTypedEvent(&types.EventBundleVote{
		PoolId:    msg.PoolId,
		Staker:   msg.Staker,
		StorageId: msg.StorageId,
		Vote:      msg.Vote,
	}); err != nil {
		return nil, err
	}

	return &types.MsgVoteProposalResponse{}, nil
}
