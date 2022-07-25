package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DelegatePool handles the logic of an SDK message that allows
// delegation to a protocol node from a specified pool.
func (k msgServer) DelegatePool(goCtx context.Context, msg *types.MsgDelegatePool) (*types.MsgDelegatePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Performs logical delegation without transferring the amount
	err := k.Delegate(ctx, msg.Staker, msg.PoolId, msg.Creator, msg.Amount)
	if err != nil {
		return nil, err
	}

	// Transfer tokens from sender to this module.
	if transferErr := util.TransferToRegistry(k.bankKeeper, ctx, types.ModuleName, msg.Creator, msg.Amount); transferErr != nil {
		return nil, err
	}

	// Emit a delegation event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventDelegatePool{
		PoolId:  msg.PoolId,
		Address: msg.Creator,
		Node:    msg.Staker,
		Amount:  msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgDelegatePoolResponse{}, nil
}
