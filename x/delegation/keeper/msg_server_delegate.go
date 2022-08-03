package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Delegate ...
func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Performs logical delegation without transferring the amount
	if err := k.performDelegation(ctx, msg.Staker, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Transfer tokens from sender to this module.
	if transferErr := util.TransferToModule(k.bankKeeper, ctx, types.ModuleName, msg.Creator, msg.Amount); transferErr != nil {
		return nil, transferErr
	}

	// Emit a delegation event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventDelegate{
		Address: msg.Creator,
		Node:    msg.Staker,
		Amount:  msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgDelegateResponse{}, nil
}
