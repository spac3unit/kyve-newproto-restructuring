package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	querytypes "github.com/KYVENetwork/chain/x/query/types"
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
})
