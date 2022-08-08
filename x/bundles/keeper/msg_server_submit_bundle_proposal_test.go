package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Submit Bundle Proposal", Ordered, func() {
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
			MaxBundleSize: 100,
			StartKey: "0",
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Amount: 100*i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.ALICE,
			Amount: 100*i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 0,
			Valaddress: i.BOB,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
		})
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyStakersModuleAssetsIntegrity()
	})

	It("Try to submit proposal", func () {
		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, found := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(found).To(BeTrue())

		Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
		Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))
		Expect(bundleProposal.Uploader).To(Equal(i.ALICE))
		Expect(bundleProposal.NextUploader).To(Equal(i.ALICE))
		Expect(bundleProposal.ByteSize).To(Equal(uint64(100)))
		Expect(bundleProposal.ToHeight).To(Equal(uint64(100)))
		Expect(bundleProposal.ToKey).To(Equal("99"))
		Expect(bundleProposal.ToValue).To(Equal("test_value"))
		Expect(bundleProposal.BundleHash).To(Equal("test_hash"))
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(ContainElement(i.ALICE))
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())
	})

	It("Try to submit proposal with non existing valaccount", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.CHARLIE,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.CHARLIE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with wrong valaccount", func () {
		// ARRANGE
		s.RunTxPoolSuccess(&pooltypes.MsgCreatePool{
			Creator: i.ALICE,
			Name:    "Moontest2",
			Config: "{}",
			Binaries: "{}",
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator: i.ALICE,
			PoolId: 1,
			Valaddress: i.CHARLIE,
		})
		
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.CHARLIE,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with empty storage id", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with empty byte size", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 0,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with invalid from height", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 1,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with bigger bundle size than allowed", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 101,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with empty bundle", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 0,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with empty value", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "",
			BundleHash: "test_hash",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	It("Try to submit proposal with empty bundle hash", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgSubmitBundleProposal{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			ByteSize: 100,
			FromHeight: 0,
			ToHeight: 100,
			FromKey: "0",
			ToKey: "99",
			ToValue: "test_value",
			BundleHash: "",
		})

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
	})

	// TODO: submit KYVE_NO_DATA_BUNDLES
	// TODO: submit bundle proposal without reaching upload interval
})
