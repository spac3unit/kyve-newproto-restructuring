package keeper

import (
	"context"

	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
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
		data = append(data, k.parsePoolResponse(ctx, &pool))
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

	return &types.QueryPoolResponse{Pool: k.parsePoolResponse(ctx, &pool)}, nil
}

func (k Keeper) parsePoolResponse(ctx sdk.Context, pool *pooltypes.Pool) types.PoolResponse {
	bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, pool.Id)
	stakers := k.stakerKeeper.GetAllStakerAddressesOfPool(ctx, pool.Id)
	totalStake := k.stakerKeeper.GetTotalStake(ctx, pool.Id)

	var poolStatus pooltypes.PoolStatus

	if pool.UpgradePlan.ScheduledAt > 0 && uint64(ctx.BlockTime().Unix()) >= pool.UpgradePlan.ScheduledAt {
		poolStatus = pooltypes.POOL_STATUS_UPGRADING
	} else if pool.Paused {
		poolStatus = pooltypes.POOL_STATUS_PAUSED
	} else if totalStake < pool.MinStake {
		poolStatus = pooltypes.POOL_STATUS_NOT_ENOUGH_STAKE
	} else if pool.TotalFunds == 0 {
		poolStatus = pooltypes.POOL_STATUS_NO_FUNDS
	} else {
		poolStatus = pooltypes.POOL_STATUS_ACTIVE
	}

	return types.PoolResponse{
		Id:             pool.Id,
		Data:           pool,
		BundleProposal: &bundleProposal,
		Stakers:        stakers,
		TotalStake:     totalStake,
		Status:         poolStatus,
	}
}
