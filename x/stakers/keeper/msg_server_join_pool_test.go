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

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create pool
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
		})

		// create staker
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100*i.KYVE,
		})
	})

	It("Staker was just created", func() {
		// ASSERT
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccounts).To(HaveLen(0))
	})

	It("Join a pool", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
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

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(100*i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Stake after joining a pool", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)
		Expect(totalStakeOfPool).To(Equal(100*i.KYVE))

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount: 50*i.KYVE,
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

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		totalStakeOfPool = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(150*i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Try to join the same pool with same valaddress again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join the same pool with a different valaddress again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})
		
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.CHARLIE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join another pool with same valaddress", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest2",
		})
		
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 1,
			Valaddress: i.BOB,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join another pool with another valaddress", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest2",
		})
		
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 1,
			Valaddress: i.CHARLIE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		Expect(valaccountsOfStaker).To(HaveLen(2))
	})
})
