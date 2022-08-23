package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/stakers/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Staking", Ordered, func() {
	s := i.NewCleanChain()

	initialBalance := s.GetBalanceFromAddress(i.STAKER_0)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Create new staker with 100 KYVE", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(BeEmpty())
	})

	It("Stake additional 50 KYVE", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  50 * i.KYVE,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(150 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(150 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Stake with more than available balance", func() {
		// ACT
		currentBalance := s.GetBalanceFromAddress(i.STAKER_0)

		s.RunTxStakersError(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  currentBalance + 1,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)
		Expect(initialBalance - balanceAfter).To(BeZero())

		_, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeFalse())
	})

	It("Create a second staker", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.BOB,
			Amount:  150 * i.KYVE,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.BOB)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.BOB)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.BOB)

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(150 * i.KYVE))

		Expect(staker.Address).To(Equal(i.BOB))
		Expect(staker.Amount).To(Equal(150 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(BeEmpty())
	})
})
