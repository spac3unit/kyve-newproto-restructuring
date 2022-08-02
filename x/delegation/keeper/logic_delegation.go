package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetDelegationAmount(ctx sdk.Context, staker string) uint64 {
	delegationData, _ := k.GetDelegationData(ctx, staker)
	return delegationData.TotalDelegation
}

func (k Keeper) PayoutRewards(ctx sdk.Context, staker string, amount uint64, payerModuleName string) (success bool) {
	// Assert there are delegators
	if k.DoesDelegationDataExist(ctx, staker) {

		// Add amount to the rewards pool
		k.AddAmountToDelegationRewards(ctx, staker, amount)

		// Transfer tokens to the delegation module
		err := util.TransferInterModule(k.bankKeeper, ctx, payerModuleName, types.ModuleName, amount)
		if err != nil {
			util.PanicHalt(k.upgradeKeeper, ctx, "Not enough tokens in module")
			return false
		}
		return true
	}
	return false
}

func (k Keeper) SlashDelegators(ctx sdk.Context, staker string, amount uint64) {
	// TODO stub
}

// Delegate performs a safe delegation with all necessary checks
// Warning: does not transfer the amount (only the rewards)
func (k Keeper) performDelegation(ctx sdk.Context, stakerAddress string, delegatorAddress string, amount uint64) error {

	// Create a new F1Distribution struct for interacting with delegations.
	f1Distribution := F1Distribution{
		k:                k,
		ctx:              ctx,
		stakerAddress:    stakerAddress,
		delegatorAddress: delegatorAddress,
	}

	// Check if the sender is already a delegator.
	_, delegatorExists := k.GetDelegator(ctx, stakerAddress, delegatorAddress)

	if delegatorExists {
		// If the sender is already a delegator, first perform an undelegation, before then delegating.
		reward := f1Distribution.Withdraw()
		err := util.TransferToAddress(k.bankKeeper, ctx, types.ModuleName, delegatorAddress, reward)
		if err != nil {
			return err
		}

		// Perform redelegation
		unDelegateAmount := f1Distribution.Undelegate()
		f1Distribution.Delegate(unDelegateAmount + amount)

	} else {
		// If the sender isn't already a delegator, simply create a new delegation entry.
		f1Distribution.Delegate(amount)
	}

	// TODO where to update staker delegation and pool delegation?
	// TODO should this even be an aggregation variable

	return nil
}

// Undelegate performs a safe undelegation
// Warning: It does not create an unbonding entry; it does not transfer the delegation back (only the rewards)
func (k Keeper) performUndelegation(ctx sdk.Context, stakerAddress string, delegatorAddress string, amount uint64) error {

	// Check if the sender is already a delegator.
	delegator, delegatorExists := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	if !delegatorExists {
		return sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNotADelegator.Error())
	}

	// Check if the sender is trying to undelegate more than they have delegated.
	if amount > delegator.DelegationAmount {
		return sdkErrors.Wrapf(sdkErrors.ErrInsufficientFunds, types.ErrNotEnoughDelegation.Error(), amount)
	}

	// Create a new F1Distribution struct for interacting with delegations.
	f1Distribution := F1Distribution{
		k:                k,
		ctx:              ctx,
		stakerAddress:    stakerAddress,
		delegatorAddress: delegatorAddress,
	}

	// Withdraw all rewards for the sender.
	reward := f1Distribution.Withdraw()

	// Transfer tokens from this module to sender.
	err := util.TransferToAddress(k.bankKeeper, ctx, types.ModuleName, delegatorAddress, reward)
	if err != nil {
		return err
	}

	// Perform an internal re-delegation.
	undelegatedAmount := f1Distribution.Undelegate()
	redelegation := undelegatedAmount - amount
	f1Distribution.Delegate(redelegation)

	return nil
}
