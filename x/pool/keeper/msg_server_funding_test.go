package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func createPool(suite *i.KeeperTestSuite, t *testing.T) {
	suite.RunTxPoolSuccess(t, &types.MsgCreatePool{
		Creator: i.ALICE,
		Name:    "Moontest",
	})

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.True(t, found)
	require.Equal(t, "Moontest", pool.Name)

	suite.CommitAfterSeconds(10)
}

func TestBasicFundingDefunding(t *testing.T) {

	suite := i.NewCleanChain()
	createPool(&suite, t)

	suite.RunTxPoolSuccess(t, &types.MsgFundPool{
		Creator: i.ALICE,
		Id:      0,
		Amount:  50 * i.KYVE,
	})

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.True(t, found)
	require.Equal(t, "Moontest", pool.Name)
	require.Len(t, pool.Funders, 1)

	require.Equal(t, i.ALICE, pool.Funders[0].Account)
	require.Equal(t, 50*i.KYVE, pool.Funders[0].Amount)
	require.Equal(t, uint64(0), pool.Funders[0].PoolId)

	aliceAcc, _ := sdk.AccAddressFromBech32(i.ALICE)
	res := suite.App().BankKeeper.GetBalance(suite.Ctx(), aliceAcc, "tkyve")
	require.Equal(t, uint64(950000000000), res.Amount.Uint64())
}

func TestFundingKickOut(t *testing.T) {
	suite := i.NewCleanChain()

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.False(t, found)
	require.Equal(t, "", pool.Name)

	createPool(&suite, t)

	for k, addr := range i.DUMMY {
		suite.RunTxPoolSuccess(t, &types.MsgFundPool{
			Creator: addr,
			Id:      0,
			Amount:  uint64(k*10 + 10),
		})
	}

	pool, _ = suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.Len(t, pool.Funders, 50)
	require.Equal(t, i.DUMMY[0], pool.GetLowestFunder().Account)
	require.Equal(t, uint64(10*50+10*25*49), pool.TotalFunds)

	suite.RunTxPoolSuccess(t, &types.MsgFundPool{
		Creator: i.ALICE,
		Id:      0,
		Amount:  15, // Should kick out lowest staker with 10 and still be lowest staker
	})
	pool, _ = suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.Len(t, pool.Funders, 50)
	require.Equal(t, i.ALICE, pool.GetLowestFunder().Account)
	require.Equal(t, uint64(10*50+10*25*49-10+15), pool.TotalFunds)
}

// TODO test for two funders with same amount

// TODO test modifying current funds

// TODO test funder payout
