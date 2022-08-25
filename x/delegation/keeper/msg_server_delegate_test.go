package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

/*

TEST CASES - msg_server_delegate_test.go

* Delegate 10 KYVE to ALICE
* Try delegating to non-existent staker
* Delegate more than available
* Payout delegators
* Don't pay out rewards twice

*/

var _ = Describe("Delegation - Delegation", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		CreateFundedPool(&s)

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.BOB,
			Amount:  200 * i.KYVE,
		})

		_, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(stakerFound).To(BeTrue())

		s.CommitAfterSeconds(7)
	})

	AfterEach(func() {
		CheckAndContinueChainForOneMonth(&s)
	})

	It("Delegate 10 KYVE to ALICE", func() {

		// Arrange
		bobBalance := s.GetBalanceFromAddress(i.BOB)

		// Act
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.BOB,
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		// Assert
		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)
		Expect(bobBalanceAfter).To(Equal(bobBalance - 10*i.KYVE))

		aliceDelegation := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegation).To(Equal(10 * i.KYVE))
	})

	It("Try delegating to non-existent staker", func() {

		// Arrange
		bobBalance := s.GetBalanceFromAddress(i.BOB)

		// Act
		s.RunTxDelegatorError(&types.MsgDelegate{
			Creator: i.BOB,
			Staker:  i.CHARLIE,
			Amount:  10 * i.KYVE,
		})

		// Assert
		Expect(s.GetBalanceFromAddress(i.BOB)).To(Equal(bobBalance))

		aliceDelegation := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegation).To(BeZero())
	})

	It("Delegate more than available", func() {

		// Arrange
		bobBalance := s.GetBalanceFromAddress(i.BOB)
		aliceDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)

		// Act
		_, delegateErr := s.RunTxDelegator(&types.MsgDelegate{
			Creator: i.BOB,
			Staker:  i.ALICE,
			Amount:  2000 * i.KYVE,
		})

		// Assert
		Expect(delegateErr).ToNot(BeNil())

		aliceDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegationBefore).To(Equal(aliceDelegationAfter))

		bobBalanceAfter := s.GetBalanceFromAddress(i.BOB)
		Expect(bobBalanceAfter).To(Equal(bobBalance))
	})

	It("Payout delegators", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  100 * i.KYVE,
		})
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  200 * i.KYVE,
		})
		poolModuleBalance := s.GetBalanceFromModule(pooltypes.ModuleName)
		Expect(poolModuleBalance).To(Equal(50 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(BeZero())
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(BeZero())

		// Act
		PayoutRewards(&s, i.ALICE, 10*i.KYVE)

		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(3_333_333_333)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(6_666_666_666)))

		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
		})

		// Assert
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(0)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(6_666_666_666)))

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(uint64(903333333333)))
		Expect(s.GetBalanceFromModule(pooltypes.ModuleName), 40*i.KYVE)
		Expect(s.GetBalanceFromModule(types.ModuleName), uint64(6_666_666_666+1))
	})

	It("Don't pay out rewards twice", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  100 * i.KYVE,
		})
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  200 * i.KYVE,
		})
		poolModuleBalance := s.GetBalanceFromModule(pooltypes.ModuleName)
		Expect(poolModuleBalance).To(Equal(50 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(BeZero())
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(BeZero())

		PayoutRewards(&s, i.ALICE, 10*i.KYVE)

		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(3_333_333_333)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(6_666_666_666)))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
		})
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
		})

		// Assert
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(0)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(6_666_666_666)))

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(uint64(903333333333)))
		Expect(s.GetBalanceFromModule(pooltypes.ModuleName), 40*i.KYVE)
		Expect(s.GetBalanceFromModule(types.ModuleName), uint64(6_666_666_666+1))
	})

})
