package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// DelegatePool handles the logic of an SDK message that allows
// delegation to a protocol node from a specified pool.
func (k msgServer) DelegatePool(
	goCtx context.Context, msg *types.MsgDelegatePool,
) (*types.MsgDelegatePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Performs logical delegation without transferring the amount
	err := k.Delegate(ctx, msg.Staker, msg.PoolId, msg.Creator, msg.Amount)
	if err != nil {
		return nil, err
	}

	//// Transfer tokens from sender to this module.
	//if transferErr := k.transferToRegistry(ctx, msg.Creator, msg.Amount); transferErr != nil {
	//	return nil, err
	//}
	//
	//// Emit a delegation event.
	//if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventDelegatePool{
	//	PoolId:  msg.Id,
	//	Address: msg.Creator,
	//	Node:    msg.Staker,
	//	Amount:  msg.Amount,
	//}); errEmit != nil {
	//	return nil, errEmit
	//}

	return &types.MsgDelegatePoolResponse{}, nil
}

func (k msgServer) WithdrawPool(ctx context.Context, pool *types.MsgWithdrawPool) (*types.MsgWithdrawPoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) UndelegatePool(ctx context.Context, pool *types.MsgUndelegatePool) (*types.MsgUndelegatePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) RedelegatePool(ctx context.Context, pool *types.MsgRedelegatePool) (*types.MsgRedelegatePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}
