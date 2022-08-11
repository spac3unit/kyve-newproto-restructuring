package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Stake handles the logic of an SDK message that allows protocol nodes to stake
func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is already a staker.
	stakerExists := k.DoesStakerExist(ctx, msg.Creator)

	if stakerExists {
		k.AddAmountToStaker(ctx, msg.Creator, msg.Amount)
	} else {
		// Append new staker
		k.AppendStaker(ctx, types.Staker{
			Address:    msg.Creator,
			Amount:     msg.Amount,
			Commission: types.DefaultCommission,
		})
	}

	// Transfer tokens from sender to this module.
	if err := util.TransferFromAddressToModule(k.bankKeeper, ctx, msg.Creator, types.ModuleName, msg.Amount); err != nil {
		return nil, err
	}

	// Event a stake event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventStakePool{
		Address: msg.Creator,
		Amount:  msg.Amount,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgStakeResponse{}, nil
}
