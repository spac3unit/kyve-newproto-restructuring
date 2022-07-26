package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) orderNewCommissionChange(ctx sdk.Context, staker string, commission string) {

	// Remove existing queue entry
	queueEntry, found := k.GetCommissionChangeEntryByIndex2(ctx, staker)
	if found {
		k.RemoveCommissionChangeEntry(ctx, &queueEntry)
	}

	queueIndex := k.getNextQueueSlot(ctx, "commission" /* TODO TYPE */)

	commissionChangeEntry := types.CommissionChangeEntry{
		Index:        queueIndex,
		Staker:       staker,
		Commission:   commission,
		CreationDate: ctx.BlockTime().Unix(),
	}

	k.SetCommissionChangeEntry(ctx, commissionChangeEntry)
}

// ProcessCommissionChangeQueue ...
func (k Keeper) ProcessCommissionChangeQueue(ctx sdk.Context) {

	k.processQueue(ctx, "commission" /* TODO TYPE */, func(index uint64) bool {

		// Get queue entry in question
		queueEntry, found := k.GetCommissionChangeEntry(ctx, index)

		if !found {
			// continue with the next entry
			return true
		} else if queueEntry.CreationDate+int64(k.CommissionChangeTime(ctx)) /* TODO PARAM */ <= ctx.BlockTime().Unix() {
			k.RemoveCommissionChangeEntry(ctx, &queueEntry)

			k.UpdateStakerCommission(ctx, queueEntry.Staker, queueEntry.Commission)

			// Event an event.
			ctx.EventManager().EmitTypedEvent(&types.EventUpdateCommission{
				Address:    queueEntry.Staker,
				Commission: queueEntry.Commission,
			})
			return true
		}
		return false
	})

}
