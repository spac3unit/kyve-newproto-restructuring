package keeper

import (
	"github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AddAmountToRewards ...
func (k Keeper) AddAmountToDelegationRewards(ctx sdk.Context, stakerAddress string, amount uint64) {
	delegationData, found := k.GetDelegationData(ctx, stakerAddress)
	if found {
		delegationData.CurrentRewards += amount
		k.SetDelegationData(ctx, delegationData)
	}
}

// SetDelegationData set a specific delegationPoolData in the store from its index
func (k Keeper) SetDelegationData(ctx sdk.Context, delegationData types.DelegationData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationDataKeyPrefix)
	b := k.cdc.MustMarshal(&delegationData)
	store.Set(types.DelegationDataKey(
		delegationData.Staker,
	), b)
}

// GetDelegationData returns a delegationPoolData from its index
func (k Keeper) GetDelegationData(ctx sdk.Context, stakerAddress string) (val types.DelegationData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationDataKeyPrefix)

	b := store.Get(types.DelegationDataKey(stakerAddress))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoesDelegationDataExist ...
func (k Keeper) DoesDelegationDataExist(ctx sdk.Context, stakerAddress string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationDataKeyPrefix)
	return store.Has(types.DelegationDataKey(stakerAddress))
}

// RemoveDelegationData removes a delegationPoolData from the store
func (k Keeper) RemoveDelegationData(ctx sdk.Context, stakerAddress string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationDataKeyPrefix)
	store.Delete(types.DelegationDataKey(stakerAddress))
}

// GetAllDelegationData returns all delegationPoolData
func (k Keeper) GetAllDelegationData(ctx sdk.Context) (list []types.DelegationData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DelegationDataKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DelegationData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
