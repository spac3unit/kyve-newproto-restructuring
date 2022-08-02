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

	_, valaccountFound := k.GetValaccount(ctx, msg.PoolId, msg.Creator)
	if valaccountFound {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrInvalidRequest, types.ErrAlreadyJoinedPool.Error())
	}

	errFreeSlot := k.ensureFreeSlot(ctx, msg.PoolId, staker.Amount)
	if errFreeSlot != nil {
		return nil, errFreeSlot
	}

	k.AddValaccountToPool(ctx, msg.PoolId, msg.Creator, msg.Valaddress)

	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventJoinPool{
		PoolId:  msg.PoolId,
		Staker: msg.Creator,
		Valaddress: msg.Valaddress,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgJoinPoolResponse{}, nil
}
