package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetBundleProposal set a specific staker in the store from its index
func (k Keeper) SetBundleProposal(ctx sdk.Context, bundleProposal types.BundleProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BundleKeyPrefix)
	b := k.cdc.MustMarshal(&bundleProposal)
	store.Set(types.BundleProposalKey(
		bundleProposal.PoolId,
	), b)
}

// GetBundleProposal returns a staker from its index
func (k Keeper) GetBundleProposal(
	ctx sdk.Context,
	poolId uint64,
) (val types.BundleProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BundleKeyPrefix)

	b := store.Get(types.BundleProposalKey(poolId))
	if b == nil {
		val.PoolId = poolId
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetFinalizedBundle set a specific staker in the store from its index
func (k Keeper) SetFinalizedBundle(ctx sdk.Context, finalizedBundle types.FinalizedBundle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FinalizedBundlePrefix)
	b := k.cdc.MustMarshal(&finalizedBundle)
	store.Set(types.FinalizedBundleKey(
		finalizedBundle.PoolId,
		finalizedBundle.Id,
	), b)
}

func (k Keeper) GetAllFinalizedBundles(
	ctx sdk.Context,
) (list []types.FinalizedBundle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FinalizedBundlePrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		var val types.FinalizedBundle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetFinalizedBundlesByPool(
	ctx sdk.Context,
	poolId uint64,
) (list []types.FinalizedBundle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FinalizedBundlePrefix)
	iterator := sdk.KVStorePrefixIterator(store, util.GetByteKey(poolId))

	for ; iterator.Valid(); iterator.Next() {
		var val types.FinalizedBundle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFinalizedBundle returns a staker from its index
func (k Keeper) GetFinalizedBundle(
	ctx sdk.Context,
	poolId uint64,
	id uint64,
) (val types.FinalizedBundle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FinalizedBundlePrefix)

	b := store.Get(types.FinalizedBundleKey(poolId, id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetPaginatedFinalizedBundleQuery(ctx sdk.Context, pagination *query.PageRequest, poolId uint64) ([]types.FinalizedBundle, *query.PageResponse, error) {
	var data []types.FinalizedBundle

	store := prefix.NewStore(ctx.KVStore(k.storeKey), util.GetByteKey(types.FinalizedBundlePrefix, poolId))

	pageRes, err := query.FilteredPaginate(store, pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			var finalizedBundle types.FinalizedBundle
			if err := k.cdc.Unmarshal(value, &finalizedBundle); err != nil {
				return false, err
			}

			data = append(data, finalizedBundle)
		}

		return true, nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return data, pageRes, nil
}
