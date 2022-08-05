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

	initialBalance := s.GetBalanceFromAddress(i.ALICE)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.VerifyStakersModuleAssetsIntegrity()
	})

	It("Create new staker with 100 KYVE", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100*i.KYVE,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100*i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(100*i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Stake additional 50 KYVE", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100*i.KYVE,
		})

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  50*i.KYVE,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(150*i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(150*i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Stake with more than available balance", func() {
		// ACT
		currentBalance := s.GetBalanceFromAddress(i.ALICE)

		s.RunTxStakersError(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  currentBalance + 1,
		})

		// ASSERT
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)
		Expect(initialBalance - balanceAfter).To(BeZero())

		_, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(found).To(BeFalse())
	})

	// TODO: test updating moniker, logo and website
	// TODO: test kicking out lowest staker
})
