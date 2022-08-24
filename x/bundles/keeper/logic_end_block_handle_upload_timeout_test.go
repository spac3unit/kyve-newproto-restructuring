package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bundles module timeout tests", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create clean pool for every test case
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MaxBundleSize:  100,
			StartKey:       "0",
			MinStake:       100 * i.KYVE,
			UploadInterval: 60,
			OperatingCost:  10_000,
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
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("First staker who joins gets automatically chosen as next uploader", func() {
		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader gets removed due to pool upgrading", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		pool.UpgradePlan = &pooltypes.UpgradePlan{
			Version:     "1.0.0",
			Binaries:    "{}",
			ScheduledAt: uint64(s.Ctx().BlockTime().Unix()),
			Duration:    3600,
		}

		s.App().PoolKeeper.SetPool(s.Ctx(), pool)

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader gets removed due to pool being paused", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		pool.Paused = true

		s.App().PoolKeeper.SetPool(s.Ctx(), pool)

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader gets removed due to pool having no funds", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.RunTxPoolSuccess(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Id:      0,
			Amount:  100 * i.KYVE,
		})

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader gets removed due to pool not reaching min stake", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  50 * i.KYVE,
		})

		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(50 * i.KYVE))
	})

	It("Staker is just next uploader and upload timeout does not pass", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Staker is just next uploader and upload timeout does not pass but upload interval passed", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		// ACT
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Staker is just next uploader and upload timeout does pass together with upload interval", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		// ACT
		s.CommitAfterSeconds(s.App().BundlesKeeper.UploadTimeout(s.Ctx()))
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(BeEmpty())

		// check if next uploader got removed from pool
		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(BeEmpty())

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())

		Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(BeZero())

		// check if next uploader got slashed
		slashAmountRatio, _ := sdk.NewDecFromStr(s.App().StakersKeeper.TimeoutSlash(s.Ctx()))
		expectedBalance := 100*i.KYVE - uint64(sdk.NewDec(int64(100*i.KYVE)).Mul(slashAmountRatio).RoundInt64())

		Expect(expectedBalance).To(Equal(nextUploader.Amount))
	})

	It("Next uploader is still set after not reaching the upload interval", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader is still set after not reaching the upload timeout", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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

		// ACT
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader receives timeout slash and gets removed after not uploading on bundle which reached quorum", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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

		// ACT
		s.CommitAfterSeconds(s.App().BundlesKeeper.UploadTimeout(s.Ctx()))
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(BeEmpty())
		Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))

		// check if next uploader got removed from pool
		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(BeEmpty())

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())

		Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(BeZero())

		// check if next uploader got slashed
		slashAmountRatio, _ := sdk.NewDecFromStr(s.App().StakersKeeper.TimeoutSlash(s.Ctx()))
		expectedBalance := 100*i.KYVE - uint64(sdk.NewDec(int64(100*i.KYVE)).Mul(slashAmountRatio).RoundInt64())

		Expect(expectedBalance).To(Equal(nextUploader.Amount))
	})

	It("Bundle did not get dropped after not reaching the upload interval with a bundle with no quorum", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(2))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Bundle did get dropped after reaching the upload interval with a bundle with no quorum", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ACT
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)

		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.ByteSize).To(BeZero())
		Expect(bundleProposal.ToHeight).To(BeZero())
		Expect(bundleProposal.ToKey).To(BeEmpty())
		Expect(bundleProposal.ToValue).To(BeEmpty())
		Expect(bundleProposal.BundleHash).To(BeEmpty())
		Expect(bundleProposal.VotersValid).To(BeEmpty())
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(2))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader is still set after not reaching the upload interval with a KYVE_NO_DATA_BUNDLE", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
			PoolId:     0,
			StorageId:  "KYVE_NO_DATA_BUNDLE_0_1661331692",
			ByteSize:   0,
			FromHeight: 0,
			ToHeight:   0,
			FromKey:    "",
			ToKey:      "",
			ToValue:    "",
			BundleHash: "",
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(Equal("KYVE_NO_DATA_BUNDLE_0_1661331692"))

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(2))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader is still set after not reaching the upload interval with a KYVE_NO_DATA_BUNDLE", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
			PoolId:     0,
			StorageId:  "KYVE_NO_DATA_BUNDLE_0_1661331692",
			ByteSize:   0,
			FromHeight: 0,
			ToHeight:   0,
			FromKey:    "",
			ToKey:      "",
			ToValue:    "",
			BundleHash: "",
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ACT
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.StorageId).To(Equal("KYVE_NO_DATA_BUNDLE_0_1661331692"))

		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(2))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())
		Expect(nextUploader.Amount).To(Equal(100 * i.KYVE))
	})

	It("Next uploader is still set after not reaching the upload interval with a KYVE_NO_DATA_BUNDLE", func() {
		// ARRANGE
		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
			PoolId:     0,
			StorageId:  "KYVE_NO_DATA_BUNDLE_0_1661331692",
			ByteSize:   0,
			FromHeight: 0,
			ToHeight:   0,
			FromKey:    "",
			ToKey:      "",
			ToValue:    "",
			BundleHash: "",
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ACT
		s.CommitAfterSeconds(s.App().BundlesKeeper.UploadTimeout(s.Ctx()))
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		bundleProposal, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_1))
		Expect(bundleProposal.StorageId).To(Equal("KYVE_NO_DATA_BUNDLE_0_1661331692"))

		// check if next uploader got removed from pool
		poolStakers := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
		Expect(poolStakers).To(HaveLen(1))

		nextUploader, found := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		Expect(found).To(BeTrue())

		Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(Equal(100 * i.KYVE))

		// check if next uploader got slashed
		slashAmountRatio, _ := sdk.NewDecFromStr(s.App().StakersKeeper.TimeoutSlash(s.Ctx()))
		expectedBalance := 100*i.KYVE - uint64(sdk.NewDec(int64(100*i.KYVE)).Mul(slashAmountRatio).RoundInt64())

		Expect(expectedBalance).To(Equal(nextUploader.Amount))
	})
})
