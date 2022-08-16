package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// Delegate performs a safe delegation with all necessary checks
// Warning: does not transfer the amount (only the rewards)
func (k Keeper) performDelegation(ctx sdk.Context, stakerAddress string, delegatorAddress string, amount uint64) error {

	// Check if the sender is already a delegator.
	_, delegatorExists := k.GetDelegator(ctx, stakerAddress, delegatorAddress)

	if delegatorExists {
		// If the sender is already a delegator, first perform an undelegation, before then delegating.
		reward := k.f1WithdrawRewards(ctx, stakerAddress, delegatorAddress)
		err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, delegatorAddress, reward)
		if err != nil {
			return err
		}

		// Perform redelegation
		unDelegateAmount := k.f1RemoveDelegator(ctx, stakerAddress, delegatorAddress)
		k.f1CreateDelegator(ctx, stakerAddress, delegatorAddress, unDelegateAmount+amount)
	} else {
		// If the sender isn't already a delegator, simply create a new delegation entry.
		k.f1CreateDelegator(ctx, stakerAddress, delegatorAddress, amount)
	}

	// TODO where to update staker delegation and pool delegation?
	// TODO should this even be an aggregation variable

	return nil
}

// Undelegate performs a safe undelegation
// Warning: It does not create an unbonding entry; it does not transfer the delegation back (only the rewards)
func (k Keeper) performUndelegation(ctx sdk.Context, stakerAddress string, delegatorAddress string, amount uint64) error {

	// Check if the sender is already a delegator.
	_, delegatorExists := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	if !delegatorExists {
		return sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNotADelegator.Error())
	}

	// Check if the sender is trying to undelegate more than they have delegated.
	if amount > k.GetDelegationAmountOfDelegator(ctx, stakerAddress, delegatorAddress) {
		return sdkErrors.Wrapf(sdkErrors.ErrInsufficientFunds, types.ErrNotEnoughDelegation.Error(), amount)
	}

	// Withdraw all rewards for the sender.
	reward := k.f1WithdrawRewards(ctx, stakerAddress, delegatorAddress)

	// Transfer tokens from this module to sender.
	err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, delegatorAddress, reward)
	if err != nil {
		return err
	}

	// Perform an internal re-delegation.
	undelegatedAmount := k.f1RemoveDelegator(ctx, stakerAddress, delegatorAddress)
	redelegation := undelegatedAmount - amount
	k.f1CreateDelegator(ctx, stakerAddress, delegatorAddress, redelegation)

	return nil
}
