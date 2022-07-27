package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// #####################
// === QUEUE ENTRIES ===
// #####################

// SetUnbondingStakeEntry ...
func (k Keeper) SetUnbondingStakeEntry(ctx sdk.Context, unbondingStakeEntry types.UnbondingStakeEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefix)
	b := k.cdc.MustMarshal(&unbondingStakeEntry)
	store.Set(types.UnbondingStakeEntryKey(
		unbondingStakeEntry.Index,
	), b)

	// Insert the same entry with a different key prefix for query lookup
	indexStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefixIndex2)
	indexStore.Set(types.UnbondingStakeEntryKeyIndex2(
		unbondingStakeEntry.Staker,
		unbondingStakeEntry.Index,
	), []byte{1})
}

// GetUnbondingStakeEntry returns a UnbondingStakingQueueEntry from its index
func (k Keeper) GetUnbondingStakeEntry(ctx sdk.Context, index uint64) (val types.UnbondingStakeEntry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefix)

	b := store.Get(types.UnbondingStakeEntryKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUnbondingStakeEntry removes a UnbondingStakingQueueEntry from the store
func (k Keeper) RemoveUnbondingStakeEntry(ctx sdk.Context, unbondingStakeEntry *types.UnbondingStakeEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefix)
	store.Delete(types.UnbondingStakeEntryKey(unbondingStakeEntry.Index))

	indexStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefixIndex2)
	indexStore.Delete(types.UnbondingStakeEntryKeyIndex2(
		unbondingStakeEntry.Staker,
		unbondingStakeEntry.Index,
	))
}

// GetAllUnbondingStakeEntries returns all staker unbondings
func (k Keeper) GetAllUnbondingStakeEntries(ctx sdk.Context) (list []types.UnbondingStakeEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakingEntryKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UnbondingStakeEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
