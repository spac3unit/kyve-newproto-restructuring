package keeper

import (
	"github.com/KYVENetwork/chain/x/pool/types"
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
