package keeper_test

import (
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

	It("Delegate 10 KYVE to ALICE from DUMMY_0", func() {
		dummy0BalanceBefore := s.GetBalanceFromAddress(i.DUMMY[0])

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		dummy0BalanceAfter := s.GetBalanceFromAddress(i.DUMMY[0])
		Expect(dummy0BalanceBefore).To(Equal(dummy0BalanceAfter + 10*i.KYVE))
	})

	It("Redelegate 1 KYVE to Bob", func() {

		aliceDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		bobDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)

		s.RunTxDelegatorSuccess(&types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.ALICE,
			ToStaker:   i.BOB,
			Amount:     1 * i.KYVE,
		})
		s.CommitAfterSeconds(10)

		aliceDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegationBefore).To(Equal(aliceDelegationAfter + 1*i.KYVE))

		bobDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)
		Expect(bobDelegationBefore).To(Equal(bobDelegationAfter - 1*i.KYVE))
	})

	It("Redelegate without delegation", func() {

		charlieDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.CHARLIE)
		bobDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)

		s.RunTxDelegatorError(&types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.CHARLIE,
			ToStaker:   i.BOB,
			Amount:     1 * i.KYVE,
		})

		charlieDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.CHARLIE)
		Expect(charlieDelegationBefore).To(Equal(charlieDelegationAfter))

		bobDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)
		Expect(bobDelegationBefore).To(Equal(bobDelegationAfter))
	})

	It("Exhaust all redelegation spells", func() {

		aliceDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		bobDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)

		redelegationMessage := types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.ALICE,
			ToStaker:   i.BOB,
			Amount:     1 * i.KYVE,
		}

		s.RunTxDelegatorSuccess(&redelegationMessage)
		s.CommitAfterSeconds(10)
		s.RunTxDelegatorSuccess(&redelegationMessage)
		s.CommitAfterSeconds(10)
		s.RunTxDelegatorSuccess(&redelegationMessage)
		s.CommitAfterSeconds(10)
		s.RunTxDelegatorSuccess(&redelegationMessage)
		s.CommitAfterSeconds(10)

		aliceDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegationBefore).To(Equal(aliceDelegationAfter + 4*i.KYVE))

		bobDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)
		Expect(bobDelegationBefore).To(Equal(bobDelegationAfter - 4*i.KYVE))

		// Now all redelegation spells are exhausted
		s.RunTxDelegatorError(&redelegationMessage)
	})

	It("Expire redelegation spells", func() {

		s.CommitAfterSeconds(s.App().DelegationKeeper.RedelegationCooldown(s.Ctx()) - 50)
		s.CommitAfterSeconds(1)

		// One redelegation available
		redelegationMessage := types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.ALICE,
			ToStaker:   i.BOB,
			Amount:     1 * i.KYVE,
		}

		s.RunTxDelegatorSuccess(&redelegationMessage)

		s.CommitAfterSeconds(1)

		// Redelegations are now all used again
		s.RunTxDelegatorError(&redelegationMessage)

		// Expire next two spells
		s.CommitAfterSeconds(25)

		s.RunTxDelegatorSuccess(&redelegationMessage)
		// No two delegation
		s.RunTxDelegatorError(&redelegationMessage)
	})

})
