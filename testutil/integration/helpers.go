package integration

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var DEBUG_PREFIX_ADD_LEN = 0

func (suite *KeeperTestSuite) GetBalanceFromAddress(address string) uint64 {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return 0
	}

	balance := suite.App().BankKeeper.GetBalance(suite.Ctx(), accAddress, "tkyve")

	return uint64(balance.Amount.Int64())
}
