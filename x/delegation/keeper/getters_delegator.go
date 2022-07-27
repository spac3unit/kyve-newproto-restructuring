package keeper

import (
	"github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDelegator set a specific delegator in the store from its index
func (k Keeper) SetDelegator(ctx sdk.Context, delegator types.Delegator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefix)
	b := k.cdc.MustMarshal(&delegator)
	store.Set(types.DelegatorKey(
		delegator.Staker,
		delegator.Delegator,
	), b)

	indexStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefixIndex2)
	indexStore.Set(types.DelegatorKeyIndex2(
		delegator.Delegator,
		delegator.Staker,
	), []byte{1})
}

// GetDelegator returns a delegator from its index
func (k Keeper) GetDelegator(
	ctx sdk.Context,
	stakerAddress string,
	delegatorAddress string,
) (val types.Delegator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefix)

	b := store.Get(types.DelegatorKey(
		stakerAddress,
		delegatorAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegator removes a delegator from the store
func (k Keeper) RemoveDelegator(
	ctx sdk.Context,
	stakerAddress string,
	delegatorAddress string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefix)
	store.Delete(types.DelegatorKey(
		stakerAddress,
		delegatorAddress,
	))
	indexStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefixIndex2)
	indexStore.Delete(types.DelegatorKeyIndex2(
		delegatorAddress,
		stakerAddress,
	))
}

// GetAllDelegator returns all delegator
func (k Keeper) GetAllDelegator(ctx sdk.Context) (list []types.Delegator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegatorKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Delegator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
