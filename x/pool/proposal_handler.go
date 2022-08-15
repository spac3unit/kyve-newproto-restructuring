package pool

import (
	"encoding/json"

	"github.com/KYVENetwork/chain/x/pool/keeper"
	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func NewRegistryProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CreatePoolProposal:
			return handleCreatePoolProposal(ctx, k, c)
		// case *types.UpdatePoolProposal:
		// 	return handleUpdatePoolProposal(ctx, k, c)
		// case *types.PausePoolProposal:
		// 	return handlePausePoolProposal(ctx, k, c)
		// case *types.UnpausePoolProposal:
		// 	return handleUnpausePoolProposal(ctx, k, c)
		// case *types.SchedulePoolUpgradeProposal:
		// 	return handleSchedulePoolUpgradeProposal(ctx, k, c)
		// case *types.CancelPoolUpgradeProposal:
		// 	return handleCancelPoolUpgradeProposal(ctx, k, c)

		default:
			return sdkErrors.Wrapf(sdkErrors.ErrUnknownRequest, "unrecognized pool proposal content type: %T", c)
		}
	}
}

func handleCreatePoolProposal(ctx sdk.Context, k keeper.Keeper, p *types.CreatePoolProposal) error {
	// Validate config json
	if !json.Valid([]byte(p.Config)) {
		return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), p.Config)
	}

	// Validate binaries json
	if !json.Valid([]byte(p.Binaries)) {
		return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), p.Binaries)
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
		return errEmit
	}

	return nil
}

func handleUpdatePoolProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdatePoolProposal) error {
	// Validate config json
	if !json.Valid([]byte(p.Config)) {
		return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidJson.Error(), p.Config)
	}

	pool, poolFound := k.GetPool(ctx, p.Id)

	if !poolFound {
		return sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error())
	}

	pool.Name = p.Name
	pool.Runtime = p.Runtime
	pool.Logo = p.Logo
	pool.Config = p.Config
	pool.UploadInterval = p.UploadInterval
	pool.OperatingCost = p.OperatingCost
	pool.MinStake = p.MinStake
	pool.MaxBundleSize = p.MaxBundleSize

	// if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventCreatePool{
	// 	Id:             k.GetPoolCount(ctx),
	// 	Name:           p.Name,
	// 	Runtime:        p.Runtime,
	// 	Logo:           p.Logo,
	// 	Config:         p.Config,
	// 	StartKey:       p.StartKey,
	// 	UploadInterval: p.UploadInterval,
	// 	OperatingCost:  p.OperatingCost,
	// 	MinStake:       p.MinStake,
	// 	MaxBundleSize:  p.MaxBundleSize,
	// 	Version:        p.Version,
	// 	Binaries:       p.Binaries,
	// }); errEmit != nil {
	// 	return errEmit
	// }

	return nil
}