package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDelegationSlashEntry ...
func (k Keeper) SetDelegationSlashEntry(ctx sdk.Context, slashEntry types.DelegationSlash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationSlashEntriesKeyPrefix)
	b := k.cdc.MustMarshal(&slashEntry)
	store.Set(types.DelegationEntriesKey(
		slashEntry.Staker,
		slashEntry.KIndex,
	), b)
}

// GetDelegationSlashEntry ...
func (k Keeper) GetDelegationSlashEntry(
	ctx sdk.Context,
	stakerAddress string,
	kIndex uint64,

) (val types.DelegationSlash, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationSlashEntriesKeyPrefix)

	b := store.Get(types.DelegationSlashEntriesKey(
		stakerAddress,
		kIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDelegationSlashEntry ...
func (k Keeper) RemoveDelegationSlashEntry(
	ctx sdk.Context,
	stakerAddress string,
	kIndex uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationSlashEntriesKeyPrefix)
	store.Delete(types.DelegationSlashEntriesKey(
		stakerAddress,
		kIndex,
	))
}

// GetAllDelegationSlashEntries ...
func (k Keeper) GetAllDelegationSlashEntries(ctx sdk.Context) (list []types.DelegationSlash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationSlashEntriesKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationSlash
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetAllDelegationSlashesBetween(ctx sdk.Context, staker string, start uint64, end uint64) (list []types.DelegationSlash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationSlashEntriesKeyPrefix)
	//iterator := sdk.KVStorePrefixIterator(store, util.GetByteKey(staker))

	iterator := store.Iterator(util.GetByteKey(staker, start), util.GetByteKey(staker, end+1))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationSlash
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
