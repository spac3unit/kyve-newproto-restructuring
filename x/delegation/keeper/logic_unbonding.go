package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ####################
// ==== DELEGATION ====
// ####################

func (k Keeper) StartUnbondingDelegator(ctx sdk.Context, staker string, delegatorAddress string, amount uint64) (error error) {

	//// unbondingState stores the start and the end of the queue with all unbonding entries
	//// the queue is ordered by time
	//unbondingQueueState := k.GetUnbondingDelegationQueueState(ctx)
	//
	//// Increase topIndex as a new entry is about to be appended
	//unbondingQueueState.HighIndex += 1
	//k.SetUnbondingDelegationQueueState(ctx, unbondingQueueState)
	//
	//// UnbondingEntry stores all the information which are needed to perform
	//// the undelegation at the end of the unbonding time
	//unbondingQueueEntry := types.UnbondingDelegationQueueEntry{
	//	Delegator:    delegatorAddress,
	//	Index:        unbondingQueueState.HighIndex,
	//	Staker:       staker,
	//	PoolId:       poolId,
	//	Amount:       amount,
	//	CreationTime: uint64(ctx.BlockTime().Unix()),
	//}
	//
	//k.SetUnbondingDelegationQueueEntry(ctx, unbondingQueueEntry)

	return nil
}

// ProcessDelegatorUnbondingQueue is called at the end of every block and checks the
// tail of the UnbondingDelegationQueue for Undelegations that can be performed
// This O(t) with t being the amount of undelegation-transactions which has been performed within
// a timeframe of one block
func (k Keeper) ProcessDelegatorUnbondingQueue(ctx sdk.Context) {

	//// Get Queue information
	//unbondingQueueState := k.GetUnbondingDelegationQueueState(ctx)
	//
	//// Check if queue is currently empty
	//if unbondingQueueState.LowIndex == unbondingQueueState.HighIndex {
	//	return
	//}
	//
	//// flag for computing every entry at the end of the queue which is due.
	//undelegationPerformed := true
	//// start processing the end of the queue
	//for undelegationPerformed {
	//	undelegationPerformed = false
	//
	//	// Get end of queue
	//	unbondingDelegationEntry, found := k.GetUnbondingDelegationQueueEntry(ctx, unbondingQueueState.LowIndex+1)
	//
	//	// Check if unbonding time is over
	//	if found && unbondingDelegationEntry.CreationTime+uint64(k.UnbondingDelegationTime(ctx)) < uint64(ctx.BlockTime().Unix()) {
	//
	//		// Transfer the money
	//		err := k.TransferToAddress(ctx, unbondingDelegationEntry.Delegator, unbondingDelegationEntry.Amount)
	//		if err != nil {
	//			k.PanicHalt(ctx, "Not enough money in module: "+err.Error())
	//		}
	//
	//		k.RemoveUnbondingDelegationQueueEntry(ctx, &unbondingDelegationEntry)
	//
	//		// Update tailIndex (lowIndex) of queue
	//		unbondingQueueState.LowIndex += 1
	//		k.SetUnbondingDelegationQueueState(ctx, unbondingQueueState)
	//
	//		// flags
	//		undelegationPerformed = true
	//	}
	//}
}
