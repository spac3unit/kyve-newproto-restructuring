package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Update Metadata", Ordered, func() {
	s := i.NewCleanChain()

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
		s.PerformValidityChecks()
	})

	It("Get default metadata", func() {
		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
	})

	It("Update metadata", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateMetadata{
			Creator: i.ALICE,
			Moniker: "KYVE Node Runner",
			Website: "https://kyve.network",
			Logo:    "https://arweave.net/Tewyv2P5VEG8EJ6AUQORdqNTectY9hlOrWPK8wwo-aU",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)

		Expect(staker.Moniker).To(Equal("KYVE Node Runner"))
		Expect(staker.Website).To(Equal("https://kyve.network"))
		Expect(staker.Logo).To(Equal("https://arweave.net/Tewyv2P5VEG8EJ6AUQORdqNTectY9hlOrWPK8wwo-aU"))
	})

	It("Reset metadata", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgUpdateMetadata{
			Creator: i.ALICE,
			Moniker: "",
			Website: "",
			Logo:    "",
		})

		// ASSERT
		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)

		Expect(staker.Moniker).To(BeEmpty())
		Expect(staker.Website).To(BeEmpty())
		Expect(staker.Logo).To(BeEmpty())
	})
})
