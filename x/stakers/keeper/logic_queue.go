package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO maybe generalise in global module with generics ?

// insertQueueEntry inserts an entry into the queue identified by `identifier`
// It automatically updates the queue state and uses the block time.
func (k Keeper) getNextQueueSlot(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER) (index uint64) {

	// unbondingState stores the start and the end of the queue with all unbonding entries
	// the queue is ordered by time
	queueState := k.GetQueueState(ctx, identifier)

	// Increase topIndex as a new entry is about to be appended
	queueState.HighIndex += 1

	k.SetQueueState(ctx, identifier, queueState)

	return queueState.HighIndex
}

// processQueue ...
func (k Keeper) processQueue(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER, processEntry func(index uint64) bool) {
	// Get Queue information
	queueState := k.GetQueueState(ctx, identifier)

	// flag for computing every entry at the end of the queue which is due.
	// start processing the end of the queue
	for commissionChangePerformed := true; commissionChangePerformed; {
		commissionChangePerformed = false

		// TODO maybe do index height check before (here)
		entryRemoved := processEntry(queueState.LowIndex + 1)

		if entryRemoved {
			if queueState.LowIndex < queueState.HighIndex {
				queueState.LowIndex += 1
				commissionChangePerformed = true
			}
		}

	}
	k.SetQueueState(ctx, identifier, queueState)
}
