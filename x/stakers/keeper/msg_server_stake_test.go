package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	"github.com/KYVENetwork/chain/x/stakers/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Staking", Ordered, func() {
	s := i.NewCleanChain()

	initialBalanceAlice := s.GetBalanceFromAddress(i.ALICE)

	BeforeAll(func() {
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		_, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(found).To(BeTrue())
	})

	It("Create new staker with 100 KYVE", func() {
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(found).To(BeTrue())

		Expect(initialBalanceAlice - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Stake additional 50 KYVE", func() {
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  50 * i.KYVE,
		})

		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(found).To(BeTrue())

		Expect(initialBalanceAlice - balanceAfter).To(Equal(150 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(150 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})
})

// func createPool(suite *i.KeeperTestSuite, t *testing.T) {
// 	suite.RunTxPoolSuccess(t, &pooltypes.MsgCreatePool{
// 		Creator: i.ALICE,
// 		Name:    "Moontest",
// 	})

// 	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
// 	require.True(t, found)
// 	require.Equal(t, "Moontest", pool.Name)

// 	suite.CommitAfterSeconds(10)
// }

// func TestBasicStaking(t *testing.T) {

// 	s := i.NewCleanChain()
// 	createPool(&s, t)

// 	// create staker
// 	s.RunTxStakersSuccess(t, &stakerstypes.MsgStake{
// 		Creator: i.ALICE,
// 		Amount:  100 * i.KYVE,
// 	})

// 	staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
// 	require.True(t, found)
// 	require.Equal(t, 100 * i.KYVE, staker.Amount)
// 	fmt.Printf("%v\n", staker)

// 	count := s.App().StakersKeeper.GetStakerCountOfPool(s.Ctx(), 0)
// 	require.Equal(t, uint64(0), count)

// 	// join pool
// 	s.RunTxStakersSuccess(t, &stakerstypes.MsgJoinPool{
// 		Creator: i.ALICE,
// 		PoolId:  0,
// 	})

// 	count = s.App().StakersKeeper.GetStakerCountOfPool(s.Ctx(), 0)
// 	require.Equal(t, uint64(1), count)

// 	totalPoolStake := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)
// 	require.Equal(t, 100*i.KYVE, totalPoolStake)

// 	valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
// 	require.Len(t, valaccounts, 1)

// 	// add additional stake
// 	s.RunTxStakersSuccess(t, &stakerstypes.MsgStake{
// 		Creator: i.ALICE,
// 		Amount:  50 * i.KYVE,
// 	})

// 	staker, found = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
// 	require.True(t, found)
// 	require.Equal(t, 150 * i.KYVE, staker.Amount)

// 	totalPoolStake = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)
// 	require.Equal(t, 150*i.KYVE, totalPoolStake)
// }
