package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Undelegation", Ordered, func() {
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

	It("Undelegate 11 KYVE from ALICE with DUMMY_0 (more than allowed)", func() {
		res, err := s.RunTxDelegator(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  11 * i.KYVE,
		})

		fmt.Println(res)
		fmt.Println(err)

		Expect(res).To(BeNil())
		Expect(err).ToNot(BeNil())
	})

	It("Undelegate 5 KYVE from ALICE with DUMMY_0", func() {

		delegationDummyBefore := s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])

		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})

		delegationDummyAfter := s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])

		//Delegation amount stays the same (due to unbonding)
		Expect(delegationDummyBefore).To(Equal(delegationDummyAfter))
		Expect(delegationDummyAfter).To(Equal(10 * i.KYVE))

		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		delegationDummyAfter = s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])
		Expect(delegationDummyAfter).To(Equal(5 * i.KYVE))

	})

	It("Undelegate Slashed Amount", func() {

		delegationDummyBefore := s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])

		fraction, err := sdk.NewDecFromStr("0.1")
		Expect(err).To(BeNil())
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, fraction)

		delegationDummyAfter := s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])

		slashes := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
		_ = slashes

		fmt.Println("Slash")
		fmt.Println(delegationDummyBefore)
		fmt.Println(delegationDummyAfter)

		//
		////Delegation amount stays the same (due to unbonding)
		//Expect(delegationDummyBefore).To(Equal(delegationDummyAfter))
		//Expect(delegationDummyAfter).To(Equal(10 * i.KYVE))
		//
		//s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		//s.CommitAfterSeconds(1)
		//
		//delegationDummyAfter = s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])
		//Expect(delegationDummyAfter).To(Equal(5 * i.KYVE))

	})

})
