package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/stakers/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Unstaking", Ordered, func() {
	s := i.NewCleanChain()

	initialBalance := s.GetBalanceFromAddress(i.ALICE)

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
		s.VerifyStakersModuleAssetsIntegrity()
		s.VerifyPoolTotalStake()
		s.VerifyStakersQueries()
	})

	It("Unstake 50 KYVE", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  50 * i.KYVE,
		})

		// ASSERT
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(unstakingEntries).To(HaveLen(1))

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(Equal(50 * i.KYVE))
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter = s.GetBalanceFromAddress(i.ALICE)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)

		Expect(unstakingEntries).To(BeEmpty())

		Expect(initialBalance - balanceAfter).To(Equal(50 * i.KYVE))

		Expect(staker.Amount).To(Equal(50 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
	})

	It("Unstake more than staked", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  101 * i.KYVE,
		})

		// ASSERT
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())
	})

	It("Unstake everything", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		// ASSERT
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(unstakingEntries).To(HaveLen(1))

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(Equal(100 * i.KYVE))
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter = s.GetBalanceFromAddress(i.ALICE)

		_, found = s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)

		Expect(unstakingEntries).To(BeEmpty())

		Expect(initialBalance - balanceAfter).To(BeZero())

		Expect(found).To(BeFalse())
	})

	It("Unstake while already unbonding", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  25 * i.KYVE,
		})

		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  25 * i.KYVE,
		})

		// ASSERT
		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(2))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		Expect(found).To(BeTrue())
		Expect(initialBalance - balanceAfter).To(Equal(50 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(50 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Unstake more than staked while already unbonding", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  25 * i.KYVE,
		})

		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  90 * i.KYVE,
		})

		// ASSERT
		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)
		balanceAfter := s.GetBalanceFromAddress(i.ALICE)

		Expect(found).To(BeTrue())
		Expect(initialBalance - balanceAfter).To(Equal(75 * i.KYVE))

		Expect(staker.Address).To(Equal(i.ALICE))
		Expect(staker.Amount).To(Equal(75 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})
})
