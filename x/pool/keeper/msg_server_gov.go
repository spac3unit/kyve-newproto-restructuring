package keeper

import (
	"context"
	"encoding/json"
	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePool(goCtx context.Context, p *types.GovMsgCreatePool) (*types.GovMsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check module

	// Validate config json
	if !json.Valid([]byte(p.Config)) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), p.Config)
	}

	// Validate binaries json
	if !json.Valid([]byte(p.Binaries)) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), p.Binaries)
	}

	k.AppendPool(ctx, types.Pool{
		Name:           p.Name,
		Runtime:        p.Runtime,
		Logo:           p.Logo,
		Config:         p.Config,
		StartKey:       p.StartKey,
		UploadInterval: p.UploadInterval,
		OperatingCost:  p.OperatingCost,
		MinStake:       p.MinStake,
		MaxBundleSize:  p.MaxBundleSize,
		Protocol: &types.Protocol{
			Version:     p.Version,
			Binaries:    p.Binaries,
			LastUpgrade: uint64(ctx.BlockTime().Unix()),
		},
		UpgradePlan: &types.UpgradePlan{},
	})

	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventCreatePool{
		Id:             k.GetPoolCount(ctx),
		Name:           p.Name,
		Runtime:        p.Runtime,
		Logo:           p.Logo,
		Config:         p.Config,
		StartKey:       p.StartKey,
		UploadInterval: p.UploadInterval,
		OperatingCost:  p.OperatingCost,
		MinStake:       p.MinStake,
		MaxBundleSize:  p.MaxBundleSize,
		Version:        p.Version,
		Binaries:       p.Binaries,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.GovMsgCreatePoolResponse{}, nil
}

func (k msgServer) UpdatePool(goCtx context.Context, p *types.GovMsgUpdatePool) (*types.GovMsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check module

	pool, found := k.GetPool(ctx, p.Id)
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), p.Id)
	}

	type Update struct {
		Name           *string
		Runtime        *string
		Logo           *string
		Config         *string
		UploadInterval *uint64
		OperatingCost  *uint64
		MaxBundleSize  *uint64
		MinStake       *uint64
	}

	var update Update

	if err := json.Unmarshal([]byte(p.Payload), &update); err != nil {
		return nil, err
	}

	if update.Name != nil {
		pool.Name = *update.Name
	}

	if update.Runtime != nil {
		pool.Runtime = *update.Runtime
	}

	if update.Logo != nil {
		pool.Logo = *update.Logo
	}

	if update.Config != nil {
		if json.Valid([]byte(*update.Config)) {
			pool.Config = *update.Config
		} else {
			return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), *update.Config)
		}
	}

	if update.UploadInterval != nil {
		pool.UploadInterval = *update.UploadInterval
	}

	if update.OperatingCost != nil {
		pool.OperatingCost = *update.OperatingCost
	}

	if update.MaxBundleSize != nil {
		pool.MaxBundleSize = *update.MaxBundleSize
	}

	if update.MinStake != nil {
		pool.MinStake = *update.MinStake
	}

	k.SetPool(ctx, pool)

	return &types.GovMsgUpdatePoolResponse{}, nil
}
