package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
This file is responsible for implementing the F1-Fee distribution as described in
https://drops.dagstuhl.de/opus/volltexte/2020/11974/pdf/OASIcs-Tokenomics-2019-10.pdf

We recommend reading the paper first before reading the code.
This file covers all relevant methods to fully implement the algorithm.
It also takes fully care of the entire state. The only interaction needed
is covered by the available methods.

The methods starting lowerCase are only used internally
The methods starting upperCase can be freely accessed inside the keeper.
*/

// f1StartNewPeriod finishes the current period according to the F1-Paper
// It returns the index of the new period.
// delegationData is passed as a pointer and updated in this method
// it's the responsibility of the caller to save the meta-data state.
// This method only writes to the entries.
func (k Keeper) f1StartNewPeriod(ctx sdk.Context, staker string, delegationData *types.DelegationData) uint64 {
	// Ending the current period is performed by getting the Entry
	// of the previous index and adding the current quotient of
	// $T_f / n_f$

	// Get previous entry
	// F1: corresponds to $Entry_{f-1}$
	previousEntry, _ := k.GetDelegationEntries(ctx, staker, delegationData.LatestIndexK)

	// Calculate quotient of current round
	// If totalDelegation is zero the quotient is also zero
	currentPeriodValue := sdk.NewDec(0)
	if delegationData.TotalDelegation != 0 {
		decCurrentRewards := sdk.NewDec(int64(delegationData.CurrentRewards))
		decTotalDelegation := sdk.NewDec(int64(delegationData.TotalDelegation))

		// F1: $T_f / n_f$
		currentPeriodValue = decCurrentRewards.Quo(decTotalDelegation)
	}

	// Add previous entry to current one
	currentPeriodValue = currentPeriodValue.Add(previousEntry.ValueNew)

	// Increment index for the next period
	indexF := delegationData.LatestIndexK + 1

	// Add entry for new period to KV-Store
	k.SetDelegationEntries(ctx, types.DelegationEntries{
		ValueNew: currentPeriodValue,
		Staker:   staker,
		KIndex:   indexF,
	})

	// Reset the rewards for the next period back to zero
	// and update to the new index
	delegationData.CurrentRewards = 0
	delegationData.LatestIndexK = indexF

	// TODO efficient pruning?

	return indexF
}

// f1CreateDelegator creates a new delegator within the f1-logic.
// It is assumed that no delegator exists.
func (k Keeper) f1CreateDelegator(ctx sdk.Context, staker string, delegator string, amount uint64) {

	if amount == 0 {
		return
	}

	// Fetch metadata
	delegationData, found := k.GetDelegationData(ctx, staker)

	// Init default data-set, if this is the first delegator
	if !found {
		delegationData = types.DelegationData{
			Staker: staker,
		}
	}

	// Finish current round
	k.f1StartNewPeriod(ctx, staker, &delegationData)

	// Update metadata
	delegationData.TotalDelegation += amount
	delegationData.DelegatorCount += 1

	k.SetDelegator(ctx, types.Delegator{
		Staker:        staker,
		Delegator:     delegator,
		InitialAmount: amount,
		KIndex:        delegationData.LatestIndexK,
	})

	k.SetDelegationData(ctx, delegationData)
}

