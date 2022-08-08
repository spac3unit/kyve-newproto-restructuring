package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Vote Proposal", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create clean pool for every test case
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest",
			Config: "{}",
			Binaries: "{}",
		})
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyStakersModuleAssetsIntegrity()
	})

	It("Try to vote on proposal where pool is not funded", func () {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.ALICE,
			Amount: 100*i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote: bundletypes.VOTE_TYPE_YES,
		})

		// ASSERT
		_, found := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(found).To(BeFalse()) // TODO: implement

		// Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.ALICE))
		// Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.ALICE))
		// Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.ALICE))
	})

	// TODO: add test cases where pool is not created
})
