package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Undelegate handles the logic of an SDK message that allows undelegation from a protocol node.
func (k msgServer) Undelegate(goCtx context.Context, msg *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create Unbonding queue entry
	if unbondingError := k.StartUnbondingDelegator(ctx, msg.Staker, msg.Creator, msg.Amount); unbondingError != nil {
		return nil, unbondingError
	}

	return &types.MsgUndelegateResponse{}, nil
}
