package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Redelegate lets a user redelegate from one staker to another staker
func (k msgServer) Redelegate(goCtx context.Context, msg *types.MsgRedelegate) (*types.MsgRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.consumeRedelegationSpell(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Perform undelegation
	if err := k.performUndelegation(ctx, msg.FromStaker, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Perform undelegation
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
