package keeper

import (
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
		staker.Valaddress,
	), b)
}

// GetValaccount returns a Valaccount from its index
func (k Keeper) GetValaccount(
	ctx sdk.Context,
	poolId uint64,
	staker string,
	valaddress string,
) (val types.Staker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)

	b := store.Get(types.ValaccountKey(
		poolId,
		staker,
		valaddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// removeValaccount removes a Valaccount from the store
func (k Keeper) removeValaccount(ctx sdk.Context, staker types.Valaccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValaccountPrefix)
	store.Delete(types.ValaccountKey(
		staker.PoolId,
		staker.Staker,
		staker.Valaddress,
	))
}