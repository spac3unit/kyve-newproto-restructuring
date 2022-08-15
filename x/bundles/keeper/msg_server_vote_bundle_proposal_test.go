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
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:          "Moontest",
			MaxBundleSize: 100,
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.ALICE,
			PoolId:     0,
			Valaddress: i.BOB,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker:  i.ALICE,
			PoolId:  0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize:   100,
			FromHeight: 0,
			ToHeight:   100,
			FromKey:    "0",
			ToKey:      "99",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyStakersModuleAssetsIntegrity()
	})

	It("Try to vote valid on proposal", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote valid on proposal again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote invalid on proposal", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote invalid on proposal again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote abstain on proposal", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).To(ContainElement(i.BOB))
	})

	It("Try to vote abstain on proposal again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).To(ContainElement(i.BOB))
	})

	It("Try to vote valid on proposal after abstain vote", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote invalid on proposal after abstain vote", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote unspecified on proposal", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_UNSPECIFIED,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	It("Try to vote on proposal with invalid storage id", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		// ACT
		s.RunTxBundlesError(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "some_invalid_storage_id",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.VotersValid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).NotTo(ContainElement(i.BOB))
		Expect(bundleProposal.VotersAbstain).NotTo(ContainElement(i.BOB))
	})

	// TODO: try to vote on KYVE_NO_DATA_BUNDLE
})
