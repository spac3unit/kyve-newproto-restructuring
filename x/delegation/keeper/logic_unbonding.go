package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StartUnbondingDelegator ...
func (k Keeper) StartUnbondingDelegator(ctx sdk.Context, staker string, delegatorAddress string, amount uint64) {

	// the queue is ordered by time
	queueState := k.GetQueueState(ctx)

	// Increase topIndex as a new entry is about to be appended
	queueState.HighIndex += 1

	k.SetQueueState(ctx, queueState)

	// UnbondingEntry stores all the information which are needed to perform
	// the undelegation at the end of the unbonding time
	undelegationQueueEntry := types.UndelegationQueueEntry{
		Delegator:    delegatorAddress,
		Index:        queueState.HighIndex,
		Staker:       staker,
		Amount:       amount,
		CreationTime: uint64(ctx.BlockTime().Unix()),
	}

	k.SetUndelegationQueueEntry(ctx, undelegationQueueEntry)
}

// ProcessDelegatorUnbondingQueue ...
func (k Keeper) ProcessDelegatorUnbondingQueue(ctx sdk.Context) {

	// Get Queue information
	queueState := k.GetQueueState(ctx)

	// flag for computing every entry at the end of the queue which is due.
	// start processing the end of the queue
	for commissionChangePerformed := true; commissionChangePerformed; {
		commissionChangePerformed = false

		// Get end of queue
		undelegationEntry, found := k.GetUndelegationQueueEntry(ctx, queueState.LowIndex+1)

		removed := false
		// Check if unbonding time is over
		if !found {
			removed = true
		} else if undelegationEntry.CreationTime+k.UnbondingDelegationTime(ctx) < uint64(ctx.BlockTime().Unix()) {

			availableAmount := k.GetDelegationAmountOfDelegator(ctx, undelegationEntry.Staker, undelegationEntry.Delegator)
			unbondingAmount := util.MinUint64(availableAmount, undelegationEntry.Amount)

			// TODO error handling?
			k.performUndelegation(ctx, undelegationEntry.Staker, undelegationEntry.Delegator, unbondingAmount)

			// Transfer the money
			if err := util.TransferFromModuleToAddress(
				k.bankKeeper,
				ctx,
				types.ModuleName,
				undelegationEntry.Delegator,
				unbondingAmount,
			); err != nil {
				return
			}

			k.RemoveUndelegationQueueEntry(ctx, &undelegationEntry)

			// Update tailIndex (lowIndex) of queue
			queueState.LowIndex += 1
			k.SetQueueState(ctx, queueState)
			removed = true
		}

		if removed {
			if queueState.LowIndex < queueState.HighIndex {
				queueState.LowIndex += 1
				commissionChangePerformed = true
			}
		}

	}
	k.SetQueueState(ctx, queueState)
}
