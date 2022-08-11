package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UpdateCommission ...
func (k msgServer) UpdateCommission(goCtx context.Context, msg *types.MsgUpdateCommission) (*types.MsgUpdateCommissionResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is a protocol node (aka has staked into this pool).
	_, isStaker := k.GetStaker(ctx, msg.Creator)
	if !isStaker {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrNoStaker.Error())
	}

	// Validate commission.
	commission, err := sdk.NewDecFromStr(msg.Commission)
	if err != nil {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidCommission.Error(), msg.Commission)
	}

	if commission.LT(sdk.NewDec(int64(0))) || commission.GT(sdk.NewDec(int64(1))) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidCommission.Error(), msg.Commission)
	}

	k.orderNewCommissionChange(ctx, msg.Creator, msg.Commission)

	return &types.MsgUpdateCommissionResponse{}, nil
}
