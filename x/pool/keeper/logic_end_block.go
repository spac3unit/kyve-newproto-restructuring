package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// // TODO move this code to query module
//		if pool.UpgradePlan.ScheduledAt > 0 && uint64(ctx.BlockTime().Unix()) >= pool.UpgradePlan.ScheduledAt {
//			pool.Status = types.POOL_STATUS_UPGRADING
//		} else if pool.Paused {
//			pool.Status = types.POOL_STATUS_PAUSED
//		} else if len(pool.Stakers) < 2 {
//			pool.Status = types.POOL_STATUS_NOT_ENOUGH_VALIDATORS
//		} else if pool.TotalStake < pool.MinStake {
//			pool.Status = types.POOL_STATUS_NOT_ENOUGH_STAKE
//		} else if pool.TotalFunds == 0 {
//			pool.Status = types.POOL_STATUS_NO_FUNDS
//		} else {
//			pool.Status = types.POOL_STATUS_ACTIVE
//		}

func (k Keeper) HandlePoolUpgrades(goCtx context.Context) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, pool := range k.GetAllPools(ctx) {

		if pool.UpgradePlan.ScheduledAt > 0 && uint64(ctx.BlockTime().Unix()) >= pool.UpgradePlan.ScheduledAt {
			// Check if pool upgrade already has been applied
			if pool.Protocol.Version != pool.UpgradePlan.Version || pool.Protocol.Binaries != pool.UpgradePlan.Binaries {
				// perform pool upgrade
				pool.Protocol.Version = pool.UpgradePlan.Version
				pool.Protocol.Binaries = pool.UpgradePlan.Binaries
				pool.Protocol.LastUpgrade = pool.UpgradePlan.ScheduledAt
			}

			// Check if upgrade duration was reached
			if uint64(ctx.BlockTime().Unix()) >= (pool.UpgradePlan.ScheduledAt + pool.UpgradePlan.Duration) {
				// reset upgrade plan to default values
				pool.UpgradePlan = &types.UpgradePlan{}
			}

			k.SetPool(ctx, pool)
		}

	}

}
