package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	querytypes "github.com/KYVENetwork/chain/x/query/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bundles module query tests", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Can vote should fail if pool does not exist", func() {
		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("pool with id 0 does not exist: not found"))
	})

	It("Can vote should fail if pool is upgrading", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:     "Moontest",
			Protocol: &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{
				Version:     "1.0.0",
				Binaries:    "{}",
				ScheduledAt: 100,
				Duration:    3600,
			},
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("pool currently upgrading"))
	})

	It("Can vote should fail if pool is paused", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			Paused:      true,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("pool is paused"))
	})

	It("Can vote should fail if pool is out of funds", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("pool is out of funds"))
	})

	It("Can vote should fail if pool has not reached min stake", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			MinStake:    50 * i.KYVE,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
			Amount:  100 * i.KYVE,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("min stake not reached"))
	})

	It("Can vote should fail if valaccount is not authorized", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			MinStake:    50 * i.KYVE,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.CHARLIE,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("valaccount not authorized"))
	})

	It("Can vote should fail if bundle was dropped", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			MinStake:    50 * i.KYVE,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.App().BundlesKeeper.SetBundleProposal(s.Ctx(), bundletypes.BundleProposal{
			PoolId:    0,
			StorageId: "",
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("can not vote on dropped bundle"))
	})

	It("Can vote should fail if bundle is of type no data bundle", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			MinStake:    50 * i.KYVE,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.App().BundlesKeeper.SetBundleProposal(s.Ctx(), bundletypes.BundleProposal{
			PoolId:    0,
			StorageId: "KYVE_NO_DATA_BUNDLE_0_1661239718",
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "KYVE_NO_DATA_BUNDLE_0_1661239718",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("can not vote on KYVE_NO_DATA_BUNDLE"))
	})

	It("Can vote should fail if storage id differs", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:        "Moontest",
			MinStake:    50 * i.KYVE,
			Protocol:    &pooltypes.Protocol{},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		s.App().BundlesKeeper.SetBundleProposal(s.Ctx(), bundletypes.BundleProposal{
			PoolId:    0,
			StorageId: "test_storage_id",
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.ALICE,
			Voter:     i.BOB,
			StorageId: "another_test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("provided storage_id does not match current one"))
	})

	It("Can vote should fail if voter has already voted valid", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       50 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
			Amount:     0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker:  i.ALICE,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 0,
			ToHeight:   100,
			FromKey:    "0",
			ToKey:      "99",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "test_storage_id",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.BOB,
			Voter:     i.ALICE,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("has already voted valid"))
	})

	It("Can vote should fail if voter has already voted invalid", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       50 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
			Amount:     0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker:  i.ALICE,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 0,
			ToHeight:   100,
			FromKey:    "0",
			ToKey:      "99",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "test_storage_id",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.BOB,
			Voter:     i.ALICE,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeFalse())
		Expect(canVote.Reason).To(Equal("has already voted invalid"))
	})

	It("Can vote should be possible if voter voted abstain already", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       50 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
			Amount:     0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker:  i.ALICE,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 0,
			ToHeight:   100,
			FromKey:    "0",
			ToKey:      "99",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "test_storage_id",
			Vote:      bundletypes.VOTE_TYPE_ABSTAIN,
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.BOB,
			Voter:     i.ALICE,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeTrue())
		Expect(canVote.Reason).To(Equal("KYVE_VOTE_NO_ABSTAIN_ALLOWED"))
	})

	It("Can vote should be possible", func() {
		// ARRANGE
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       50 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id:      0,
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
			Amount:     0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
			Amount:     0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker:  i.ALICE,
			PoolId:  0,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 0,
			ToHeight:   100,
			FromKey:    "0",
			ToKey:      "99",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		// ACT
		canVote, err := s.App().QueryKeeper.CanVote(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanVoteRequest{
			PoolId:    0,
			Staker:    i.BOB,
			Voter:     i.ALICE,
			StorageId: "test_storage_id",
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canVote.Possible).To(BeTrue())
		Expect(canVote.Reason).To(BeEmpty())
	})
})
