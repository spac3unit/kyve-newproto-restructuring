package keeper_test

import (
	"fmt"
	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func setup(suite *i.KeeperTestSuite, t *testing.T) {
	suite.RunTxPoolSuccess(t, &pooltypes.MsgCreatePool{
		Creator: i.ALICE,
		Name:    "Moontest",
	})

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.True(t, found)
	require.Equal(t, "Moontest", pool.Name)

	suite.CommitAfterSeconds(10)

	suite.RunTxStakersSuccess(t, &stakerstypes.MsgStake{
		Creator: i.ALICE,
		Amount:  100 * i.KYVE,
	})

	staker, found := suite.App().StakersKeeper.GetStaker(suite.Ctx(), i.ALICE)
	require.True(t, found)
	fmt.Printf("%v\n", staker)

	suite.CommitAfterSeconds(1)
}

func TestBasicDelegation(t *testing.T) {

	s := i.NewCleanChain()
	setup(&s, t)

	s.RunTxDelegatorSuccess(t, &types.MsgDelegate{
		Creator: i.BOB,
		Staker:  i.ALICE,
		Amount:  10 * i.KYVE,
	})

	delegationAmount := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
	fmt.Printf("%v\n", delegationAmount)

}
