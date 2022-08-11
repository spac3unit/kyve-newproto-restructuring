package keeper

import (
	"context"
	"encoding/json"

	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate config json
	if !json.Valid([]byte(msg.Config)) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), msg.Config)
	}

	// Validate binaries json
	if !json.Valid([]byte(msg.Binaries)) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), msg.Binaries)
	}

	k.AppendPool(ctx, types.Pool{
		Name:           msg.Name,
		Runtime:        msg.Runtime,
		Logo:           msg.Logo,
		Config:         msg.Config,
		StartKey:       msg.StartKey,
		UploadInterval: msg.UploadInterval,
		OperatingCost:  msg.OperatingCost,
		MinStake:       msg.MinStake,
		MaxBundleSize:  msg.MaxBundleSize,
		Protocol: &types.Protocol{
			Version:     msg.Version,
			Binaries:    msg.Binaries,
			LastUpgrade: uint64(ctx.BlockTime().Unix()),
		},
		UpgradePlan: &types.UpgradePlan{},
	})

	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventCreatePool{
		Creator:        msg.Creator,
		Id:             k.GetPoolCount(ctx),
		Name:           msg.Name,
		Runtime:        msg.Runtime,
		Logo:           msg.Logo,
		Config:         msg.Config,
		StartKey:       msg.StartKey,
		UploadInterval: msg.UploadInterval,
		OperatingCost:  msg.OperatingCost,
		MinStake:       msg.MinStake,
		MaxBundleSize:  msg.MaxBundleSize,
		Version:        msg.Version,
		Binaries:       msg.Binaries,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgCreatePoolResponse{}, nil
}

// TODO create missing pool changes
