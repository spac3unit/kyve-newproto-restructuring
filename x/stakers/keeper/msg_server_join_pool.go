package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// JoinPool ...
func (k msgServer) JoinPool(goCtx context.Context, msg *types.MsgJoinPool) (*types.MsgJoinPoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	if poolErr := k.poolKeeper.AssertPoolExists(ctx, msg.PoolId); poolErr != nil {
		return nil, poolErr
	}

	staker, stakerFound := k.GetStaker(ctx, msg.Creator)
	if !stakerFound {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNoStaker.Error())
	}

	errFreeSlot := k.EnsureFreeSlot(ctx, msg.PoolId, staker.Amount)
	if errFreeSlot != nil {
		return nil, errFreeSlot
	}

	k.AddStakerToPool(ctx, msg.PoolId, msg.Creator)

	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventJoinPool{
		PoolId:  msg.PoolId,
		Address: msg.Creator,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgJoinPoolResponse{}, nil
}
