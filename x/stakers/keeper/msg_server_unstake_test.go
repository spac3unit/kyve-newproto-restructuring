package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/stakers/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

/*

TEST CASES - msg_server_unstake.go

* Unstake 50 KYVE from a staker who has previously staked 100 KYVE
* Try to unstake more KYVE than the staker has actually staked
* Unstake full staking amount from a staker who has previously staked 100 KYVE
* Unstake 25 KYVE while staker is already unbonding 25 KYVE from 100 KYVE in total
* Unstake 90 KYVE while staker is already unbonding 25 KYVE although staker only staked 100 KYVE in total
* TODO: full unstake with delegators

*/

var _ = Describe("msg_server_unstake.go", Ordered, func() {
	s := i.NewCleanChain()

	initialBalance := s.GetBalanceFromAddress(i.STAKER_0)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create staker
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Unstake 50 KYVE from a staker who has previously staked 100 KYVE", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  50 * i.KYVE,
		})

		// ASSERT
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(unstakingEntries).To(HaveLen(1))

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(Equal(50 * i.KYVE))
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))

		s.PerformValidityChecks()

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter = s.GetBalanceFromAddress(i.STAKER_0)

		staker, _ = s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)

		Expect(unstakingEntries).To(BeEmpty())

		Expect(initialBalance - balanceAfter).To(Equal(50 * i.KYVE))

		Expect(staker.Amount).To(Equal(50 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
	})

	It("Try to unstake more KYVE than the staker has actually staked", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  101 * i.KYVE,
		})

		// ASSERT
		s.PerformValidityChecks()
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())
	})

	It("Unstake full staking amount from a staker who has previously staked 100 KYVE", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		// ASSERT
		s.PerformValidityChecks()
		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(unstakingEntries).To(HaveLen(1))

		Expect(found).To(BeTrue())

		Expect(initialBalance - balanceAfter).To(Equal(100 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(100 * i.KYVE))
		Expect(staker.UnbondingAmount).To(Equal(100 * i.KYVE))
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
		s.PerformValidityChecks()

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		balanceAfter = s.GetBalanceFromAddress(i.STAKER_0)

		_, found = s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)

		Expect(unstakingEntries).To(BeEmpty())

		Expect(initialBalance - balanceAfter).To(BeZero())

		Expect(found).To(BeFalse())
	})

	It("Unstake 25 KYVE while staker is already unbonding 25 KYVE from 100 KYVE in total", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  25 * i.KYVE,
		})

		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  25 * i.KYVE,
		})

		// ASSERT
		s.PerformValidityChecks()
		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(2))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		Expect(found).To(BeTrue())
		Expect(initialBalance - balanceAfter).To(Equal(50 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(50 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})

	It("Unstake 90 KYVE while staker is already unbonding 25 KYVE although staker only staked 100 KYVE in total", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  25 * i.KYVE,
		})

		unstakingEntries := s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  90 * i.KYVE,
		})

		// ASSERT
		s.PerformValidityChecks()
		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(HaveLen(1))

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		unstakingEntries = s.App().StakersKeeper.GetAllUnbondingStakeEntries(s.Ctx())
		Expect(unstakingEntries).To(BeEmpty())

		staker, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)
		balanceAfter := s.GetBalanceFromAddress(i.STAKER_0)

		Expect(found).To(BeTrue())
		Expect(initialBalance - balanceAfter).To(Equal(75 * i.KYVE))

		Expect(staker.Address).To(Equal(i.STAKER_0))
		Expect(staker.Amount).To(Equal(75 * i.KYVE))
		Expect(staker.UnbondingAmount).To(BeZero())
		Expect(staker.Commission).To(Equal(types.DefaultCommission))

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())

		Expect(valaccounts).To(HaveLen(0))
	})
})
