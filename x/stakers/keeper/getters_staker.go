package keeper

import (
	"encoding/binary"
	"math"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// AddValaccountToPool adds a valaccount to a pool.
// If valaccount already belongs to pool, nothing happens.
func (k Keeper) AddValaccountToPool(ctx sdk.Context, poolId uint64, stakerAddress string, valaddress string) {
	staker, found := k.GetStaker(ctx, stakerAddress)

	if found {
		if !k.DoesValaccountExist(ctx, poolId, stakerAddress) {
			k.setValaccount(ctx, types.Valaccount{
				PoolId:     poolId,
				Staker:     stakerAddress,
				Valaddress: valaddress,
			})
			k.addToTotalStake(ctx, poolId, staker.Amount)
			k.addOneToCount(ctx, poolId)
		}
	}
}

// RemoveValaccountFromPool removes a valaccount from a given pool and updates
// all aggregated variables. If the valaccount is not in the pool nothing happens.
func (k Keeper) RemoveValaccountFromPool(ctx sdk.Context, poolId uint64, stakerAddress string) {
	// get valaccount
	valaccount, valaccountFound := k.GetValaccount(ctx, poolId, stakerAddress)

	// if valaccount was found on pool continue
	if valaccountFound {
		// remove valaccount from pool
		k.removeValaccount(ctx, valaccount)
		k.subtractOneFromCount(ctx, poolId)
		staker, _ := k.GetStaker(ctx, stakerAddress)
		k.subtractFromTotalStake(ctx, poolId, staker.Amount)
	}
}

// AddAmountToStaker adds the given amount to an already existing staker
// It also checks the status of the staker and adjust the corresponding pool stake.
func (k Keeper) AddAmountToStaker(ctx sdk.Context, stakerAddress string, amount uint64) {
	staker, found := k.GetStaker(ctx, stakerAddress)
	if found {
		staker.Amount += amount

		for _, valaccount := range k.GetValaccountsFromStaker(ctx, stakerAddress) {
			k.addToTotalStake(ctx, valaccount.PoolId, amount)
		}

		k.setStaker(ctx, staker)
	}
}

func (k Keeper) RemoveAmountFromStaker(ctx sdk.Context, stakerAddress string, amount uint64, isUnstake bool) {
	staker, found := k.GetStaker(ctx, stakerAddress)
	if found {
		// if amount is smaller equal zero remove staker
		if amount >= staker.Amount {
			k.removeStaker(ctx, staker)

			// remove all valaccounts from pools and subtract stake
			for _, valaccount := range k.GetValaccountsFromStaker(ctx, stakerAddress) {
				k.removeValaccount(ctx, valaccount)
				k.subtractOneFromCount(ctx, valaccount.PoolId)
				k.subtractFromTotalStake(ctx, valaccount.PoolId, staker.Amount)
			}
		} else {
			staker.Amount -= amount

			// TODO: verify isUnstake
			if isUnstake {
				staker.UnbondingAmount -= amount
			}

			// subtract stake of pools
			for _, valaccount := range k.GetValaccountsFromStaker(ctx, stakerAddress) {
				k.subtractFromTotalStake(ctx, valaccount.PoolId, amount)
			}

			k.setStaker(ctx, staker)
		}
	}
}

func (k Keeper) GetLowestStaker(ctx sdk.Context, poolId uint64) (val types.Staker, found bool) {
	minAmount := uint64(math.Inf(0))

	for _, staker := range k.getAllStakersOfPool(ctx, poolId) {
		if staker.Amount <= minAmount {
			minAmount = staker.Amount
			val = staker
		}
	}

	return
}

func (k Keeper) AppendStaker(ctx sdk.Context, staker types.Staker) {
	k.setStaker(ctx, staker)
	// TODO add total staker count
}

// #############################
// #  Raw KV-Store operations  #
// #############################

func (k Keeper) getLowestStakerIndex(ctx sdk.Context, poolId uint64) (staker types.Staker, found bool) {
	// TODO implement
	return types.Staker{}, false
}

func (k Keeper) getAllStakersOfPool(ctx sdk.Context, poolId uint64) []types.Staker {
	valaccounts := k.GetAllValaccountsOfPool(ctx, poolId)

	stakers := make([]types.Staker, 0)

	for _, valaccount := range valaccounts {
		staker, _ := k.GetStaker(ctx, valaccount.Staker)
		stakers = append(stakers, staker)
	}

	return stakers
}

// removeStaker removes a staker from the store
func (k Keeper) removeStaker(ctx sdk.Context, staker types.Staker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)
	store.Delete(types.StakerKey(
		staker.Address,
	))
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

func (k Keeper) GetPaginatedStakerQuery(ctx sdk.Context, pagination *query.PageRequest) ([]types.Staker, *query.PageResponse, error) {
	var stakers []types.Staker

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.StakerKeyPrefix)

	pageRes, err := query.FilteredPaginate(store, pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var staker types.Staker
		if err := k.cdc.Unmarshal(value, &staker); err != nil {
			return false, err
		}

		if accumulate {
			stakers = append(stakers, staker)
		}

		return true, nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}
	
	return stakers, pageRes, nil
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
	bz := store.Get(util.GetByteKey(string(statType), poolId))
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
	store.Set(util.GetByteKey(string(statType), poolId), bz)
}
