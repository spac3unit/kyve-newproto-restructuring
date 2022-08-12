package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanPropose(c context.Context, req *types.QueryCanProposeRequest) (*types.QueryCanProposeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	pool, found := k.poolKeeper.GetPool(ctx, req.PoolId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, req.PoolId)

	if err := k.bundleKeeper.AssertPoolCanRun(ctx, req.PoolId); err != nil {
		return &types.QueryCanProposeResponse{
			Possible: false,
			Reason: err.Error(),
		}, nil
	}

	// Check if sender is a staker in pool
	if err := k.stakerKeeper.AssertValaccountAuthorized(ctx, req.PoolId, req.Staker, req.Proposer); err != nil {
		return &types.QueryCanProposeResponse{
			Possible: false,
			Reason: "valaccount not authorized",
		}, nil
	}

	// Check if from_height matches
	if bundleProposal.ToHeight != req.FromHeight {
		return &types.QueryCanProposeResponse{
			Possible: false,
			Reason:   "invalid from_height",
		}, nil
	}

	// Check if designated uploader
	if bundleProposal.NextUploader != req.Proposer {
		return &types.QueryCanProposeResponse{
			Possible: false,
			Reason:   "not designated uploader",
		}, nil
	}

	// Check if upload interval has been surpassed
	if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval) {
		return &types.QueryCanProposeResponse{
			Possible: false,
			Reason:   "upload interval not surpassed",
		}, nil
	}

	return &types.QueryCanProposeResponse{
		Possible: true,
		Reason: "",
	}, nil
}
