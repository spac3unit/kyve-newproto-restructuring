package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Stake handles the logic of an SDK message that allows protocol nodes to stake
// If no staker object exists, a new staker will be created.
// Otherwise, the amount will be added to the existing staker
// Every user can create a staker object with some stake. However,
// only if stake + delegation is large enough to join a pool the staker
// is able to participate in the protocol
func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.DoesStakerExist(ctx, msg.Creator) {
		// If staker already exists, add amount to existing staker
		k.AddAmountToStaker(ctx, msg.Creator, msg.Amount)
	} else {
		// Create and append new staker to store
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
