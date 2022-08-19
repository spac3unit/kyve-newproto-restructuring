package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FinalizedBundles(c context.Context, req *types.QueryFinalizedBundlesRequest) (*types.QueryFinalizedBundlesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	finalizedBundles, pageRes, err := k.bundleKeeper.GetPaginatedFinalizedBundleQuery(ctx, req.Pagination, req.PoolId)
	if err != nil {
		return nil, err
	}

	return &types.QueryFinalizedBundlesResponse{FinalizedBundles: finalizedBundles, Pagination: pageRes}, nil
}

func (k Keeper) FinalizedBundle(c context.Context, req *types.QueryFinalizedBundleRequest) (*types.QueryFinalizedBundleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	finalizedBundle, found := k.bundleKeeper.GetFinalizedBundle(ctx, req.PoolId, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryFinalizedBundleResponse{FinalizedBundle: finalizedBundle}, nil
}
