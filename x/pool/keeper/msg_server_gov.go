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

func (k msgServer) PausePool(goCtx context.Context, p *types.GovMsgPausePool) (*types.GovMsgPausePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Attempt to fetch the pool, throw an error if not found.
	pool, found := k.GetPool(ctx, p.Id)
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), p.Id)
	}

	// Throw an error if the pool is already paused.
	if pool.Paused {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, "Pool is already paused.")
	}

	// Pause the pool and return.
	pool.Paused = true
	k.SetPool(ctx, pool)

	return &types.GovMsgPausePoolResponse{}, nil
}

func (k msgServer) UnpausePool(goCtx context.Context, p *types.GovMsgUnpausePool) (*types.GovMsgUnpausePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Attempt to fetch the pool, throw an error if not found.
	pool, found := k.GetPool(ctx, p.Id)
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), p.Id)
	}

	// Throw an error if the pool is already unpaused.
	if !pool.Paused {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, "Pool is already unpaused.")
	}

	// Unpause the pool and return.
	pool.Paused = false
	k.SetPool(ctx, pool)

	return &types.GovMsgUnpausePoolResponse{}, nil
}

func (k msgServer) PoolUpgrade(goCtx context.Context, p *types.GovMsgPoolUpgrade) (*types.GovMsgPoolUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if upgrade version and binaries are not empty
	if p.Version == "" || p.Binaries == "" {
		return nil, types.ErrInvalidArgs
	}

	var scheduledAt uint64

	// If upgrade time was already surpassed we upgrade immediately
	if p.ScheduledAt < uint64(ctx.BlockTime().Unix()) {
		scheduledAt = uint64(ctx.BlockTime().Unix())
	} else {
		scheduledAt = p.ScheduledAt
	}

	// go through every pool and schedule the upgrade
	for _, pool := range k.GetAllPools(ctx) {
		// Skip if runtime does not match
		if pool.Runtime != p.Runtime {
			continue
		}

		// Skip if pool is currently upgrading
		if pool.UpgradePlan.ScheduledAt > 0 {
			continue
		}

		// register upgrade plan
		pool.UpgradePlan = &types.UpgradePlan{
			Version:     p.Version,
			Binaries:    p.Binaries,
			ScheduledAt: scheduledAt,
			Duration:    p.Duration,
		}

		// Update the pool
		k.SetPool(ctx, pool)
	}

	return &types.GovMsgPoolUpgradeResponse{}, nil
}

func (k msgServer) CancelPoolUpgrade(goCtx context.Context, p *types.GovMsgCancelPoolUpgrade) (*types.GovMsgCancelPoolUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// go through every pool and cancel the upgrade
	for _, pool := range k.GetAllPools(ctx) {
		// Skip if runtime does not match
		if pool.Runtime != p.Runtime {
			continue
		}

		// Continue if there is no upgrade scheduled
		if pool.UpgradePlan.ScheduledAt == 0 {
			continue
		}

		// clear upgrade plan
		pool.UpgradePlan = &types.UpgradePlan{}

		// Update the pool
		k.SetPool(ctx, pool)
	}

	return &types.GovMsgCancelPoolUpgradeResponse{}, nil
}

// TODO move to bundles module?
//func handleResetPoolProposal(ctx sdk.Context, k keeper.Keeper, p *types.ResetPoolProposal) error {
//	// Attempt to fetch the pool, throw an error if not found.
//	pool, found := k.GetPool(ctx, p.Id)
//	if !found {
//		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, types.ErrPoolNotFound.Error(), p.Id)
//	}
//
//	// Check if proposal can be found with bundle id
//	_, foundProposal := k.GetProposalByPoolIdAndBundleId(ctx, p.Id, p.BundleId)
//	if !foundProposal {
//		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, types.ErrProposalNotFound.Error(), p.Id, p.BundleId)
//	}
//
//	fmt.Println("proposals")
//
//	// Delete all proposals created after reset proposal
//	for _, proposal := range k.GetProposalsByPoolIdSinceBundleId(ctx, p.Id, p.BundleId) {
//		fmt.Printf("%v\n", proposal)
//		k.RemoveProposal(ctx, proposal)
//	}
//
//	// Reset pool to latest bundle
//	if p.BundleId == 0 {
//		// if reset pool id is zero reset pool to "genesis state"
//		pool.CurrentHeight = 0
//		pool.TotalBundles = 0
//		pool.CurrentKey = ""
//		pool.CurrentValue = ""
//		pool.BundleProposal = &types.BundleProposal{
//			NextUploader: pool.BundleProposal.NextUploader,
//			CreatedAt:    uint64(ctx.BlockTime().Unix()),
//		}
//	} else {
//		// Check if reset proposal can be found with bundle id
//		resetProposal, foundResetProposal := k.GetProposalByPoolIdAndBundleId(ctx, p.Id, p.BundleId-1)
//		if !foundResetProposal {
//			return sdkerrors.Wrapf(sdkerrors.ErrNotFound, types.ErrProposalNotFound.Error(), p.Id, p.BundleId-1)
//		}
//
//		// reset pool to previous valid bundle
//		pool.CurrentHeight = resetProposal.ToHeight
//		pool.TotalBundles = p.BundleId
//		pool.CurrentKey = resetProposal.Key
//		pool.CurrentValue = resetProposal.Value
//		pool.BundleProposal = &types.BundleProposal{
//			NextUploader: pool.BundleProposal.NextUploader,
//			CreatedAt:    uint64(ctx.BlockTime().Unix()),
//		}
//	}
//
//	// Update the pool
//	k.SetPool(ctx, pool)
//
//	return nil
//}
