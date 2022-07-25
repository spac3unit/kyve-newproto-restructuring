package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Unstake handles the logic of an SDK message that allows protocol nodes to unstake from a specified pool.
func (k msgServer) Unstake(
	goCtx context.Context, msg *types.MsgUnstake,
) (*types.MsgUnstakeResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.StartUnbondingStaker(ctx, msg.Creator, msg.Amount)
	if err != nil {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrUnstakeTooHigh.Error())
	}

	return &types.MsgUnstakeResponse{}, nil
}
