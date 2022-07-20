package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UnstakePool handles the logic of an SDK message that allows protocol nodes to unstake from a specified pool.
func (k msgServer) UnstakePool(
	goCtx context.Context, msg *types.MsgUnstakePool,
) (*types.MsgUnstakePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	if poolErr := k.poolKeeper.AssertPoolExists(ctx, msg.PoolId); poolErr != nil {
		return nil, poolErr
	}

	err := k.StartUnbondingStaker(ctx, msg.PoolId, msg.Creator, msg.Amount)
	if err != nil {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrUnstakeTooHigh.Error(), msg.PoolId)
	}

	return &types.MsgUnstakePoolResponse{}, nil
}
