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

/*

TEST CASES - valid bundles

* Produce a valid bundle with one validator and no foreign delegations
* TODO: Produce a valid bundle with one validator and foreign delegations
* TODO: Produce a valid bundle with multiple validators and no foreign delegations
* TODO: Produce a valid bundle with multiple validators and foreign delegations
* TODO: Produce a valid bundle with only one funder
* TODO: Produce a valid bundle with multiple funders and same funding amounts
* TODO: Produce a valid bundle with multiple funders and different funding amounts

*/

var _ = Describe("valid bundles", Ordered, func() {
	s := i.NewCleanChain()

	initialBalanceStaker := s.GetBalanceFromAddress(i.STAKER_0)
	initialBalanceValaddress := s.GetBalanceFromAddress(i.VALADDRESS_0)

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
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
		})

		s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
			Creator: i.VALADDRESS_0,
			Staker:  i.STAKER_0,
			PoolId:  0,
		})

		initialBalanceStaker = s.GetBalanceFromAddress(i.STAKER_0)
		initialBalanceValaddress = s.GetBalanceFromAddress(i.VALADDRESS_0)

		s.CommitAfterSeconds(60)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	PIt("Produce a valid bundle with one validator and no foreign delegations", func() {
		// ARRANGE
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

		s.CommitAfterSeconds(60)

		// ACT
		s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
			Creator:    i.VALADDRESS_0,
			Staker:     i.STAKER_0,
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
		Expect(finalizedBundle.Uploader).To(Equal(i.STAKER_0))
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
		Expect(bundleProposal.Uploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
		Expect(bundleProposal.ByteSize).To(Equal(uint64(100)))
		Expect(bundleProposal.ToHeight).To(Equal(uint64(200)))
		Expect(bundleProposal.ToKey).To(Equal("199"))
		Expect(bundleProposal.ToValue).To(Equal("test_value2"))
		Expect(bundleProposal.BundleHash).To(Equal("test_hash2"))
		Expect(bundleProposal.CreatedAt).NotTo(BeZero())
		Expect(bundleProposal.VotersValid).To(ContainElement(i.STAKER_0))
		Expect(bundleProposal.VotersInvalid).To(BeEmpty())
		Expect(bundleProposal.VotersAbstain).To(BeEmpty())

		// check uploader status
		valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)
		Expect(valaccountUploader.Points).To(BeZero())

		balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
		Expect(balanceValaddress).To(Equal(initialBalanceValaddress))

		balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)

		// calculate uploader reward
		totalReward := 100*s.App().BundlesKeeper.StorageCost(s.Ctx()) + pool.OperatingCost
		networkFee, _ := sdk.NewDecFromStr(s.App().BundlesKeeper.NetworkFee(s.Ctx()))

		treasuryReward := uint64(sdk.NewDec(int64(totalReward)).Mul(networkFee).RoundInt64())
		uploaderReward := totalReward - treasuryReward

		Expect(balanceStaker + s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.STAKER_0, i.STAKER_0)).To(Equal(initialBalanceStaker + uploaderReward))

		// check pool funds
		pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		funder, _ := pool.GetFunder(i.ALICE)

		Expect(pool.Funders).To(HaveLen(1))
		Expect(funder.Amount).To(Equal(100*i.KYVE - totalReward))
	})
})
