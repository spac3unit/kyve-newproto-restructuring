package keeper

import (
	"encoding/binary"
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateStakerMetadata ...
func (k Keeper) UpdateStakerMetadata(ctx sdk.Context, address string, moniker string, website string, logo string) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		staker.Moniker = moniker
		staker.Website = website
		staker.Logo = logo
		k.setStaker(ctx, staker)
	}
}

// UpdateStakerCommission ...
func (k Keeper) UpdateStakerCommission(ctx sdk.Context, address string, commission string) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		staker.Commission = commission
		k.setStaker(ctx, staker)
	}
}

// RemoveStakerFromPool removes a staker from a given pool and updates
// all aggregated variables. If the staker is not in the pool nothing happens.
func (k Keeper) RemoveStakerFromPool(ctx sdk.Context, poolId uint64, address string) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		newPools, removed := util.RemoveFromArrayStable(staker.Pools, poolId)
		if removed {
			staker.Pools = newPools

			k.subtractFromTotalStake(ctx, poolId, staker.Amount)
			k.subtractOneFromCount(ctx, poolId)
			// TODO what about delegation ??

			k.setStaker(ctx, staker)
		}
	}
}

// AddStakerToPool adds a staker to a pool.
// If staker already belongs to pool, nothing happens.
// TODO consider using a sorted list for faster lookup times.
func (k Keeper) AddStakerToPool(ctx sdk.Context, poolId uint64, address string) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		if !util.Contains(staker.Pools, poolId) {
			staker.Pools = append(staker.Pools, poolId)

			k.addToTotalStake(ctx, poolId, staker.Amount)
			k.addOneToCount(ctx, poolId)
			// TODO what about delegation ??

			k.setStaker(ctx, staker)
		}
	}
}

// AddAmountToStaker adds the given amount to an already existing staker
// It also checks the status of the staker and adjust the corresponding pool stake.
func (k Keeper) AddAmountToStaker(ctx sdk.Context, address string, amount uint64) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		staker.Amount += amount

		for _, poolId := range staker.Pools {
			k.addToTotalStake(ctx, poolId, amount)
		}

		k.setStaker(ctx, staker)
	}
}

// RemoveAmountFromStaker ...
// Ensure that amount <= staker.Amount -> otherwise overflow
// TODO maybe cap it to prevent an overflow. Or maybe do nothing if that happens?
func (k Keeper) RemoveAmountFromStaker(ctx sdk.Context, address string, amount uint64) {
	staker, found := k.GetStaker(ctx, address)
	if found {
		staker.Amount -= amount
		for _, poolId := range staker.Pools {
			k.subtractFromTotalStake(ctx, poolId, amount)
		}

		k.setStaker(ctx, staker)
	}
}

func (k Keeper) GetLowestStaker(ctx sdk.Context, poolId uint64) (val types.Staker, found bool) {
	// TODO implement
	return types.Staker{}, false
}

func (k Keeper) AppendStaker(ctx sdk.Context, staker types.Staker) {
	k.setStaker(ctx, staker)
	// TODO add total staker count
}

// #############################
// #  Raw KV-Store operations  #
// #############################

// RemoveStaker removes a staker from the store
// TODO only called very rarely,
func (k Keeper) removeStaker(ctx sdk.Context, staker types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	store.Delete(types.StakerKey(
		staker.Address,
	))

	// TODO remove stake from all pools
	// TODO What about delegation ?
}

// SetStaker set a specific staker in the store from its index
func (k Keeper) setStaker(ctx sdk.Context, staker types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	b := k.cdc.MustMarshal(&staker)
	store.Set(types.StakerKey(
		staker.Address,
	), b)
}

// GetStaker returns a staker from its index
func (k Keeper) GetStaker(
	ctx sdk.Context,
	staker string,
) (val types.Staker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)

	b := store.Get(types.StakerKey(
		staker,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DoesStakerExist returns true if the staker exists
func (k Keeper) DoesStakerExist(ctx sdk.Context, staker string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	return store.Has(types.StakerKey(staker))
}

// GetAllStakers returns all staker
func (k Keeper) GetAllStakers(ctx sdk.Context) (list []types.Staker) {
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

// #############################
// #     Aggregation Data      #
// #############################

func (k Keeper) GetStakerCountOfPool(ctx sdk.Context, poolId uint64) uint64 {
	return k.getStat(ctx, poolId, types.STAKER_STATS_COUNT)
}

func (k Keeper) addOneToCount(ctx sdk.Context, poolId uint64) {
	count := k.getStat(ctx, poolId, types.STAKER_STATS_COUNT)
	k.setStat(ctx, poolId, types.STAKER_STATS_COUNT, count+1)
}

func (k Keeper) subtractOneFromCount(ctx sdk.Context, poolId uint64) {
	count := k.getStat(ctx, poolId, types.STAKER_STATS_COUNT)
	k.setStat(ctx, poolId, types.STAKER_STATS_COUNT, count-1)
}

func (k Keeper) addToTotalStake(ctx sdk.Context, poolId uint64, amount uint64) {
	stake := k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
	k.setStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE, stake+amount)
}

func (k Keeper) subtractFromTotalStake(ctx sdk.Context, poolId uint64, amount uint64) {
	stake := k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
	k.setStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE, stake-amount)
}

// getStat get the total number of pool
func (k Keeper) getStat(ctx sdk.Context, poolId uint64, statType types.STAKER_STATS) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(util.GetByteKey(statType, poolId))
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// setStat set the total number of pool
func (k Keeper) setStat(ctx sdk.Context, poolId uint64, statType types.STAKER_STATS, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(util.GetByteKey(statType, poolId), bz)
}
