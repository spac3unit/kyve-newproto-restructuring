package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PayoutAmount equally splits the amount between all funders and removes
// the appropriate amount from each funder.
// All funders who can't afford the amount, are kicked out.
// Their remaining amount is transferred to the Treasury
// Function throws an error if pool ran out of funds.
func (k Keeper) ChargeFundersOfPool(ctx sdk.Context, poolId uint64, amount uint64) error {

	pool, poolErr := k.GetPoolWithError(ctx, poolId)
	if poolErr != nil {
		return poolErr
	}

	var amountPerFunder uint64
	var amountRemainder uint64

	var slashedFunds uint64

	// Remove all funders who can't afford amountPerFunder
	for len(pool.Funders) > 0 {
		amountPerFunder = amount / uint64(len(pool.Funders))
		amountRemainder = amount - amountPerFunder*uint64(len(pool.Funders))

		lowestFunder := pool.GetLowestFunder()

		if amountRemainder+amountPerFunder > lowestFunder.Amount {
			// TODO: check if funder gets properly removed from pool
			pool.RemoveFunder(lowestFunder)
			// TODO: emit defund event
			slashedFunds += lowestFunder.Amount
		} else {
			break
		}
	}

	if slashedFunds > 0 {
		// transfer to treasury
		// TODO: transfer to treasury, (summarize all slashes and transfer in one call)
		
	}

	// Remove amount from funders
	for _, funder := range pool.Funders {
		funder.Amount -= amountPerFunder
		pool.UpdateFunder(*funder)
	}

	lowestFunder := pool.GetLowestFunder()
	lowestFunder.Amount -= amountRemainder
	pool.UpdateFunder(lowestFunder)

	return nil
}
