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

	It("Stake with more than available balance", func() {
		currentBalance := s.GetBalanceFromAddress(i.ALICE)

		s.RunTxStakersError(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  currentBalance + 1,
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

	// TODO: test updating moniker, logo and website
	// TODO: test kicking out lowest staker
})
