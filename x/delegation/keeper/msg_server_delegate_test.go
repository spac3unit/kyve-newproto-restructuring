package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Delegation", Ordered, func() {
	s := i.NewCleanChain()

	BeforeAll(func() {
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.CommitAfterSeconds(7)

		_, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		_, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(stakerFound).To(BeTrue())

		// Create Staker
		s.RunTxStakersSuccess(&stakerstypes.MsgUnstake{
			Creator: i.ALICE,
			Amount:  50 * i.KYVE,
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
			Amount:  50,
		})

		s.CommitAfterSeconds(7)
	})

	It("Delegate 10 KYVE to ALICE", func() {

		bobBalance := s.GetBalanceFromAddress(i.BOB)

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.BOB,
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)

		Expect(bobBalanceAfter).To(Equal(bobBalance - 10*i.KYVE))

		aliceDelegation := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)

		Expect(aliceDelegation).To(Equal(10 * i.KYVE))
	})

	It("Delegate more than available", func() {

		bobBalance := s.GetBalanceFromAddress(i.BOB)
		aliceDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)

		_, delegateErr := s.RunTxDelegator(&types.MsgDelegate{
			Creator: i.BOB,
			Staker:  i.ALICE,
			Amount:  2000 * i.KYVE,
		})

		Expect(delegateErr).ToNot(BeNil())

		aliceDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegationBefore).To(Equal(aliceDelegationAfter))

		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)
		Expect(bobBalanceAfter).To(Equal(bobBalance))
	})

	It("Payout delegators", func() {
		s.App().DelegationKeeper.PayoutRewards(s.Ctx(), i.ALICE, 1*i.KYVE, pooltypes.ModuleName)

		outstandingRewardsBefore := s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.BOB)
		Expect(outstandingRewardsBefore).To(Equal(1 * i.KYVE))

		bobBalanceBefore := s.GetBalanceFromAddress(i.BOB)
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.BOB,
			Staker:  i.ALICE,
		})
		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)

		Expect(bobBalanceAfter).To(Equal(bobBalanceBefore + 1*i.KYVE))

		outstandingRewardsAfter := s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.BOB)
		Expect(outstandingRewardsAfter).To(Equal(0 * i.TKYVE))
	})

	It("Don't pay out rewards twice", func() {
		outstandingRewardsBefore := s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.BOB)

		bobBalanceBefore := s.GetBalanceFromAddress(i.BOB)
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.BOB,
			Staker:  i.ALICE,
		})
		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)

		Expect(bobBalanceAfter).To(Equal(bobBalanceBefore))

		outstandingRewardsAfter := s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.BOB)
		Expect(outstandingRewardsAfter).To(Equal(outstandingRewardsBefore))
	})

})
