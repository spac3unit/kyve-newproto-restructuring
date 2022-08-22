package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Leave Pool", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create pool
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// create staker
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		// join pool
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.ALICE,
			PoolId:     0,
			Valaddress: i.BOB,
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Leave a pool", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId:  0,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.ALICE))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal(i.BOB))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeTrue())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))

		// wait for leave pool
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		valaccountsOfStaker = s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(BeEmpty())

		_, found = s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)

		Expect(found).To(BeFalse())

		valaccountsOfPool = s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(BeEmpty())

		totalStakeOfPool = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(BeZero())
	})

	It("Try to leave pool again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId:  0,
		})

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId:  0,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(HaveLen(1))

		// wait for leave pool
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		valaccountsOfStaker = s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(BeEmpty())
	})

	It("Try to leave multiple pools", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.ALICE,
			PoolId:     1,
			Valaddress: i.CHARLIE,
		})

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId:  1,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(2))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 1, i.ALICE)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.ALICE))
		Expect(valaccount.PoolId).To(Equal(uint64(1)))
		Expect(valaccount.Valaddress).To(Equal(i.CHARLIE))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeTrue())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 1)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 1)

		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))

		// wait for leave pool
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		valaccountsOfStaker = s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		_, found = s.App().StakersKeeper.GetValaccount(s.Ctx(), 1, i.ALICE)

		Expect(found).To(BeFalse())

		valaccountsOfPool = s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 1)

		Expect(valaccountsOfPool).To(BeEmpty())

		totalStakeOfPool = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 1)

		Expect(totalStakeOfPool).To(BeZero())
	})
})
