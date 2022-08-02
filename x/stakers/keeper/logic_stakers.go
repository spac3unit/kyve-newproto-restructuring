package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Slash removed a certain amount of the user and transfers it to the treasury
// If a user loses all tokens, the function takes care of removing the user completely
func (k Keeper) Slash(
	ctx sdk.Context, poolId uint64, stakerAddress string, slashType types.SlashType,
) (slash uint64) {
	staker, found := k.GetStaker(ctx, stakerAddress)

	if found {
		// Parse the provided slash percentage and panic on any errors.
		slashAmountRatio, _ := sdk.NewDecFromStr("0.01") // TODO use params

		// Compute how much we're going to slash the staker.
		slash = uint64(sdk.NewDec(int64(staker.Amount)).Mul(slashAmountRatio).RoundInt64())

		if staker.Amount == slash {
			// If we are slashing the entire staking amount, remove the staker.
			k.removeStaker(ctx, staker)
		} else {
			// Subtract slashing amount from staking amount, and update the pool's total stake.
			k.RemoveAmountFromStaker(ctx, staker.Address, slash, false)
		}

		// emit slashing event
		ctx.EventManager().EmitTypedEvent(&types.EventSlash{
			PoolId:    poolId,
			Address:   staker.Address,
			Amount:    slash,
			SlashType: slashType,
		})

		// TODO Transfer money to treasury
	}

	return slash
}

func (k Keeper) GetTotalStake(ctx sdk.Context, poolId uint64) uint64 {
	return k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
}

// AddPoint adds a point to the user and deactivates (+ Slashes ?? ) the staker
func (k Keeper) AddPoint(ctx sdk.Context, poolId uint64, address string) {

}

func (k Keeper) ResetPoints(ctx sdk.Context, poolId uint64, address string) {

}

func (k Keeper) ensureFreeSlot(ctx sdk.Context, poolId uint64, stakeAmount uint64) error {
	// check if slots are still available
	if k.GetStakerCountOfPool(ctx, poolId) >= types.MaxStakers {
		// if not - get lowest staker
		lowestStaker, _ := k.GetLowestStaker(ctx, poolId)

		// if new pool joiner has more stake than lowest staker kick him out
		if stakeAmount > lowestStaker.Amount {
			// remove lowest staker from pool
			k.RemoveValaccountFromPool(ctx, poolId, lowestStaker.Address)

			// emit event
			if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventLeavePool{
				PoolId: poolId,
				Staker: lowestStaker.Address,
			}); errEmit != nil {
				return errEmit
			}
		} else {
			return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrStakeTooLow.Error(), lowestStaker.Amount)
		}
	}

	return nil
}

func (k Keeper) GetAuthorizedStaker(ctx sdk.Context, stakerAddress string, authAddress string) {
	// TODO
}

func (k Keeper) AssertAuthorized(ctx sdk.Context, stakerAddress string, authAddress string, poolId uint64) error {
	// // Check if the sender is a protocol node (aka has staked into this pool).
	//	staker, isStaker := k.GetStaker(ctx, msg.Creator, msg.Id)
	//	if !isStaker {
	//		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrNoStaker.Error())
	//	}
	// TODO
	return nil
}
