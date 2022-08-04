package keeper

import (
	"github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
