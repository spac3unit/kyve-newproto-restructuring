package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: validator does not vote n times in a row

var _ = Describe("Bundles module integration tests", Ordered, func() {
	s := i.NewCleanChain()

	initialBalanceAlice := s.GetBalanceFromAddress(i.ALICE)
	initialBalanceBob := s.GetBalanceFromAddress(i.ALICE)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create clean pool for every test case
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MaxBundleSize:  100,
			StartKey:       "0",
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

		s.RunTxStakersSuccess(&stakertypes.MsgCreateStaker{
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

		initialBalanceAlice = s.GetBalanceFromAddress(i.ALICE)
		initialBalanceBob = s.GetBalanceFromAddress(i.BOB)

		s.CommitAfterSeconds(60)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Produce valid bundles with one node", func() {
		// ARRANGE
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

		s.CommitAfterSeconds(60)

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value2",
			BundleHash: "test_hash2",
		})

		// ASSERT
		// check if bundle got finalized on pool
		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.CurrentKey).To(Equal("99"))
		Expect(pool.CurrentValue).To(Equal("test_value"))
		Expect(pool.CurrentHeight).To(Equal(uint64(100)))
		Expect(pool.TotalBundles).To(Equal(uint64(1)))

		// check if finalized bundle got saved
		finalizedBundle, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
		Expect(finalizedBundleFound).To(BeTrue())

		Expect(finalizedBundle.PoolId).To(Equal(uint64(0)))
		Expect(finalizedBundle.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))
		Expect(finalizedBundle.Uploader).To(Equal(i.ALICE))
		Expect(finalizedBundle.FromHeight).To(Equal(uint64(0)))
		Expect(finalizedBundle.ToHeight).To(Equal(uint64(100)))
		Expect(finalizedBundle.Key).To(Equal("99"))
		Expect(finalizedBundle.Value).To(Equal("test_value"))
		Expect(finalizedBundle.BundleHash).To(Equal("test_hash"))
		Expect(finalizedBundle.FinalizedAt).NotTo(BeZero())

		// check if next bundle proposal got registered
		bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposalFound).To(BeTrue())

		Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
		Expect(bundleProposal.StorageId).To(Equal("P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg"))
		Expect(bundleProposal.Uploader).To(Equal(i.ALICE))
		Expect(bundleProposal.NextUploader).To(Equal(i.ALICE))
		Expect(bundleProposal.ByteSize).To(Equal(uint64(100)))
		Expect(bundleProposal.ToHeight).To(Equal(uint64(200)))
		Expect(bundleProposal.ToKey).To(Equal("199"))
		Expect(bundleProposal.ToValue).To(Equal("test_value2"))
		Expect(bundleProposal.BundleHash).To(Equal("test_hash2"))
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(ContainElement(i.ALICE))
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceBob))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		// calculate uploader reward
		totalReward := 100*s.App().BundlesKeeper.StorageCost(s.Ctx()) + pool.OperatingCost
		networkFee, _ := sdk.NewDecFromStr(s.App().BundlesKeeper.NetworkFee(s.Ctx()))

		treasuryReward := uint64(sdk.NewDec(int64(totalReward)).Mul(networkFee).RoundInt64())
		uploaderReward := totalReward - treasuryReward

		Expect(balanceStaker + s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.ALICE)).To(Equal(initialBalanceAlice + uploaderReward))

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		funder, _ := pool.GetFunder(i.ALICE)

		Expect(pool.Funders).To(HaveLen(1))
		Expect(funder.Amount).To(Equal(100*i.KYVE - totalReward))
	})

	PIt("Produce valid bundles with two nodes", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakertypes.MsgCreateStaker{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		initialBalanceBob = s.GetBalanceFromAddress(i.BOB)

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

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		s.CommitAfterSeconds(60)

		bundle, _ := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		_ = bundle

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.ALICE,
			Staker:     i.BOB,
			PoolId:     0,
			StorageId:  "P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value2",
			BundleHash: "test_hash2",
		})

		// ASSERT
		// check if bundle got finalized on pool
		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.CurrentKey).To(Equal("99"))
		Expect(pool.CurrentValue).To(Equal("test_value"))
		Expect(pool.CurrentHeight).To(Equal(uint64(100)))
		Expect(pool.TotalBundles).To(Equal(uint64(1)))

		// check if finalized bundle got saved
		finalizedBundle, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
		Expect(finalizedBundleFound).To(BeTrue())

		Expect(finalizedBundle.PoolId).To(Equal(uint64(0)))
		Expect(finalizedBundle.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))
		Expect(finalizedBundle.Uploader).To(Equal(i.ALICE))
		Expect(finalizedBundle.FromHeight).To(Equal(uint64(0)))
		Expect(finalizedBundle.ToHeight).To(Equal(uint64(100)))
		Expect(finalizedBundle.Key).To(Equal("99"))
		Expect(finalizedBundle.Value).To(Equal("test_value"))
		Expect(finalizedBundle.BundleHash).To(Equal("test_hash"))
		Expect(finalizedBundle.FinalizedAt).NotTo(BeZero())

		// check if next bundle proposal got registered
		bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposalFound).To(BeTrue())

		Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
		Expect(bundleProposal.StorageId).To(Equal("P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg"))
		Expect(bundleProposal.Uploader).To(Equal(i.BOB))
		Expect(bundleProposal.NextUploader).To(Equal(i.BOB))
		Expect(bundleProposal.ByteSize).To(Equal(uint64(100)))
		Expect(bundleProposal.ToHeight).To(Equal(uint64(200)))
		Expect(bundleProposal.ToKey).To(Equal("199"))
		Expect(bundleProposal.ToValue).To(Equal("test_value2"))
		Expect(bundleProposal.BundleHash).To(Equal("test_hash2"))
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(ContainElement(i.BOB))
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceBob))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		// calculate uploader reward
		totalReward := 100*s.App().BundlesKeeper.StorageCost(s.Ctx()) + pool.OperatingCost
		networkFee, _ := sdk.NewDecFromStr(s.App().BundlesKeeper.NetworkFee(s.Ctx()))

		treasuryReward := uint64(sdk.NewDec(int64(totalReward)).Mul(networkFee).RoundInt64())
		uploaderReward := totalReward - treasuryReward

		Expect(balanceStaker).To(Equal(initialBalanceAlice + uploaderReward))

		// check voter status
		valaccountVoter, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.BOB)
		Expect(valaccountVoter.Points).To(BeZero())

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		funder, _ := pool.GetFunder(i.ALICE)

		Expect(pool.Funders).To(HaveLen(1))
		Expect(funder.Amount).To(Equal(100*i.KYVE - totalReward))
	})

	PIt("Produce invalid bundles with two nodes", func() {
		// ARRANGE
		// stake a bit more than first node so >50% is reached
		s.RunTxStakersSuccess(&stakertypes.MsgCreateStaker{
			Creator: i.BOB,
			Amount:  200 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		initialBalanceBob = s.GetBalanceFromAddress(i.BOB)

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

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_NO,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value2",
			BundleHash: "test_hash2",
		})

		// ASSERT
		// check if bundle got not finalized on pool
		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.CurrentKey).To(Equal(""))
		Expect(pool.CurrentValue).To(BeEmpty())
		Expect(pool.CurrentHeight).To(BeZero())
		Expect(pool.TotalBundles).To(BeZero())

		// check if finalized bundle exists
		_, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
		Expect(finalizedBundleFound).To(BeFalse())

		// check if bundle proposal got dropped
		bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposalFound).To(BeTrue())

		Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
		Expect(bundleProposal.NextUploader).To(Equal(i.ALICE))
		Expect(bundleProposal.ByteSize).To(BeZero())
		Expect(bundleProposal.ToHeight).To(BeZero())
		Expect(bundleProposal.ToKey).To(BeEmpty())
		Expect(bundleProposal.ToValue).To(BeEmpty())
		Expect(bundleProposal.BundleHash).To(BeEmpty())
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(BeEmpty())
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceBob))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		Expect(balanceStaker).To(Equal(initialBalanceAlice))

		_, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), valaccountUploader.Staker)
		Expect(stakerFound).To(BeTrue())

		// Upload Slash
		fraction, _ := sdk.NewDecFromStr(s.App().StakersKeeper.UploadSlash(s.Ctx()))
		expectedBalance := sdk.NewDec(int64(100 * i.KYVE)).Mul(sdk.NewDec(1).Sub(fraction)).TruncateInt().Uint64()
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.ALICE)).To(Equal(expectedBalance))
		Expect(s.App().DelegationKeeper.GetDelegationOfPool(s.Ctx(), 0)).To(Equal(200*i.KYVE + expectedBalance))

		// check voter status
		valaccountVoter, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.BOB)
		Expect(valaccountVoter.Points).To(BeZero())

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		funder, _ := pool.GetFunder(i.ALICE)

		Expect(pool.Funders).To(HaveLen(1))
		Expect(funder.Amount).To(Equal(100 * i.KYVE))
	})

	PIt("Produce dropped bundle because nodes do not vote", func() {
		// ARRANGE
		// stake a bit more than first node so >50% is reached
		s.RunTxStakersSuccess(&stakertypes.MsgCreateStaker{
			Creator: i.BOB,
			Amount:  200 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		initialBalanceBob = s.GetBalanceFromAddress(i.BOB)

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

		// ACT
		// do not vote so bundle gets dropped
		s.CommitAfterSeconds(60)
		s.CommitAfterSeconds(1)

		// ASSERT
		// check if bundle got not finalized on pool
		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.CurrentKey).To(Equal(""))
		Expect(pool.CurrentValue).To(BeEmpty())
		Expect(pool.CurrentHeight).To(BeZero())
		Expect(pool.TotalBundles).To(BeZero())

		// check if finalized bundle exists
		_, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
		Expect(finalizedBundleFound).To(BeFalse())

		// check if bundle proposal got dropped
		bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposalFound).To(BeTrue())

		Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
		Expect(bundleProposal.StorageId).To(BeEmpty())
		Expect(bundleProposal.Uploader).To(BeEmpty())
		Expect(bundleProposal.NextUploader).To(Equal(i.ALICE))
		Expect(bundleProposal.ByteSize).To(BeZero())
		Expect(bundleProposal.ToHeight).To(BeZero())
		Expect(bundleProposal.ToKey).To(BeEmpty())
		Expect(bundleProposal.ToValue).To(BeEmpty())
		Expect(bundleProposal.BundleHash).To(BeEmpty())
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(BeEmpty())
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceBob))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		Expect(balanceStaker).To(Equal(initialBalanceAlice))

		//staker, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), valaccountUploader.Staker)
		//Expect(stakerFound).To(BeTrue()) TODO
		//
		//Expect(staker.Amount).To(Equal(100 * i.KYVE))
		//Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(Equal(300 * i.KYVE))

		// check voter status
		valaccountVoter, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.BOB)
		Expect(valaccountVoter.Points).To(Equal(uint64(1)))

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		funder, _ := pool.GetFunder(i.ALICE)

		Expect(pool.Funders).To(HaveLen(1))
		Expect(funder.Amount).To(Equal(100 * i.KYVE))
	})

	PIt("Produce dropped bundle because pool has not enough funds", func() {
		// ARRANGE
		s.RunTxPoolSuccess(&pooltypes.MsgDefundPool{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		// fund amount which definetely not cover bundle reward
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Amount:  1,
		})

		initialBalanceAlice := s.GetBalanceFromAddress(i.ALICE)

		// stake a bit more than first node so >50% is reached
		s.RunTxStakersSuccess(&stakertypes.MsgCreateStaker{
			Creator: i.BOB,
			Amount:  200 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.BOB,
			PoolId:     0,
			Valaddress: i.ALICE,
		})

		initialBalanceBob = s.GetBalanceFromAddress(i.BOB)

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

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
			Creator:   i.ALICE,
			Staker:    i.BOB,
			PoolId:    0,
			StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
			Vote:      bundletypes.VOTE_TYPE_YES,
		})

		s.CommitAfterSeconds(60)

		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.BOB,
			Staker:     i.ALICE,
			PoolId:     0,
			StorageId:  "P9edn0bjEfMU_lecFDIPLvGO2v2ltpFNUMWp5kgPddg",
			ByteSize:   100,
			FromHeight: 100,
			ToHeight:   200,
			FromKey:    "99",
			ToKey:      "199",
			ToValue:    "test_value2",
			BundleHash: "test_hash2",
		})

		// ASSERT
		// check if bundle got not finalized on pool
		pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		Expect(pool.CurrentKey).To(Equal(""))
		Expect(pool.CurrentValue).To(BeEmpty())
		Expect(pool.CurrentHeight).To(BeZero())
		Expect(pool.TotalBundles).To(BeZero())

		// check if finalized bundle exists
		_, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
		Expect(finalizedBundleFound).To(BeFalse())

		// check if bundle proposal got dropped
		bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(bundleProposalFound).To(BeTrue())

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

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.ALICE)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceBob))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		Expect(balanceStaker).To(Equal(initialBalanceAlice))

		//staker, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), valaccountUploader.Staker)
		//Expect(stakerFound).To(BeTrue())

		//Expect(staker.Amount).To(Equal(100 * i.KYVE)) TODO
		//Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(Equal(300 * i.KYVE))

		// check voter status
		valaccountVoter, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.BOB)
		Expect(valaccountVoter.Points).To(BeZero())

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(pool.Funders).To(BeEmpty())
	})
})

// TODO: Test if network fee gets correctly transferred to treasury (not so easy)
