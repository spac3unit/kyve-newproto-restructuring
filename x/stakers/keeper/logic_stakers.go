package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Slash() {
	// TODO implement me
}

func (k Keeper) EnsureFreeSlot(ctx sdk.Context, poolId uint64, stakeAmount uint64) error {
	if k.GetStakerCount() >= types.MaxStakers {
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
