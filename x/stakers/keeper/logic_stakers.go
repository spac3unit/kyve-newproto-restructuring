package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/stakers/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

func (k Keeper) GetSlashFraction(ctx sdk.Context, slashType types.SlashType) (slashAmountRatio sdk.Dec) {
	// Retrieve slash fraction from params
	switch slashType {
	case types.SLASH_TYPE_TIMEOUT:
		slashAmountRatio, _ = sdk.NewDecFromStr(k.TimeoutSlash(ctx))
	case types.SLASH_TYPE_VOTE:
		slashAmountRatio, _ = sdk.NewDecFromStr(k.VoteSlash(ctx))
	case types.SLASH_TYPE_UPLOAD:
		slashAmountRatio, _ = sdk.NewDecFromStr(k.UploadSlash(ctx))
	}
	return
}

// Slash removed a certain amount of the user and transfers it to the treasury
// If a user loses all tokens, the function takes care of removing the user completely
func (k Keeper) Slash(
	ctx sdk.Context, poolId uint64, stakerAddress string, slashType types.SlashType,
) (slashedAmount uint64) {
	staker, found := k.GetStaker(ctx, stakerAddress)
	if found {

		// Retrieve slash fraction from params
		var slashAmountRatio = k.GetSlashFraction(ctx, slashType)

		// Compute how much we're going to slash the staker.
		slash := uint64(sdk.NewDec(int64(staker.Amount)).Mul(slashAmountRatio).RoundInt64())

		// remove amount - staker gets removed slash if greater equal than stake
		k.RemoveAmountFromStaker(ctx, staker.Address, slash, false)

		// send slash to treasury
		if err := util.TransferFromModuleToTreasury(k.accountKeeper, k.distrkeeper, ctx, stakertypes.ModuleName, slash); err != nil {
			util.PanicHalt(k.upgradeKeeper, ctx, "not enough money in module")
			return 0
		}

		if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventSlash{
			PoolId:    poolId,
			Address:   staker.Address,
			Amount:    slash,
			SlashType: slashType,
		}); errEmit != nil {
			// TODO log metrics
		}

		return slash
	}

	return 0
}

func (k Keeper) getLowestStaker(ctx sdk.Context, poolId uint64) (val types.Staker, found bool) {
	var minAmount uint64 = math.MaxUint64

	for _, staker := range k.getAllStakersOfPool(ctx, poolId) {
		if staker.Amount+k.delegationKeeper.GetDelegationAmount(ctx, staker.Address) <= minAmount {
			minAmount = staker.Amount
			val = staker
		}
	}

	return
}

func (k Keeper) ensureFreeSlot(ctx sdk.Context, poolId uint64, stakeAmount uint64) error {
	// check if slots are still available
	if k.GetStakerCountOfPool(ctx, poolId) >= types.MaxStakers {
		// if not - get lowest staker
		lowestStaker, _ := k.getLowestStaker(ctx, poolId)

		// if new pool joiner has more stake than lowest staker kick him out
		if stakeAmount > lowestStaker.Amount {
			// remove lowest staker from pool
			k.RemoveValaccountFromPool(ctx, poolId, lowestStaker.Address)

			// emit event
			if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventLeavePool{
				PoolId: poolId,
				Staker: lowestStaker.Address,
			}); errEmit != nil {
				return errEmit
			}
		} else {
			return sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrStakeTooLow.Error(), lowestStaker.Amount)
		}
	}

	return nil
}

func (k Keeper) AssertValaccountAuthorized(ctx sdk.Context, poolId uint64, stakerAddress string, valaddress string) error {
	valaccount, found := k.GetValaccount(ctx, poolId, stakerAddress)

	if !found {
		return types.ErrValaccountUnauthorized
	}

	if valaccount.Valaddress != valaddress {
		return types.ErrValaccountUnauthorized
	}

	return nil
}
