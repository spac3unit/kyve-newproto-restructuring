package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/util"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Delegate handles the transaction of delegating a specific amount of $KYVE to a staker
// The only requirement for the transaction to succeed is that the staker exists
// and the user has enough balance.
func (k msgServer) Delegate(goCtx context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check if staker exist?

	// Performs logical delegation without transferring the amount
	if err := k.performDelegation(ctx, msg.Staker, msg.Creator, msg.Amount); err != nil {
		return nil, err
	}

	// Transfer tokens from sender to this module.
	if transferErr := util.TransferFromAddressToModule(k.bankKeeper, ctx, msg.Creator, types.ModuleName, msg.Amount); transferErr != nil {
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
