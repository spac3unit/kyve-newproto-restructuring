package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LeavePool(goCtx context.Context, msg *types.MsgLeavePool) (*types.MsgLeavePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.orderLeavePool(ctx, msg.Creator, msg.PoolId)
	if err != nil {
		return nil, err
	}

	return &types.MsgLeavePoolResponse{}, nil
}
