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
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
			Amount:  50,
		})

		s.CommitAfterSeconds(7)

		_, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		_, aliceFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(aliceFound).To(BeTrue())

		_, bobFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.BOB)
		Expect(bobFound).To(BeTrue())

		s.CommitAfterSeconds(7)
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

		println("PArams:")
		println(s.App().StakersKeeper.UploadSlash(s.Ctx()))
		println(s.App().DelegationKeeper.RedelegationMaxAmount(s.Ctx()))

		aliceDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		bobDelegationBefore := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)

		s.RunTxDelegatorSuccess(&types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.ALICE,
			ToStaker:   i.BOB,
			Amount:     1 * i.KYVE,
		})

		aliceDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)
		Expect(aliceDelegationBefore).To(Equal(aliceDelegationAfter + 1*i.KYVE))

		bobDelegationAfter := s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.BOB)
		Expect(bobDelegationBefore).To(Equal(bobDelegationAfter - 1*i.KYVE))
	})

})
