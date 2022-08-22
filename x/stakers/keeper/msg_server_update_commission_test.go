package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Update Commission", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create staker
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Get default commission", func() {
		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))
	})

	It("Update commission", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.5",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))

		// wait for update
		s.CommitAfterSeconds(s.App().StakersKeeper.CommissionChangeTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal("0.5"))
	})

	It("Update commission to 0%", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))

		// wait for update
		s.CommitAfterSeconds(s.App().StakersKeeper.CommissionChangeTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal("0"))
	})

	It("Update commission to 100%", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "1",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))

		// wait for update
		s.CommitAfterSeconds(s.App().StakersKeeper.CommissionChangeTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal("1"))
	})

	It("Update commission with invalid number", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "teset",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))
	})

	It("Update commission with negative number", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "-0.5",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))
	})

	It("Update commission with to high number", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "2",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))
	})

	It("Update commission during change time", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.5",
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.2",
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.3",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))

		// wait for update
		s.CommitAfterSeconds(s.App().StakersKeeper.CommissionChangeTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal("0.3"))
	})

	It("Update commission during change time to same value", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.5",
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: "0.2",
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateCommission{
			Creator:    i.ALICE,
			Commission: stakerstypes.DefaultCommission,
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))

		// wait for update
		s.CommitAfterSeconds(s.App().StakersKeeper.CommissionChangeTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(staker.Commission).To(Equal(stakerstypes.DefaultCommission))
	})
})
