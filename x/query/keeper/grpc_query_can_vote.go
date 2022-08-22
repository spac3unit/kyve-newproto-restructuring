package keeper

import (
	"context"
	"strings"

	"github.com/KYVENetwork/chain/util"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanVote(c context.Context, req *types.QueryCanVoteRequest) (*types.QueryCanVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, req.PoolId)

	if err := k.bundleKeeper.AssertPoolCanRun(ctx, req.PoolId); err != nil {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   err.Error(),
		}, nil
	}

	// Check if sender is a staker in pool
	if err := k.stakerKeeper.AssertValaccountAuthorized(ctx, req.PoolId, req.Staker, req.Voter); err != nil {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "valaccount not authorized",
		}, nil
	}

	// Check if dropped bundle
	if bundleProposal.StorageId == "" {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "can not vote on dropped bundle",
		}, nil
	}

	// Check if empty bundle
	if strings.HasPrefix(bundleProposal.StorageId, bundletypes.KYVE_NO_DATA_BUNDLE) {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "can not vote on KYVE_NO_DATA_BUNDLE",
		}, nil
	}

	// Check if tx matches current bundleProposal
	if req.StorageId != bundleProposal.StorageId {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "provided storage_id does not match current one",
		}, nil
	}

	// Check if the sender has already voted on the bundle.
	hasVotedValid := util.ContainsString(bundleProposal.VotersValid, req.Staker)
	hasVotedInvalid := util.ContainsString(bundleProposal.VotersInvalid, req.Staker)
	hasVotedAbstain := util.ContainsString(bundleProposal.VotersAbstain, req.Staker)

	if hasVotedValid {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "has already voted valid",
		}, nil
	}

	if hasVotedInvalid {
		return &types.QueryCanVoteResponse{
			Possible: false,
			Reason:   "has already voted invalid",
		}, nil
	}

	if hasVotedAbstain {
		return &types.QueryCanVoteResponse{
			Possible: true,
			Reason:   "KYVE_VOTE_NO_ABSTAIN_ALLOWED",
		}, nil
	}

	return &types.QueryCanVoteResponse{
		Possible: true,
		Reason:   "",
	}, nil
}
