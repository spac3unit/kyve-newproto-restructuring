package keeper

import (
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// PayoutAmount equally splits the amount between all funders and removes
// the appropriate amount from each funder.
// All funders who can't afford the amount, are kicked out.
// Their remaining amount is transferred to the Treasury
// Function throws an error if pool ran out of funds.
func (k Keeper) PayoutAmount(ctx sdk.Context, fromPoolId uint64, toAddress string, amount uint64) error {

	pool, poolErr := k.GetPoolWithError(ctx, fromPoolId)
	if poolErr != nil {
		return poolErr
	}

	var amountPerFunder uint64
	var amountRemainder uint64

	// Ensure lowest funder can pay
	for true {
		if len(pool.Funders) == 0 {
			return sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, types.ErrFundsTooLow.Error())
		}

		// len(pool.Funders) > 0
		amountPerFunder = amount / uint64(len(pool.Funders))
		amountRemainder = amount - amountPerFunder*uint64(len(pool.Funders))

		lowestFunder := pool.GetLowestFunder()

		if amountRemainder+amountPerFunder > lowestFunder.Amount {
			pool.RemoveFunder(*lowestFunder)

			//transfer to treasury
			// TODO transfer to treasury
		} else {
			break
		}
	}

	// Remove amount from funders
	for _, funder := range pool.Funders {
		funder.Amount -= amountPerFunder
		pool.UpdateFunder(*funder)
	}

	lowestFunder := pool.GetLowestFunder()
	lowestFunder.Amount -= amountRemainder
	pool.UpdateFunder(*lowestFunder)

	err := util.TransferToAddress(k.bankKeeper, ctx, types.ModuleName, toAddress, amount)
	if err != nil {
		return err
	}

	return nil
}
