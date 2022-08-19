package keeper

import (
	"github.com/KYVENetwork/chain/util"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ChargeFundersOfPool equally splits the amount between all funders and removes
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
		// send slash to treasury
		if err := util.TransferFromModuleToTreasury(k.accountKeeper, k.distrkeeper, ctx, pooltypes.ModuleName, slashedFunds); err != nil {
			// TODO: return 0, err?
		}
	}

	// Remove amount from funders
	for _, funder := range pool.Funders {
		pool.SubFromFunder(funder.Address, amountPerFunder)
	}

	lowestFunder := pool.GetLowestFunder()
	pool.SubFromFunder(lowestFunder.Address, amountRemainder)

	k.SetPool(ctx, pool)

	return nil
}
