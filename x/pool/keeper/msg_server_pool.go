package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.AppendPool(ctx, types.Pool{
		Name:    msg.Name,
		// TODO fill rest
		Runtime:        "",
		Logo:           "",
		Config:         "",
		UploadInterval: 0,
		OperatingCost:  0,
		Paused:         false,
		MaxBundleSize:  0,
		Protocol:       nil,
		UpgradePlan:    nil,
		StartKey:       "",
		CurrentKey:     "",
		CurrentValue:   "",
		MinStake:       0,
		Funders:        nil,
		TotalFunds:     0,
	})

	// TODO emit event ?

	return &types.MsgCreatePoolResponse{}, nil
}

// TODO create missing pool changes
