package keeper

import (
	"context"
	"strings"

	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	"github.com/KYVENetwork/chain/x/query/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pools []types.PoolResponse
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, types.KeyPrefix(pooltypes.StoreKey))

	pageRes, err := query.FilteredPaginate(poolStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var pool pooltypes.Pool
		if err := k.cdc.Unmarshal(value, &pool); err != nil {
			return false, err
		}

		// filter search
		if !strings.Contains(strings.ToLower(pool.Name), strings.ToLower(req.Search)) {
			return false, nil
		}

		// filter runtime
		if req.Runtime != "" && req.Runtime != pool.Runtime {
			return false, nil
		}

		// filter paused
		if req.Paused != pool.Paused {
			return false, nil
		}

		if accumulate {
			bundleProposal, _ := k.bundleKeeper.GetBundleProposal(ctx, pool.Id)

			pools = append(pools, types.PoolResponse{
				Pool: &pool,
				BundleProposal: &bundleProposal,
			})
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPoolsResponse{Pools: pools, Pagination: pageRes}, nil
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

	return &types.QueryPoolResponse{Pool: types.PoolResponse{
		Pool: &pool,
		BundleProposal: &bundleProposal,
	}}, nil
}