// f1RemoveDelegator performs a full undelegation and removes the delegator from the f1-logic
// This method returns the amount of tokens that were undelegated
// Due to slashing the undelegated amount can be lower than the initial delegated amount
func (k Keeper) f1RemoveDelegator(ctx sdk.Context, stakerAddress string, delegatorAddress string) (amount uint64) {

	// Check if delegator exists
	delegator, found := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	if !found {
		return 0
	}

	// Fetch metadata
	delegationData, found := k.GetDelegationData(ctx, stakerAddress)
	if !found {
		// Should never happen, if so there is an error in the f1-implementation
		util.PanicHalt(k.upgradeKeeper, ctx, "No delegationData although somebody is delegating")
	}

	// TODO calculate remaining balance regarding slashes
	balance := sdk.NewDec(0)

	// Start new period
	k.f1StartNewPeriod(ctx, stakerAddress, &delegationData)

	// TODO
	//delegationData.LatestIndexWasUndelegation = true

	// Update Metadata
	delegationData.TotalDelegation -= uint64(balance.RoundInt64())
	delegationData.DelegatorCount -= 1

	//Remove Delegator
	k.RemoveDelegator(ctx, delegator.Staker, delegator.Delegator)
	//Remove old entry
	k.RemoveDelegationEntries(ctx, stakerAddress, delegator.KIndex)

	// Final cleanup
	if delegationData.DelegatorCount == 0 {
		k.RemoveDelegationData(ctx, delegationData.Staker)
		k.RemoveDelegationEntries(ctx, stakerAddress, delegationData.LatestIndexK)
	} else {
		k.SetDelegationData(ctx, delegationData)
	}

	return uint64(balance.RoundInt64())
}

// f1Slash performs a slash within the f1-logic.
// It ends the current round and start a new one with reduced total delegation.
// A slash entry is created which is needed to calculate the correct delegation amount
// of every delegator.
func (k Keeper) f1Slash(ctx sdk.Context, stakerAddress string, fraction sdk.Dec) {

	delegationData, _ := k.GetDelegationData(ctx, stakerAddress)

	// Finish current period because in the new one there will be
	// a reduced total delegation for the slashed staker
	// The slash will be accounted to the period with index `slashedIndex`
	slashedIndex := k.f1StartNewPeriod(ctx, stakerAddress, &delegationData)

	k.SetDelegationSlashEntry(ctx, types.DelegationSlash{
		Staker:   stakerAddress,
		KIndex:   slashedIndex,
		Fraction: fraction,
	})

	// TODO check for rounding errors
	// remaining_total_delegation = total_delegation * (1 - fraction)
	totalDelegation := sdk.NewDec(int64(delegationData.TotalDelegation))
	slashFactor := sdk.NewDec(1).Sub(fraction)
	newTotalDelegation := totalDelegation.Mul(slashFactor)

	// Update metadata
	delegationData.TotalDelegation = uint64(newTotalDelegation.RoundInt64())
	k.SetDelegationData(ctx, delegationData)
}

// f1WithdrawRewards calculates all outstanding rewards and withdraws them from
// the f1-logic. A new period starts.
func (k Keeper) f1WithdrawRewards(ctx sdk.Context, stakerAddress string, delegatorAddress string) (rewards uint64) {
	delegator, found := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	if !found {
		return 0
	}

	// Fetch metadata
	delegationData, found := k.GetDelegationData(ctx, stakerAddress)
	if !found {
		util.PanicHalt(k.upgradeKeeper, ctx, "No delegationData although somebody is delegating")
	}

	// End current period and use it for calculating the reward
	endIndex := k.f1StartNewPeriod(ctx, stakerAddress, &delegationData)

	// According to F1 the reward is calculated as the difference between two entries multiplied by the
	// delegation amount for the period.
	// To incorporate slashing one needs to iterate all slashes and calculate the reward for every period
	// separately and then sum it.
	reward := sdk.NewDec(0)
	k.f1IterateConstantDelegationPeriods(ctx, stakerAddress, delegatorAddress, delegator.KIndex, endIndex,
		func(startIndex uint64, endIndex uint64, delegation sdk.Dec) {
			// entry difference
			firstEntry, found := k.GetDelegationEntries(ctx, stakerAddress, startIndex)
			if !found {
				panic("Entry 1 does not exist") // TODO replace with panic halt
			}

			secondEntry, found := k.GetDelegationEntries(ctx, stakerAddress, endIndex)
			if !found {
				panic("Entry 1 does not exist") // TODO replace with panic halt
			}

			difference := secondEntry.ValueNew.Sub(firstEntry.ValueNew)

			periodReward := difference.Mul(delegation)

			reward = reward.Add(periodReward)
		})

	return uint64(reward.RoundInt64())
}

