package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ChangeStakerStatus sets the status of the user, adjusts the sizes for active/inactive count
// and adjusts the pool total stakes / inactive stakes
func (k Keeper) UpdateStakerMetadata(ctx sdk.Context, poolId uint64, address string, moniker string, website string, logo string) {
	// TODO implement me
}

// ChangeStakerStatus sets the status of the user, adjusts the sizes for active/inactive count
// and adjusts the pool total stakes / inactive stakes
func (k Keeper) ChangeStakerStatus(ctx sdk.Context, poolId uint64, address string, status types.StakerStatus) {
	// TODO implement me
}

// AddAmountToStaker adds the given amount to an already existing staker
// It also checks the status of the staker and adjust the corresponding pool stake.
func (k Keeper) AddAmountToStaker(ctx sdk.Context, poolId uint64, address string, amount uint64) {

}

func (k Keeper) AppendStaker(ctx sdk.Context, staker types.Staker) {
	// TODO implement
}

func (k Keeper) GetLowestStaker(ctx sdk.Context, poolId uint64) (val types.Staker, found bool) {
	return types.Staker{}, false
}

func (k Keeper) GetStakerCount() uint64 {
	// TODO implement
	return 0
}

//// SetStaker set a specific staker in the store from its index
//func (k Keeper) SetStaker(ctx sdk.Context, staker types.Staker) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
//	b := k.cdc.MustMarshal(&staker)
//	store.Set(types.StakerKey(
//		staker.Account,
//		staker.PoolId,
//	), b)
//}

// GetStaker returns a staker from its index
func (k Keeper) GetStaker(
	ctx sdk.Context,
	staker string,
	poolId uint64,
) (val types.Staker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)

	b := store.Get(types.StakerKey(
		staker,
		poolId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStaker removes a staker from the store
func (k Keeper) RemoveStaker(
	ctx sdk.Context,
	staker string,
	poolId uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	store.Delete(types.StakerKey(
		staker,
		poolId,
	))
}

// GetAllStaker returns all staker
func (k Keeper) GetAllStaker(ctx sdk.Context) (list []types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Staker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
