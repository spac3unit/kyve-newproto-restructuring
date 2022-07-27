package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueueState ...
func (k Keeper) GetQueueState(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER) (state types.QueueState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	b := store.Get(identifier)

	if b == nil {
		return state
	}

	k.cdc.MustUnmarshal(b, &state)
	return
}

// SetQueueState ...
func (k Keeper) SetQueueState(ctx sdk.Context, identifier types.QUEUE_IDENTIFIER, state types.QueueState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	b := k.cdc.MustMarshal(&state)
	store.Set(identifier, b)
}
