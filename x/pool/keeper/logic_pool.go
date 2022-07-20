package keeper

import (
	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetPoolWithError(ctx sdk.Context, poolId uint64) (types.Pool, error) {
	pool, found := k.GetPool(ctx, poolId)
	if !found {
		return types.Pool{}, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), poolId)
	}
	return pool, nil
}

// GetPool returns a pool from its id
func (k Keeper) AssertPoolExists(ctx sdk.Context, poolId uint64) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolKey)
	if store.Has(GetPoolIDBytes(poolId)) {
		return nil
	}
	return sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), poolId)
}
