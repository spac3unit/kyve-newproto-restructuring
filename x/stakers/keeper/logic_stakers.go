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

func (k Keeper) ensureFreeSlot(ctx sdk.Context, poolId uint64, stakeAmount uint64) error {

	if k.GetStakerCountOfPool(ctx, poolId) >= types.MaxStakers /* TODO introduce param */ {
		lowestStaker, _ := k.GetLowestStaker(ctx, poolId)

		if stakeAmount > lowestStaker.Amount {

			// TODO emit leave pool event

			// Move the lowest staker to inactive staker set
			k.RemoveStakerFromPool(ctx, poolId, lowestStaker.Address)

		} else {
			return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrStakeTooLow.Error(), lowestStaker.Amount)
		}
	}

	return nil
}
