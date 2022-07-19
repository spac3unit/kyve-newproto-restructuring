package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) StartUnbondingStaker(ctx sdk.Context, poolId uint64, staker string, amount uint64) (error error) {

	//// Check if user is able to unstake more
	//unbondingStaker, foundUnbondingStaker := k.GetUnbondingStaker(ctx, poolId, staker)
	//if !foundUnbondingStaker {
	//	unbondingStaker.Staker = staker
	//	unbondingStaker.PoolId = poolId
	//}
	//poolStaker, stakerFound := k.GetStaker(ctx, staker, poolId)
	//if !stakerFound {
	//	return errors.New("staker does not exist")
	//}
	//
	//if amount > poolStaker.Amount-unbondingStaker.UnbondingAmount {
	//	return errors.New("amount is too high")
	//}
	//
	//// unbondingState stores the start and the end of the queue with all unbonding entries
	//// the queue is ordered by time
	//unbondingQueueState := k.GetUnbondingStakingQueueState(ctx)
	//
	//// Increase topIndex as a new entry is about to be appended
	//unbondingQueueState.HighIndex += 1
	//k.SetUnbondingStakingQueueState(ctx, unbondingQueueState)
	//
	//// UnbondingEntry stores all the information which are needed to perform
	//// the undelegation at the end of the unbonding time
	//unbondingQueueEntry := types.UnbondingStakingQueueEntry{
	//	Index:        unbondingQueueState.HighIndex,
	//	Staker:       staker,
	//	PoolId:       poolId,
	//	Amount:       amount,
	//	CreationTime: uint64(ctx.BlockTime().Unix()),
	//}
	//
	//k.SetUnbondingStakingQueueEntry(ctx, unbondingQueueEntry)
	//
	//unbondingStaker.UnbondingAmount += amount
	//k.SetUnbondingStaker(ctx, unbondingStaker)

	return nil
}

// ProcessStakerUnbondingQueue is called at the end of every block and checks the
// tail of the UnbondingStakingQueue for Undelegations that can be performed
// This O(t) with t being the amount of undelegation-transactions which has been performed within
// a timeframe of one block
func (k Keeper) ProcessStakerUnbondingQueue(ctx sdk.Context) {

	//// Get Queue information
	//unbondingQueueState := k.GetUnbondingStakingQueueState(ctx)
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
	//	unbondingStakingEntry, found := k.GetUnbondingStakingQueueEntry(ctx, unbondingQueueState.LowIndex+1)
	//
	//	// Check if unbonding time is over
	//	if found && unbondingStakingEntry.CreationTime+uint64(k.UnbondingStakingTime(ctx)) <= uint64(ctx.BlockTime().Unix()) {
	//
	//		// Update internal UnbondingStaker value
	//		unbondingStaker, foundUnbondingStaker := k.GetUnbondingStaker(ctx, unbondingStakingEntry.PoolId, unbondingStakingEntry.Staker)
	//		if foundUnbondingStaker {
	//			unbondingStaker.UnbondingAmount -= unbondingStakingEntry.Amount
	//			k.SetUnbondingStaker(ctx, unbondingStaker)
	//
	//			// Update Pool Stakers and logic
	//
	//			pool, foundPool := k.GetPool(ctx, unbondingStakingEntry.PoolId)
	//			if !foundPool {
	//				k.PanicHalt(ctx, "Pool should exist")
	//			}
	//
	//			staker, foundStaker := k.GetStaker(ctx, unbondingStakingEntry.Staker, unbondingStakingEntry.PoolId)
	//			if foundStaker {
	//				// Check if stake decreased during unbonding time
	//				var unstakeAmount uint64 = 0
	//				if unbondingStakingEntry.Amount >= staker.Amount {
	//					// Remove user
	//					k.removeStaker(ctx, &pool, &staker)
	//					unstakeAmount = staker.Amount
	//
	//					//Remove unbondingStaker entry
	//					k.RemoveUnbondingStaker(ctx, &unbondingStaker)
	//				} else {
	//					// Reduce stake of user
	//					unstakeAmount = unbondingStakingEntry.Amount
	//
	//					staker.Amount -= unstakeAmount
	//
	//					if staker.Status == types.STAKER_STATUS_ACTIVE {
	//						pool.TotalStake -= unbondingStakingEntry.Amount
	//					} else if staker.Status == types.STAKER_STATUS_INACTIVE {
	//						pool.TotalInactiveStake -= unbondingStakingEntry.Amount
	//					}
	//
	//					k.SetStaker(ctx, staker)
	//				}
	//
	//				k.updateLowestStaker(ctx, &pool)
	//				k.SetPool(ctx, pool)
	//
	//				// Transfer the money
	//				transferError := k.TransferToAddress(ctx, unbondingStakingEntry.Staker, unstakeAmount)
	//				if transferError != nil {
	//					k.PanicHalt(ctx, "Not enough money in module: "+transferError.Error())
	//				}
	//
	//				ctx.EventManager().EmitTypedEvent(&types.EventUnstakePool{
	//					PoolId:  pool.Id,
	//					Address: unbondingStakingEntry.Staker,
	//					Amount:  unstakeAmount,
	//				})
	//			}
	//		}
	//
	//		k.RemoveUnbondingStakingQueueEntry(ctx, &unbondingStakingEntry)
	//
	//		// Update tailIndex (lowIndex) of queue
	//		unbondingQueueState.LowIndex += 1
	//		k.SetUnbondingStakingQueueState(ctx, unbondingQueueState)
	//
	//		// flags
	//		undelegationPerformed = true
	//	}
	//}
}
