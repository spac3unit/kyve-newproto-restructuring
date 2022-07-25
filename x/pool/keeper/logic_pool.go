package keeper

import (
	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetPoolWithError ...
func (k Keeper) GetPoolWithError(ctx sdk.Context, poolId uint64) (types.Pool, error) {
	pool, found := k.GetPool(ctx, poolId)
	if !found {
		return types.Pool{}, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), poolId)
	}
	return pool, nil
}

// AssertPoolExists ...
func (k Keeper) AssertPoolExists(ctx sdk.Context, poolId uint64) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolKey)
	if store.Has(GetPoolIDBytes(poolId)) {
		return nil
	}
	return sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), poolId)
}

// AssertPoolCanRun ...
func (k Keeper) AssertPoolCanRun(ctx sdk.Context, poolId uint64) error {

	pool, poolErr := k.GetPoolWithError(ctx, poolId)
	if poolErr != nil {
		return poolErr
	}

	// Error if the pool has no funds.
	if len(pool.Funders) == 0 {
		return sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, types.ErrFundsTooLow.Error())
	}

	// Error if the pool is paused.
	if pool.Paused {
		return sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrPoolPaused.Error())
	}

	// Error if the pool is upgrading.
	if pool.UpgradePlan.ScheduledAt > 0 && uint64(ctx.BlockTime().Unix()) >= pool.UpgradePlan.ScheduledAt {
		return sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrPoolCurrentlyUpgrading.Error())
	}

	return nil
}
