package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetDelegationAmount(ctx sdk.Context, staker string) uint64 {
	delegationData, found := k.GetDelegationData(ctx, staker)

	if found {
		return delegationData.TotalDelegation
	}

	return 0
}

func (k Keeper) GetDelegationAmountOfDelegator(ctx sdk.Context, stakerAddress string, delegatorAddress string) uint64 {
	return k.f1GetCurrentDelegation(ctx, stakerAddress, delegatorAddress)
}

func (k Keeper) PayoutRewards(ctx sdk.Context, staker string, amount uint64, payerModuleName string) (success bool) {
	// Assert there are delegators
	if k.DoesDelegationDataExist(ctx, staker) {

		// Add amount to the rewards pool
		k.AddAmountToDelegationRewards(ctx, staker, amount)

		// Transfer tokens to the delegation module
		err := util.TransferFromModuleToModule(k.bankKeeper, ctx, payerModuleName, types.ModuleName, amount)
		if err != nil {
			util.PanicHalt(k.upgradeKeeper, ctx, "Not enough tokens in module")
			return false
		}
		return true
	}
	return false
}

func (k Keeper) SlashDelegators(ctx sdk.Context, staker string, fraction sdk.Dec) {
	slashedAmount := k.f1Slash(ctx, staker, fraction)

	// Transfer tokens to the delegation module
	if err := util.TransferFromModuleToTreasury(k.accountKeeper, k.distrkeeper, ctx, types.ModuleName, slashedAmount); err != nil {
		util.PanicHalt(k.upgradeKeeper, ctx, "Not enough tokens in module")
	}
}
