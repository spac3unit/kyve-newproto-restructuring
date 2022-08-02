package keeper

import (
	"encoding/binary"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// #############################
// #  Raw KV-Store operations  #
// #############################

// setValaccount set a specific Valaccount in the store from its index
func (k Keeper) setValaccount(ctx sdk.Context, staker types.Valaccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)
	b := k.cdc.MustMarshal(&staker)
	store.Set(types.ValaccountKey(
		staker.PoolId,
		staker.Staker,
	), b)

	storeIndex2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefixIndex2)
	storeIndex2.Set(types.ValaccountKeyIndex2(
		staker.Staker,
		staker.PoolId,
	), []byte{})
}

// GetValaccount returns a Valaccount from its index
func (k Keeper) GetValaccount(
	ctx sdk.Context,
	poolId uint64,
	staker string,
) (val types.Valaccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)

	b := store.Get(types.ValaccountKey(
		poolId,
		staker,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// getAllValaccountsOfPool returns a Valaccount from its index
func (k Keeper) getAllValaccountsOfPool(
	ctx sdk.Context,
	poolId uint64,
) (val []types.Valaccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)

	iterator := sdk.KVStorePrefixIterator(store, util.GetByteKey(poolId))

	defer iterator.Close()

	for; iterator.Valid(); iterator.Next() {

		valaccount := types.Valaccount{}

		iterator.Key()
		
		k.cdc.MustUnmarshal(iterator.Value(), &valaccount)

		val = append(val, valaccount)
	}

	return 
}

// getValaccountsFromStaker returns all pools the staker has valaccounts in
func (k Keeper) getValaccountsFromStaker(
	ctx sdk.Context,
	stakerAddress string,
) (val []types.Valaccount) {
	storeIndex2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefixIndex2)

	iterator := sdk.KVStorePrefixIterator(storeIndex2, util.GetByteKey(stakerAddress))

	defer iterator.Close()

	for; iterator.Valid(); iterator.Next() {
		poolId := binary.BigEndian.Uint64(iterator.Key()[44:44+8])
		valaccount, valaccountFound := k.GetValaccount(ctx, poolId, stakerAddress)

		if valaccountFound {
			val = append(val, valaccount)
		}
	}

	return val
}

// removeValaccount removes a Valaccount from the store
func (k Keeper) removeValaccount(ctx sdk.Context, staker types.Valaccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)
	store.Delete(types.ValaccountKey(
		staker.PoolId,
		staker.Staker,
	))

	storeIndex2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefixIndex2)
	storeIndex2.Delete(types.ValaccountKeyIndex2(
		staker.Staker,
		staker.PoolId,
	))
}