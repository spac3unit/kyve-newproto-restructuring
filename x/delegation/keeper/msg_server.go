package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

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

func (k msgServer) WithdrawPool(goCtx context.Context, msg *types.MsgWithdrawPool) (*types.MsgWithdrawPoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, poolErr := k.poolKeeper.GetPoolWithError(ctx, msg.PoolId)
	if poolErr != nil {
		return nil, poolErr
	}

	// Check if the sender is a delegator in this pool.
	_, isDelegator := k.GetDelegator(ctx, msg.PoolId, msg.Staker, msg.Creator)
	if !isDelegator {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNotADelegator.Error())
	}

	// Create a new F1Distribution struct for interacting with delegations.
	f1Distribution := F1Distribution{
		k:                k.Keeper,
		ctx:              ctx,
		poolId:           msg.PoolId,
		stakerAddress:    msg.Staker,
		delegatorAddress: msg.Creator,
	}

	// Withdraw all rewards for the sender.
	reward := f1Distribution.Withdraw()

	// Transfer tokens from this module to sender.
	err := util.TransferToAddress(k.bankKeeper, ctx, types.ModuleName, msg.Creator, reward)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawPoolResponse{}, nil
}

func (k msgServer) UndelegatePool(ctx context.Context, pool *types.MsgUndelegatePool) (*types.MsgUndelegatePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) RedelegatePool(ctx context.Context, pool *types.MsgRedelegatePool) (*types.MsgRedelegatePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}
