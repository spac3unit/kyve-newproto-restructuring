package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakePool handles the logic of an SDK message that allows protocol nodes to stake in a specified pool.
func (k msgServer) StakePool(goCtx context.Context, msg *types.MsgStakePool) (*types.MsgStakePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO Create a PoolExists function on pool module which doesnt do unmarshalling etc.
	_, poolErr := k.poolKeeper.GetPoolWithError(ctx, msg.PoolId)
	if poolErr != nil {
		return nil, poolErr
	}

	// Check if the sender is already a staker.
	_, stakerExists := k.GetStaker(ctx, msg.Creator, msg.PoolId)

	if stakerExists {
		k.AddAmountToStaker(ctx, msg.PoolId, msg.Creator, msg.Amount)
	} else {
		// Check if we have reached the maximum number of stakers.
		// If we are staking more than the lowest staker, remove them.
		freeSlotErr := k.EnsureFreeSlot(ctx, msg.PoolId, msg.Amount)
		if freeSlotErr != nil {
			return nil, freeSlotErr
		}

		// Append new staker
		k.AppendStaker(ctx, types.Staker{
			Account:    msg.Creator,
			PoolId:     msg.PoolId,
			Amount:     msg.Amount,
			Commission: types.DefaultCommission,
			Status:     types.STAKER_STATUS_ACTIVE,
		})
	}

	// Transfer tokens from sender to this module.
	err := util.TransferToRegistry(k.bankKeeper, ctx, types.ModuleName, msg.Creator, msg.Amount)
	if err != nil {
		return nil, err
	}

	// Event a stake event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventStakePool{
		PoolId:  msg.PoolId,
		Address: msg.Creator,
		Amount:  msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgStakePoolResponse{}, nil
}
