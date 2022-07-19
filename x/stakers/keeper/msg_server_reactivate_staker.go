package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ReactivateStaker ...
func (k msgServer) ReactivateStaker(goCtx context.Context, msg *types.MsgReactivateStaker) (*types.MsgReactivateStakerResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO Create a PoolExists function on pool module which doesn't do unmarshalling etc.
	_, poolErr := k.poolKeeper.GetPoolWithError(ctx, msg.PoolId)
	if poolErr != nil {
		return nil, poolErr
	}

	staker, stakerFound := k.GetStaker(ctx, msg.Creator, msg.PoolId)
	if !stakerFound {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNoStaker.Error())
	}

	if staker.Status != types.STAKER_STATUS_INACTIVE {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrStakerAlreadyActive.Error())
	}

	errFreeSlot := k.EnsureFreeSlot(ctx, msg.PoolId, staker.Amount)
	if errFreeSlot != nil {
		return nil, errFreeSlot
	}

	k.ChangeStakerStatus(ctx, msg.PoolId, msg.Creator, types.STAKER_STATUS_ACTIVE)

	// Emit a delegation event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventStakerStatusChanged{
		PoolId:  msg.PoolId,
		Address: msg.Creator,
		Status:  types.STAKER_STATUS_ACTIVE,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgReactivateStakerResponse{}, nil
}
