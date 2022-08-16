package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Redelegate lets a user redelegate from one staker to another staker
// The user has N redelegation spells. When this transaction is executed
// one spell is used. When all spells are consumed the transaction fails.
// The user then needs to wait for the oldest spell to expire to call
// this transaction again.
func (k msgServer) Redelegate(goCtx context.Context, msg *types.MsgRedelegate) (*types.MsgRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Only errors if all spells are currently on cooldown
	if err := k.consumeRedelegationSpell(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// The redelegation is translated into an undelegation from the old staker ...
	if err := k.performUndelegation(ctx, msg.FromStaker, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// ... and a new delegation to the new staker
	if err := k.performDelegation(ctx, msg.ToStaker, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Emit a delegation event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventRedelegate{
		Address:  msg.Creator,
		FromNode: msg.FromStaker,
		ToNode:   msg.ToStaker,
		Amount:   msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgRedelegateResponse{}, nil
}
