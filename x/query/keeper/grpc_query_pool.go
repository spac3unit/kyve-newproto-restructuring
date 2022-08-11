package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	pools, pageRes, err := k.poolKeeper.GetPaginatedPoolsQuery(ctx, req.Pagination, req.Search, req.Runtime, req.Paused)
	if err != nil {
		return nil, err
	}

	data := make([]types.PoolResponse, 0)

	for _, pool := range pools {
		bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, pool.Id)
		valaccounts := k.stakerKeeper.GetAllValaccountsOfPool(ctx, pool.Id)
		totalStake := k.stakerKeeper.GetTotalStake(ctx, pool.Id)

		data = append(data, types.PoolResponse{
			Id: pool.Id,
			Pool: &pool,
			BundleProposal: &bundleProposal,
			Valaccounts: valaccounts,
			TotalStake: totalStake,
		})
	}

	return &types.QueryPoolsResponse{Pools: data, Pagination: pageRes}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.poolKeeper.GetPool(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, pool.Id)
	valaccounts := k.stakerKeeper.GetAllValaccountsOfPool(ctx, pool.Id)
	totalStake := k.stakerKeeper.GetTotalStake(ctx, pool.Id)

	return &types.QueryPoolResponse{Pool: types.PoolResponse{
		Id: pool.Id,
		Pool: &pool,
		BundleProposal: &bundleProposal,
		Valaccounts: valaccounts,
		TotalStake: totalStake,
	}}, nil
}
