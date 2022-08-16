package keeper

import (
	"github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDelegationEntry set a specific delegationEntries in the store from its index
func (k Keeper) SetDelegationEntry(ctx sdk.Context, delegationEntries types.DelegationEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationEntriesKeyPrefix)
	b := k.cdc.MustMarshal(&delegationEntries)
	store.Set(types.DelegationEntriesKey(
		delegationEntries.Staker,
		delegationEntries.KIndex,
	), b)
}

// GetDelegationEntry returns a delegationEntries from its index
func (k Keeper) GetDelegationEntry(
	ctx sdk.Context,
	stakerAddress string,
	kIndex uint64,

) (val types.DelegationEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationEntriesKeyPrefix)

	b := store.Get(types.DelegationEntriesKey(
		stakerAddress,
		kIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegationEntry removes a delegationEntries from the store
func (k Keeper) RemoveDelegationEntry(
	ctx sdk.Context,
	stakerAddress string,
	kIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationEntriesKeyPrefix)
	store.Delete(types.DelegationEntriesKey(
		stakerAddress,
		kIndex,
	))
}

// GetAllDelegationEntries returns all delegationEntries
func (k Keeper) GetAllDelegationEntries(ctx sdk.Context) (list []types.DelegationEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationEntriesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
