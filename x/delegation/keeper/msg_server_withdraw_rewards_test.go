package keeper_test

import (
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
)

var _ = Describe("Delegation", Ordered, func() {
	s := i.NewCleanChain()

	BeforeAll(func() {
		initPoolWithStakersAliceAndBob(&s)
	})

	It("Create three delegators", func() {

		dummyBalance0 := s.GetBalanceFromAddress(i.DUMMY[0])
		dummyBalance1 := s.GetBalanceFromAddress(i.DUMMY[1])
		dummyBalance2 := s.GetBalanceFromAddress(i.DUMMY[2])

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[2],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		dummyBalance0After := s.GetBalanceFromAddress(i.DUMMY[0])
		dummyBalance1After := s.GetBalanceFromAddress(i.DUMMY[1])
		dummyBalance2After := s.GetBalanceFromAddress(i.DUMMY[2])

		Expect(dummyBalance0After).To(Equal(dummyBalance0 - 10*i.KYVE))
		Expect(dummyBalance1After).To(Equal(dummyBalance1 - 10*i.KYVE))
		Expect(dummyBalance2After).To(Equal(dummyBalance2 - 10*i.KYVE))

		aliceDelegation := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegation).To(Equal(30 * i.KYVE))
	})

	It("Split rewards; test rounding", func() {

		delegationModuleBalanceBefore := s.GetBalanceFromModule(types.ModuleName)
		poolModuleBalanceBefore := s.GetBalanceFromModule(pooltypes.ModuleName)

		success := s.App().DelegationKeeper.PayoutRewards(s.Ctx(), i.ALICE, 20*i.KYVE, pooltypes.ModuleName)
		Expect(success).To(BeTrue())

		delegationModuleBalanceAfter := s.GetBalanceFromModule(types.ModuleName)
		poolModuleBalanceAfter := s.GetBalanceFromModule(pooltypes.ModuleName)

		Expect(delegationModuleBalanceAfter).To(Equal(delegationModuleBalanceBefore + 20*i.KYVE))
		Expect(poolModuleBalanceAfter).To(Equal(poolModuleBalanceBefore - 20*i.KYVE))

		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(6666666666)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(6666666666)))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[2])).To(Equal(uint64(6666666666)))

		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
		})
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
		})
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[2],
			Staker:  i.ALICE,
		})

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(uint64(996666666666)))
		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(uint64(996666666666)))
		Expect(s.GetBalanceFromAddress(i.DUMMY[2])).To(Equal(uint64(996666666666)))

		Expect(s.GetBalanceFromModule(types.ModuleName)).To(Equal(uint64(30000000002)))
	})

	It("Withdraw from not delegator", func() {

		balanceDummy1Before := s.GetBalanceFromAddress(i.DUMMY[0])
		balanceCharlieBefore := s.GetBalanceFromAddress(i.CHARLIE)
		balanceAliceBefore := s.GetBalanceFromAddress(i.ALICE)
		delegationBalance := s.GetBalanceFromModule(types.ModuleName)

		s.RunTxDelegatorError(&types.MsgWithdrawRewards{
			Creator: i.CHARLIE,
			Staker:  i.ALICE,
		})

		s.RunTxDelegatorError(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[0],
			Staker:  i.CHARLIE,
		})

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(balanceDummy1Before))
		Expect(s.GetBalanceFromAddress(i.CHARLIE)).To(Equal(balanceCharlieBefore))
		Expect(s.GetBalanceFromAddress(i.ALICE)).To(Equal(balanceAliceBefore))
		Expect(s.GetBalanceFromModule(types.ModuleName)).To(Equal(delegationBalance))

	})

	It("Test invalid Payout", func() {
		success := s.App().DelegationKeeper.PayoutRewards(s.Ctx(), i.ALICE, 20000*i.KYVE, pooltypes.ModuleName)
		Expect(success).To(BeFalse())

		success = s.App().DelegationKeeper.PayoutRewards(s.Ctx(), i.DUMMY[20], 0*i.KYVE, pooltypes.ModuleName)
		Expect(success).To(BeFalse())
	})

	// TODO test withdraw with multiple slashes

})
