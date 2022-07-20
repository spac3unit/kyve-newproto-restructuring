package keeper

import (
	"encoding/binary"
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateStakerMetadata ...
func (k Keeper) UpdateStakerMetadata(ctx sdk.Context, poolId uint64, address string, moniker string, website string, logo string) {
	staker, found := k.GetStaker(ctx, address, poolId)
	if found {
		staker.Moniker = moniker
		staker.Website = website
		staker.Logo = logo
		k.setStaker(ctx, staker)
	}
}

// ChangeStakerStatus sets the status of the user, adjusts the sizes for active/inactive count
// and adjusts the pool total stakes / inactive stakes
func (k Keeper) ChangeStakerStatus(ctx sdk.Context, poolId uint64, address string, status types.StakerStatus) {
	staker, found := k.GetStaker(ctx, address, poolId)
	if found {
		if staker.Status == status {
			// Nothing needs to be changed
			return
		} else {
			if status == types.STAKER_STATUS_ACTIVE {
				// Change to active
				k.subtractFromTotalInactiveStake(ctx, poolId, staker.Amount)
				k.addToTotalStake(ctx, poolId, staker.Amount)
			} else if status == types.STAKER_STATUS_INACTIVE {
				// Change to inactive
				k.subtractFromTotalStake(ctx, poolId, staker.Amount)
				k.addToTotalInactive(ctx, poolId, staker.Amount)
			}
		}
	}
}

// AddAmountToStaker adds the given amount to an already existing staker
// It also checks the status of the staker and adjust the corresponding pool stake.
func (k Keeper) AddAmountToStaker(ctx sdk.Context, poolId uint64, address string, amount uint64) {
	staker, found := k.GetStaker(ctx, address, poolId)
	if found {
		staker.Amount += amount
		if staker.Status == types.STAKER_STATUS_ACTIVE {
			k.addToTotalStake(ctx, poolId, amount)
		} else if staker.Status == types.STAKER_STATUS_INACTIVE {
			k.addToTotalInactive(ctx, poolId, amount)
		}
		k.setStaker(ctx, staker)
	}
}

// RemoveAmountFromStaker ...
func (k Keeper) RemoveAmountFromStaker(ctx sdk.Context, poolId uint64, address string, amount uint64) {
	staker, found := k.GetStaker(ctx, address, poolId)
	if found {
		staker.Amount -= amount
		if staker.Status == types.STAKER_STATUS_ACTIVE {
			k.subtractFromTotalStake(ctx, poolId, amount)
		} else if staker.Status == types.STAKER_STATUS_INACTIVE {
			k.subtractFromTotalInactiveStake(ctx, poolId, amount)
		}
		k.setStaker(ctx, staker)
	}
}

func (k Keeper) AppendStaker(ctx sdk.Context, staker types.Staker) {
	k.setStaker(ctx, staker)
	k.addOneToCount(ctx, staker.PoolId)
}

func (k Keeper) GetLowestStaker(ctx sdk.Context, poolId uint64) (val types.Staker, found bool) {
	// TODO implement
	return types.Staker{}, false
}

func (k Keeper) GetStakerCount(ctx sdk.Context, poolId uint64) uint64 {
	return k.getStat(ctx, poolId, types.STAKER_STATS_COUNT)
}

// RemoveStaker removes a staker from the store
func (k Keeper) RemoveStaker(ctx sdk.Context, staker types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	store.Delete(types.StakerKey(
		staker.Account,
		staker.PoolId,
	))

	if staker.Status == types.STAKER_STATUS_ACTIVE {
		k.subtractFromTotalStake(ctx, staker.PoolId, staker.Amount)
	} else if staker.Status == types.STAKER_STATUS_INACTIVE {
		k.subtractFromTotalInactiveStake(ctx, staker.PoolId, staker.Amount)
	}

	k.subtractOneFromCount(ctx, staker.PoolId)
}

// SetStaker set a specific staker in the store from its index
func (k Keeper) setStaker(ctx sdk.Context, staker types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	b := k.cdc.MustMarshal(&staker)
	store.Set(types.StakerKey(
		staker.Account,
		staker.PoolId,
	), b)
}

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

// Aggregated data
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

func (k Keeper) addToTotalInactive(ctx sdk.Context, poolId uint64, amount uint64) {
	inactiveStake := k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_INACTIVE_STAKE)
	k.setStat(ctx, poolId, types.STAKER_STATS_TOTAL_INACTIVE_STAKE, inactiveStake+amount)
}

func (k Keeper) subtractFromTotalStake(ctx sdk.Context, poolId uint64, amount uint64) {
	stake := k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
	k.setStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE, stake-amount)
}

func (k Keeper) subtractFromTotalInactiveStake(ctx sdk.Context, poolId uint64, amount uint64) {
	stake := k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_INACTIVE_STAKE)
	k.setStat(ctx, poolId, types.STAKER_STATS_TOTAL_INACTIVE_STAKE, stake-amount)
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
