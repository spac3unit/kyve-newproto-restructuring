package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Unstake handles the logic of an SDK message that allows protocol nodes to unstake
// The amount is not transferred immediately, instead an unstake entry is created
// and pushed to a queue. Once the unbonding time is reached the user will
// receive the balance. If the node got slashed during the unbonding the user might
// receive less balance than unstaked.
func (k msgServer) Unstake(goCtx context.Context, msg *types.MsgUnstake) (*types.MsgUnstakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Try to create unbonding entry
	err := k.StartUnbondingStaker(ctx, msg.Creator, msg.Amount)
	if err != nil {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrUnstakeTooHigh.Error())
	}

	return &types.MsgUnstakeResponse{}, nil
}
