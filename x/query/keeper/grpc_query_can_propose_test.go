package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	querytypes "github.com/KYVENetwork/chain/x/query/types"
	"github.com/KYVENetwork/chain/x/registry/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Can Propose Tests", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       200 * i.KYVE,
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
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     0,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
			Amount:     0,
		})

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
			Creator:   i.VALADDRESS_1,
			Staker:    i.STAKER_1,
			PoolId:    0,
			StorageId: "test_storage_id",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		s.CommitAfterSeconds(60)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Can propose should fail if pool does not exist", func() {
		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     1,
			Staker:     i.STAKER_1,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), 1).Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_1,
			PoolId:     1,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	It("Can propose should fail if pool is upgrading", func() {
		// ARRANGE
		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		pool.UpgradePlan = &pooltypes.UpgradePlan{
			Version:     "1.0.0",
			Binaries:    "{}",
			ScheduledAt: 100,
			Duration:    3600,
		}

		s.App().PoolKeeper.SetPool(s.Ctx(), pool)

		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     0,
			Staker:     i.STAKER_1,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(bundletypes.ErrPoolCurrentlyUpgrading.Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_1,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	It("Can propose should fail if pool is paused", func() {
		// ARRANGE
		pool, _ := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		pool.Paused = true

		s.App().PoolKeeper.SetPool(s.Ctx(), pool)

		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     0,
			Staker:     i.STAKER_1,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(bundletypes.ErrPoolPaused.Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_1,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	It("Can propose should fail if pool is out of funds", func() {
		// ARRANGE
		s.RunTxPoolSuccess(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Id:      0,
			Amount:  100 * i.KYVE,
		})

		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     0,
			Staker:     i.STAKER_1,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(bundletypes.ErrPoolOutOfFunds.Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_1,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	It("Can propose should fail if pool has not reached min stake", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgUnstake{
			Creator: i.STAKER_0,
			Amount:  50 * i.KYVE,
		})

		// wait for unbonding
		s.CommitAfterSeconds(s.App().StakersKeeper.UnbondingStakingTime(s.Ctx()))
		s.CommitAfterSeconds(1)

		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     0,
			Staker:     i.STAKER_1,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(bundletypes.ErrMinStakeNotReached.Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_1,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	It("Can propose should fail if valaccount is not authorized", func() {
		// ACT
		canPropose, err := s.App().QueryKeeper.CanPropose(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanProposeRequest{
			PoolId:     0,
			Staker:     i.STAKER_0,
			Proposer:   i.VALADDRESS_1,
			FromHeight: 100,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canPropose.Possible).To(BeFalse())
		Expect(canPropose.Reason).To(Equal(stakertypes.ErrValaccountUnauthorized.Error()))

		_, txErr := s.RunTxBundles(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_1,
			Staker:     i.STAKER_0,
			PoolId:     0,
			StorageId:  "test_storage_id",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value",
			BundleHash: "test_hash",
		})

		Expect(txErr).NotTo(BeNil())
		Expect(txErr.Error()).To(Equal(canPropose.Reason))
	})

	// TODO: not designated uploader
	// TODO: upload interval not reached
	// TODO: invalid from_height
	// TODO: can actually propose
})
