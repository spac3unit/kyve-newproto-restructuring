package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// These functions are meant to be called from external modules
// For now this is the bundles module and the delegation module
// which need to interact stakers module.

// All these functions are safe in the way that they do not return errors
// and every edge case is handled within the function itself.

// GetTotalStake returns the sum of stake of all stakers who are
// currently participating in the given pool
//func (k Keeper) GetTotalStake(ctx sdk.Context, poolId uint64) uint64 {
//	return k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
//}

// GetAllStakerAddressesOfPool returns a list of all stakers
// which have a currently a valaccount registered for the given pool
func (k Keeper) GetAllStakerAddressesOfPool(ctx sdk.Context, poolId uint64) (stakers []string) {
	for _, valaccount := range k.GetAllValaccountsOfPool(ctx, poolId) {
		stakers = append(stakers, valaccount.Staker)
	}

	return
}

// GetStakeInPool returns the amount of the staker if the staker is currently
// participating in the given pool. If the staker is not in that pool
// the function returns zero as the current stake for that staker is zero.
//func (k Keeper) GetStakeInPool(ctx sdk.Context, poolId uint64, stakerAddress string) uint64 {
//	if k.DoesValaccountExist(ctx, poolId, stakerAddress) {
//		staker, _ := k.GetStaker(ctx, stakerAddress)
//		return staker.Amount
//	}
//	return 0
//}

// GetCommission returns the commission of a staker as a parsed sdk.Dec
func (k Keeper) GetCommission(ctx sdk.Context, stakerAddress string) sdk.Dec {
	staker, _ := k.GetStaker(ctx, stakerAddress)
	uploaderCommission, err := sdk.NewDecFromStr(staker.Commission)
	if err != nil {
		// TODO -> log error to metrics
	}
	return uploaderCommission
}
