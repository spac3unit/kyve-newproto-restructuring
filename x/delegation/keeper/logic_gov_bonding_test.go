package keeper_test

import (
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
)

var _ = Describe("Delegation Gov Logic", Ordered, func() {
	s := i.NewCleanChain()
	aliceAcc, _ := sdk.AccAddressFromBech32(i.ALICE)
	bobAcc, _ := sdk.AccAddressFromBech32(i.BOB)
	charlieAcc, _ := sdk.AccAddressFromBech32(i.CHARLIE)

	BeforeEach(func() {
		println("BEFORE EACH")
		s := i.NewCleanChain()

		Expect(s.App().DelegationKeeper.TotalProtocolBonding(s.Ctx())).To(Equal(sdk.NewInt(0)))
	})

	It("Simple Staking", func() {

		// Arrange
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		// Assert
		Expect(s.App().DelegationKeeper.TotalProtocolBonding(s.Ctx())).To(Equal(sdk.NewInt(int64(100 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), aliceAcc)).To(Equal(sdk.NewInt(int64(100 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), bobAcc)).To(Equal(sdk.NewInt(int64(0 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), charlieAcc)).To(Equal(sdk.NewInt(int64(0 * i.KYVE))))
	})

	It("Multiple Staking", func() {
		s := i.NewCleanChain()
		Expect(s.App().DelegationKeeper.TotalProtocolBonding(s.Ctx())).To(Equal(sdk.NewInt(0)))

		// Arrange
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.BOB,
			Amount:  50 * i.KYVE,
		})

		// Assert
		Expect(s.App().DelegationKeeper.TotalProtocolBonding(s.Ctx())).To(Equal(sdk.NewInt(int64(150 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), aliceAcc)).To(Equal(sdk.NewInt(int64(100 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), bobAcc)).To(Equal(sdk.NewInt(int64(50 * i.KYVE))))
		Expect(s.App().DelegationKeeper.GetBondingOfAddress(s.Ctx(), charlieAcc)).To(Equal(sdk.NewInt(int64(0 * i.KYVE))))
	})

	// TODO test with delegation
})
