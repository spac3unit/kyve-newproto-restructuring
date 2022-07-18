package keeper

// PayoutAmount equally splits the amount between all funders and removes
// the appropriate amount from each funder.
// All funders who can't afford the amount, are kicked out.
// Their remaining amount is transferred to the Treasury
// Function throws an error if pool ran out of funds.
func (k Keeper) PayoutAmount(amount uint64) error {

	return nil
}
