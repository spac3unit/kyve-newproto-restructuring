package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Join Pool", Ordered, func() {
	s := i.NewCleanChain()

	BeforeAll(func() {
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		_, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		_, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(stakerFound).To(BeTrue())

		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Leave a pool", func() {
		s.RunTxStakersSuccess(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId: 0,
		})

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.ALICE))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal(i.BOB))
		Expect(valaccount.Points).To(BeZero())

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

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		totalStakeOfPool = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(BeZero())
	})

	It("Try to leave pool again", func() {
		_, err := s.RunTxStakers(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId: 0,
		})

		Expect(err).ToNot(BeNil())
	})

	It("Try to leave pool twice", func() {
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(HaveLen(1))

		_, err := s.RunTxStakers(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId: 0,
		})

		Expect(err).To(BeNil())

		_, err = s.RunTxStakers(&stakerstypes.MsgLeavePool{
			Creator: i.ALICE,
			PoolId: 0,
		})

		Expect(err).NotTo(BeNil())
	})
})