// f1IterateConstantDelegationPeriods iterates all periods between minIndex and maxIndex (both inclusive)
// and calls handler() for every period with constant delegation amount
// This method iterates all slashes and additionally calls handler at least once if no slashes occurred
func (k Keeper) f1IterateConstantDelegationPeriods(ctx sdk.Context, stakerAddress string, delegatorAddress string,
	minIndex uint64, maxIndex uint64, handler func(startIndex uint64, endIndex uint64, delegation sdk.Dec)) {

	slashes := k.GetAllDelegationSlashesBetween(ctx, stakerAddress, minIndex, maxIndex)

	delegator, _ := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	delegatorBalance := sdk.NewDec(int64(delegator.InitialAmount))

	if len(slashes) == 0 {
		handler(minIndex, maxIndex, delegatorBalance)
		return
	}

	prevIndex := minIndex
	for _, slash := range slashes {
		// TODO handle rounding errors
		// TODO correct do index+1 for periods
		handler(prevIndex, slash.KIndex, delegatorBalance)
		delegatorBalance = delegatorBalance.Mul(sdk.NewDec(1).Sub(slash.Fraction))
		prevIndex = slash.KIndex
	}
	handler(prevIndex, maxIndex, delegatorBalance)
}

// f1GetCurrentDelegation calculates the current delegation of a delegator.
// I.e. the initial amount minus the slashes
func (k Keeper) f1GetCurrentDelegation(ctx sdk.Context, stakerAddress string, delegatorAddress string) uint64 {
	delegator, found := k.GetDelegator(ctx, stakerAddress, delegatorAddress)
	if !found {
		return 0
	}

	// Fetch metadata
	delegationData, found := k.GetDelegationData(ctx, stakerAddress)
	if !found {
		util.PanicHalt(k.upgradeKeeper, ctx, "No delegationData although somebody is delegating")
	}

	latestBalance := sdk.NewDec(int64(delegator.InitialAmount))
	k.f1IterateConstantDelegationPeriods(ctx, stakerAddress, delegatorAddress, delegator.KIndex, delegationData.LatestIndexK,
		func(startIndex uint64, endIndex uint64, delegation sdk.Dec) {
			latestBalance = delegation
		})

	return uint64(latestBalance.RoundInt64())
}

// f1GetOutstandingRewards calculates the current outstanding rewards without modifying the f1-state.
// This method can be used for queries.
func (k Keeper) f1GetOutstandingRewards(ctx sdk.Context) uint64 {
	return 0
}

// ================

//// F1Distribution contains all necessary objects to operate the fee-distribution algorithm
//type F1Distribution struct {
//	k                Keeper
//	ctx              sdk.Context
//	stakerAddress    string
//	delegatorAddress string
//}

