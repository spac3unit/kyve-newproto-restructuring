package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	// TODO: look into object creation belonging to pools like (delegation, valaccount etc.)
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: fill other pool configs
	k.AppendPool(ctx, types.Pool{
		Name:    msg.Name,
		Protocol:       &types.Protocol{},
		UpgradePlan:    &types.UpgradePlan{},
	})

	// TODO emit event ?

	return &types.MsgCreatePoolResponse{}, nil
}

// TODO create missing pool changes
