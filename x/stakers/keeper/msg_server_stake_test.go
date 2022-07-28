package keeper_test

import (
	"fmt"
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func createPool(suite *i.KeeperTestSuite, t *testing.T) {
	suite.RunTxPoolSuccess(t, &pooltypes.MsgCreatePool{
		Creator: i.ALICE,
		Name:    "Moontest",
	})

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	require.True(t, found)
	require.Equal(t, "Moontest", pool.Name)

	suite.CommitAfterSeconds(10)
}

func TestBasicStaking(t *testing.T) {

	s := i.NewCleanChain()
	createPool(&s, t)

	s.RunTxStakersSuccess(t, &stakerstypes.MsgStake{
		Creator: i.ALICE,
		Amount:  100 * i.KYVE,
	})

	staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
	require.True(t, found)
	fmt.Printf("%v\n", staker)

	count := s.App().StakersKeeper.GetStakerCountOfPool(s.Ctx(), 0)
	require.Equal(t, uint64(0), count)

	// Enter pool
	s.RunTxStakersSuccess(t, &stakerstypes.MsgJoinPool{
		Creator: i.ALICE,
		PoolId:  0,
	})

	count = s.App().StakersKeeper.GetStakerCountOfPool(s.Ctx(), 0)
	require.Equal(t, uint64(1), count)

	totalPoolStake := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)
	require.Equal(t, 100*i.KYVE, totalPoolStake)
}