//func (f1 F1Distribution) updateEntries(
//	fMinus1Index uint64,
//	currentRewards uint64,
//	totalDelegation uint64,
//	deleteOldEntry bool,
//) (entryFBalance sdk.Dec, indexF uint64) {
//	// F1Paper: Current period f = delegationPoolData.LatestIndexK + 1
//
//	// get last but one entry for F1Distribution, init with zero if it is the first delegator
//	// F1Paper: Entry_{f-1}
//	f1fMinus1, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, fMinus1Index)
//	f1fMinus1Balance := sdk.NewDec(0)
//	if found {
//		f1fMinus1Balance, _ = sdk.NewDecFromStr(f1fMinus1.Value)
//	}
//
//	// F1Paper: T_f / n_f
//	f1FinalBalance := sdk.NewDec(0)
//	if totalDelegation != 0 {
//		decCurrentRewards := sdk.NewDec(int64(currentRewards))
//		decTotalDelegation := sdk.NewDec(int64(totalDelegation))
//
//		f1FinalBalance = decCurrentRewards.Quo(decTotalDelegation)
//	}
//
//	// F1Paper Entry_f
//	entryFBalance = f1fMinus1Balance.Add(f1FinalBalance)
//
//	indexF = fMinus1Index + 1
//
//	if deleteOldEntry {
//		//Remove Old entry
//		f1.k.RemoveDelegationEntries(f1.ctx, f1.stakerAddress, fMinus1Index)
//	}
//
//	// Insert Entry_F
//	f1.k.SetDelegationEntries(f1.ctx, types.DelegationEntries{
//		Value:  entryFBalance.String(),
//		Staker: f1.stakerAddress,
//		KIndex: indexF,
//	})
//
//	return entryFBalance, indexF
//}
//
//// Delegate performs a F1-delegation on the distribution struct.
//func (f1 F1Distribution) Delegate(amount uint64) {
//
//	if amount == 0 {
//		return
//	}
//
//	// Fetch metadata
//	delegationPoolData, found := f1.k.GetDelegationData(f1.ctx, f1.stakerAddress)
//
//	// Init default data-set, if this is the first delegator
//	if !found {
//		delegationPoolData = types.DelegationData{
//			CurrentRewards:  0,
//			TotalDelegation: 0,
//			LatestIndexK:    0,
//			DelegatorCount:  0,
//			Staker:          f1.stakerAddress,
//		}
//	}
//
//	_, indexF := f1.updateEntries(delegationPoolData.LatestIndexK, delegationPoolData.CurrentRewards,
//		delegationPoolData.TotalDelegation, delegationPoolData.LatestIndexWasUndelegation)
//
//	delegationPoolData.LatestIndexWasUndelegation = false
//	// Reset Values according to F1Paper, i.e T=0
//	delegationPoolData.CurrentRewards = 0
//
//	// Update metadata
//	delegationPoolData.TotalDelegation += amount
//	delegationPoolData.DelegatorCount += 1
//
//	delegationPoolData.LatestIndexK = indexF
//
//	f1.k.SetDelegationData(f1.ctx, delegationPoolData)
//
//	f1.k.SetDelegator(f1.ctx, types.Delegator{
//		Staker:        f1.stakerAddress,
//		Delegator:     f1.delegatorAddress,
//		InitialAmount: amount,
//		KIndex:        indexF,
//	})
//}
//
//func (f1 F1Distribution) Slash(percentage string) {
//	data, _ := f1.k.GetDelegationData(f1.ctx, f1.stakerAddress)
//	f1.updateEntries(data.LatestIndexK, data.CurrentRewards, data.TotalDelegation, false /* TODO check*/)
//
//	f1.k.SetDelegationSlashEntry(f1.ctx, types.DelegationSlash{
//		Staker:     f1.stakerAddress,
//		KIndex:     data.LatestIndexK + 1, // TODO find nicer way
//		Percentage: percentage,
//	})
//
//	slashPercentage, _ := sdk.NewDecFromStr(percentage)
//	totalDelegation := sdk.NewDec(int64(data.TotalDelegation)).Mul(sdk.NewDec(1).Sub(slashPercentage))
//
//	data.TotalDelegation -= uint64(totalDelegation.RoundInt64())
//}
//
//// Undelegate
//// Undelegates the full amount.
//// Withdraw() must be called before, otherwise the reward is gone
//func (f1 F1Distribution) Undelegate() (undelegatedAmount uint64) {
//
//	delegator, found := f1.k.GetDelegator(f1.ctx, f1.stakerAddress, f1.delegatorAddress)
//	if !found {
//		return 0
//	}
//
//	// Fetch metadata
//	delegationData, found := f1.k.GetDelegationData(f1.ctx, f1.stakerAddress)
//	if !found {
//		util.PanicHalt(f1.k.upgradeKeeper, f1.ctx, "No delegationData although somebody is delegating")
//	}
//
//	balance := sdk.NewDec(int64(delegator.InitialAmount))
//	for _, slash := range f1.k.GetAllDelegationSlashesBetween(f1.ctx, f1.stakerAddress, delegator.KIndex, delegationData.LatestIndexK) {
//		slashPercentage, _ := sdk.NewDecFromStr(slash.Percentage)
//		balance = balance.Mul(sdk.NewDec(1).Sub(slashPercentage))
//	}
//
//	_, indexF := f1.updateEntries(delegationData.LatestIndexK, delegationData.CurrentRewards,
//		delegationData.TotalDelegation, delegationData.LatestIndexWasUndelegation)
//
//	// add flag that entry can be deleted after next entry is created
//	delegationData.LatestIndexWasUndelegation = true
//
//	// Reset Values according to F1Paper, i.e T=0
//	delegationData.CurrentRewards = 0
//	delegationData.LatestIndexK = indexF
//
//	// Update Metadata
//	delegationData.TotalDelegation -= uint64(balance.RoundInt64())
//	delegationData.DelegatorCount -= 1
//
//	//Remove Delegator
//	f1.k.RemoveDelegator(f1.ctx, delegator.Staker, delegator.Delegator)
//
//	//Remove old entry
//	f1.k.RemoveDelegationEntries(f1.ctx, f1.stakerAddress, delegator.KIndex)
//
//	if delegationData.DelegatorCount == 0 {
//		f1.k.RemoveDelegationData(f1.ctx, delegationData.Staker)
//		f1.k.RemoveDelegationEntries(f1.ctx, f1.stakerAddress, indexF)
//	} else {
//		f1.k.SetDelegationData(f1.ctx, delegationData)
//	}
//
//	return delegator.InitialAmount
//}
//
//// Withdraw
//// F1Withdraw updates the states for F1-Algorithm and returns the amount of coins the user has earned.
//// The Method does NOT transfer the money.
//func (f1 F1Distribution) Withdraw() (reward uint64) {
//
//	delegator, found := f1.k.GetDelegator(f1.ctx, f1.stakerAddress, f1.delegatorAddress)
//	if !found {
//		return 0
//	}
//
//	// Fetch metadata
//	delegationData, found := f1.k.GetDelegationData(f1.ctx, f1.stakerAddress)
//	if !found {
//		util.PanicHalt(f1.k.upgradeKeeper, f1.ctx, "No delegationData although somebody is delegating")
//	}
//
//	_, indexF := f1.updateEntries(delegationData.LatestIndexK, delegationData.CurrentRewards,
//		delegationData.TotalDelegation, delegationData.LatestIndexWasUndelegation)
//
//	delegationData.LatestIndexWasUndelegation = false
//
//	// Reset Values according to F1Paper, i.e T=0
//	delegationData.CurrentRewards = 0
//	delegationData.LatestIndexK = indexF
//
//	f1.k.SetDelegationData(f1.ctx, delegationData)
//
//	//Calculate Reward
//	f1K, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, delegator.KIndex)
//	if !found {
//		util.PanicHalt(f1.k.upgradeKeeper, f1.ctx, "Delegator does not have entry")
//	}
//
//	//Remove Old entry
//	f1.k.RemoveDelegationEntries(f1.ctx, f1.stakerAddress, delegator.KIndex)
//
//	//Update Delegator
//	delegator.KIndex = indexF
//	f1.k.SetDelegator(f1.ctx, delegator)
//
//	rewards := sdk.NewDec(0)
//
//	dummyFirst := types.DelegationSlash{
//		Staker:     f1.stakerAddress,
//		KIndex:     f1K.KIndex,
//		Percentage: "0",
//	}
//	dummyLast := types.DelegationSlash{
//		Staker:     f1.stakerAddress,
//		KIndex:     indexF,
//		Percentage: "0",
//	}
//	slashes := []types.DelegationSlash{dummyFirst}
//	slashes = append(slashes, f1.k.GetAllDelegationSlashesBetween(f1.ctx, f1.stakerAddress, f1K.KIndex, indexF)...)
//	slashes = append(slashes, dummyLast)
//
//	delegatorBalance := sdk.NewDec(int64(delegator.InitialAmount))
//	for i := 1; i < len(slashes); i++ {
//		difference := f1.getEntriesDifference(slashes[i-1].KIndex, slashes[i].KIndex)
//
//		slashPercentage, _ := sdk.NewDecFromStr(slashes[i-1].Percentage)
//		delegatorBalance = delegatorBalance.Mul(sdk.NewDec(1).Sub(slashPercentage))
//
//		rewards = rewards.Add(difference.Mul(delegatorBalance))
//	}
//
//	return uint64(rewards.RoundInt64())
//}
//
//func (k Keeper) iterateSlashes(ctx sdk.Context, stakerAddress string, startIndex uint64, endIndex uint64,
//	handler func(slash types.DelegationSlash)) {
//
//	dummyFirst := types.DelegationSlash{
//		Staker:     stakerAddress,
//		KIndex:     startIndex,
//		Percentage: "0",
//	}
//	// TODO find clean way
//	dummyLast := types.DelegationSlash{
//		Staker:     stakerAddress,
//		KIndex:     endIndex,
//		Percentage: "0",
//	}
//
//	slashes := []types.DelegationSlash{dummyFirst}
//	slashes = append(slashes, k.GetAllDelegationSlashesBetween(ctx, stakerAddress, startIndex, endIndex)...)
//	slashes = append(slashes, dummyLast)
//
//	for i := 1; i < len(slashes); i++ {
//		handler(slashes[i-1])
//	}
//
//}
//
//func (f1 F1Distribution) getEntriesDifference(index1 uint64, index2 uint64) (reward sdk.Dec) {
//	// get last but one entry for F1Distribution, init with zero if it is the first delegator
//	// F1Paper: Entry_{f-1}
//	f1Index1, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, index1)
//	f1Index1Balance := sdk.NewDec(0)
//	if found {
//		f1Index1Balance, _ = sdk.NewDecFromStr(f1Index1.Value)
//	}
//
//	f1Index2, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, index2)
//	f1Index2Balance := sdk.NewDec(0)
//	if found {
//		f1Index2Balance, _ = sdk.NewDecFromStr(f1Index2.Value)
//	}
//
//	return f1Index2Balance.Sub(f1Index1Balance)
//}
//
//// getCurrentReward
//// Calculates and returns the current reward, *without* performing any state changes
//// TODO include slashing
//func (f1 F1Distribution) getCurrentReward() (reward uint64) {
//
//	delegator, found := f1.k.GetDelegator(f1.ctx, f1.stakerAddress, f1.delegatorAddress)
//	if !found {
//		return 0
//	}
//
//	// Fetch metadata
//	delegationPoolData, found := f1.k.GetDelegationData(f1.ctx, f1.stakerAddress)
//	if !found {
//		util.PanicHalt(f1.k.upgradeKeeper, f1.ctx, "No delegationData although somebody is delegating")
//	}
//
//	// get last but one entry for F1Distribution, init with zero if it is the first delegator
//	// F1Paper: Entry_{f-1}
//	f1fMinus1, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, delegationPoolData.LatestIndexK)
//	f1fMinus1Balance := sdk.NewDec(0)
//	if found {
//		f1fMinus1Balance, _ = sdk.NewDecFromStr(f1fMinus1.Value)
//	}
//
//	// F1Paper: T_f / n_f
//	f1FinalBalance := sdk.NewDec(0)
//	if delegationPoolData.TotalDelegation != 0 {
//		decCurrentRewards := sdk.NewDec(int64(delegationPoolData.CurrentRewards))
//		decTotalDelegation := sdk.NewDec(int64(delegationPoolData.TotalDelegation))
//
//		f1FinalBalance = decCurrentRewards.Quo(decTotalDelegation)
//	}
//
//	f1FinalBalance = f1FinalBalance.Add(f1fMinus1Balance)
//
//	//Calculate Reward
//	f1K, found := f1.k.GetDelegationEntries(f1.ctx, f1.stakerAddress, delegator.KIndex)
//	if !found {
//		util.PanicHalt(f1.k.upgradeKeeper, f1.ctx, "Delegator does not have an entry")
//	}
//
//	f1kBalance, _ := sdk.NewDecFromStr(f1K.Value)
//	return uint64(sdk.NewDec(int64(delegator.InitialAmount)).Mul(f1FinalBalance.Sub(f1kBalance)).RoundInt64())
//}
