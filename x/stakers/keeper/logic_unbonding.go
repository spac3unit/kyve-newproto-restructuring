package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// StartUnbondingStaker creates an unbonding entry with the given amount.
// The entry is inserted into the queue and executed after the UnbondingTime is over.
func (k Keeper) StartUnbondingStaker(ctx sdk.Context, address string, amount uint64) (error error) {

	staker, stakerFound := k.GetStaker(ctx, address)
	if !stakerFound {
		return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrNoStaker.Error())
	}

	availableUnstakeAmount := uint64(0)
	if staker.Amount > staker.UnbondingAmount {
		availableUnstakeAmount = staker.Amount - staker.UnbondingAmount
	}

	if amount > availableUnstakeAmount {
		return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrUnstakeTooHigh.Error(), availableUnstakeAmount)
	}

	queueIndex := k.getNextQueueSlot(ctx, types.QUEUE_IDENTIFIER_UNSTAKING)

	// UnbondingEntry stores all the information which are needed to perform
	// the undelegation at the end of the unbonding time
	unbondingQueueEntry := types.UnbondingStakeEntry{
		Index:        queueIndex,
		Staker:       address,
		Amount:       amount,
		CreationDate: ctx.BlockTime().Unix(),
	}

	k.SetUnbondingStakeEntry(ctx, unbondingQueueEntry)

	// WARN: Not within getters file.
	// TODO add info to specs folder
	staker.UnbondingAmount += amount
	k.setStaker(ctx, staker)

	return nil
}

// ProcessStakerUnbondingQueue is called at the end of every block and checks the
// tail of the UnbondingStakingQueue for Undelegations that can be performed
func (k Keeper) ProcessStakerUnbondingQueue(ctx sdk.Context) {

	k.processQueue(ctx, types.QUEUE_IDENTIFIER_UNSTAKING, func(index uint64) bool {

		// Get queue entry in question
		queueEntry, found := k.GetUnbondingStakeEntry(ctx, index)

		if !found {
			// continue with the next entry
			return true
		} else if queueEntry.CreationDate+int64(k.UnbondingStakingTime(ctx)) <= ctx.BlockTime().Unix() {

			k.RemoveUnbondingStakeEntry(ctx, &queueEntry)

			staker, foundStaker := k.GetStaker(ctx, queueEntry.Staker)
			if foundStaker {

				// Check if stake decreased during unbonding time
				var unstakeAmount = queueEntry.Amount

				// In case the staker got slashed during unbonding
				// only unbond the available amount
				if unstakeAmount > staker.Amount {
					unstakeAmount = staker.Amount
				}

				k.RemoveAmountFromStaker(ctx, staker.Address, unstakeAmount, true)

				// Transfer tokens from sender to this module.
				if err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, staker.Address, unstakeAmount); err != nil {
					util.PanicHalt(k.upgradeKeeper, ctx, "stakers module out of funds")
				}

				ctx.EventManager().EmitTypedEvent(&types.EventUnstakePool{
					Address: staker.Address,
					Amount:  unstakeAmount,
				})
			}

			// continue with next entry
			return true
		}
		// stop processing queue
		return false
	})
}
