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
	staker, found := k.GetStaker(ctx, stakerAddress, poolId)

	if found {
		// Parse the provided slash percentage and panic on any errors.
		slashAmountRatio, _ := sdk.NewDecFromStr("0.01") // TODO use params

		// Compute how much we're going to slash the staker.
		slash = uint64(sdk.NewDec(int64(staker.Amount)).Mul(slashAmountRatio).RoundInt64())

		if staker.Amount == slash {
			// If we are slashing the entire staking amount, remove the staker.
			k.RemoveStaker(ctx, staker)
		} else {
			// Subtract slashing amount from staking amount, and update the pool's total stake.
			k.RemoveAmountFromStaker(ctx, staker.PoolId, staker.Account, slash)
		}

		// emit slashing event
		ctx.EventManager().EmitTypedEvent(&types.EventSlash{
			PoolId:    staker.PoolId,
			Address:   staker.Account,
			Amount:    slash,
			SlashType: slashType,
		})

		// TODO Transfer money to treasury
	}

	return slash
}

func (k Keeper) GetActiveStake(ctx sdk.Context, poolId uint64, address string) uint64 {
	staker, found := k.GetStaker(ctx, address, poolId)
	if found {
		if staker.Status == types.STAKER_STATUS_ACTIVE {
			return staker.Amount
		}
	}
	return 0
}

func (k Keeper) GetTotalStake(ctx sdk.Context, poolId uint64) uint64 {
	return k.getStat(ctx, poolId, types.STAKER_STATS_TOTAL_STAKE)
}

// AddPoint adds a point to the user and deactivates (+ Slashes ?? ) the staker
func (k Keeper) AddPoint(ctx sdk.Context, poolId uint64, address string) {

}

func (k Keeper) EnsureFreeSlot(ctx sdk.Context, poolId uint64, stakeAmount uint64) error {
	if k.GetStakerCount(ctx, poolId) >= types.MaxStakers {
		lowestStaker, _ := k.GetLowestStaker(ctx, poolId)

		if stakeAmount > lowestStaker.Amount {

			if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventStakerStatusChanged{
				PoolId:  poolId,
				Address: lowestStaker.Account,
				Status:  types.STAKER_STATUS_INACTIVE,
			}); errEmit != nil {
				return errEmit
			}

			// Move the lowest staker to inactive staker set
			k.ChangeStakerStatus(ctx, poolId, lowestStaker.Account, types.STAKER_STATUS_INACTIVE)

		} else {
			return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrStakeTooLow.Error(), lowestStaker.Amount)
		}
	}

	return nil
}
