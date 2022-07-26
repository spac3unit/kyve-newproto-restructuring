package keeper

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetUnbondingStakingQueueState returns the state for the unstaking queue
func (k Keeper) GetQueueState(ctx sdk.Context, identifier string) (state types.QueueState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	// TODO
	b := store.Get(types.UnbondingStakingQueueStateKey)

	if b == nil {
		return state
	}

	k.cdc.MustUnmarshal(b, &state)
	return
}

// SetUnbondingStakingQueueState saves the unstaking queue state
func (k Keeper) SetQueueState(ctx sdk.Context, identifier string, state types.QueueState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	// TODO
	b := k.cdc.MustMarshal(&state)
	store.Set(types.UnbondingStakingQueueStateKey, b)
}
