package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Bundles module integration tests", Ordered, func() {
	s := i.NewCleanChain()

	//initialBalanceStaker0 := s.GetBalanceFromAddress(i.STAKER_0)
	//initialBalanceValaddress0 := s.GetBalanceFromAddress(i.VALADDRESS_0)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create clean pool for every test case
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MaxBundleSize:  100,
			StartKey:       "0",
			MinStake:       1 * i.KYVE,
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

		//initialBalanceStaker0 = s.GetBalanceFromAddress(i.STAKER_0)
		//initialBalanceValaddress0 = s.GetBalanceFromAddress(i.VALADDRESS_0)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	//It("If no next uploader is defined and timeout is reached the next uploader should be chosen", func() {
	//	// ACT
	//	s.CommitAfterSeconds(s.App().BundlesKeeper.UploadTimeout(s.Ctx()))
	//	s.CommitAfterSeconds(60)
	//	s.CommitAfterSeconds(1)
	//
	//	// ASSERT
	//	bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
	//	Expect(bundleProposalFound).To(BeTrue())
	//
	//	fmt.Println(bundleProposal)
	//
	//	Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
	//})
	//
	//It("If bundle proposal reached quorum and next uploader does not upload slash and remove him", func() {
	//	// ARRANGE
	//	s.RunTxBundlesSuccess(&bundletypes.MsgClaimUploaderRole{
	//		Creator: i.VALADDRESS_0,
	//		Staker:  i.STAKER_0,
	//		PoolId:  0,
	//	})
	//
	//	s.CommitAfterSeconds(60)
	//
	//	s.RunTxBundlesSuccess(&bundletypes.MsgSubmitBundleProposal{
	//		Creator:    i.VALADDRESS_0,
	//		Staker:     i.STAKER_0,
	//		PoolId:     0,
	//		StorageId:  "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
	//		ByteSize:   100,
	//		FromHeight: 0,
	//		ToHeight:   100,
	//		FromKey:    "0",
	//		ToKey:      "99",
	//		ToValue:    "test_value",
	//		BundleHash: "test_hash",
	//	})
	//
	//	s.RunTxStakersSuccess(&stakertypes.MsgStake{
	//		Creator: i.STAKER_1,
	//		Amount:  100 * i.KYVE,
	//	})
	//
	//	s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
	//		Creator:    i.STAKER_1,
	//		PoolId:     0,
	//		Valaddress: i.VALADDRESS_1,
	//		Amount:     0,
	//	})
	//
	//	s.RunTxBundlesSuccess(&bundletypes.MsgVoteBundleProposal{
	//		Creator:   i.VALADDRESS_1,
	//		Staker:    i.STAKER_1,
	//		PoolId:    0,
	//		StorageId: "y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI",
	//		Vote:      bundletypes.VOTE_TYPE_YES,
	//	})
	//
	//	bundleProposal, bundleProposalFound := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
	//	Expect(bundleProposalFound).To(BeTrue())
	//
	//	Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_0))
	//
	//	// ACT
	//	s.CommitAfterSeconds(s.App().BundlesKeeper.UploadTimeout(s.Ctx()))
	//	s.CommitAfterSeconds(60)
	//	s.CommitAfterSeconds(1)
	//
	//	// ASSERT
	//	stakersOfPool := s.App().StakersKeeper.GetAllStakerAddressesOfPool(s.Ctx(), 0)
	//	fmt.Println(stakersOfPool)
	//
	//	// check if bundle got not finalized on pool
	//	pool, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
	//	Expect(poolFound).To(BeTrue())
	//
	//	Expect(pool.CurrentKey).To(Equal(""))
	//	Expect(pool.CurrentValue).To(BeEmpty())
	//	Expect(pool.CurrentHeight).To(BeZero())
	//	Expect(pool.TotalBundles).To(BeZero())
	//
	//	// check if finalized bundle exists
	//	_, finalizedBundleFound := s.App().BundlesKeeper.GetFinalizedBundle(s.Ctx(), 0, 0)
	//	Expect(finalizedBundleFound).To(BeFalse())
	//
	//	// check if bundle proposal got only changed by next uploader
	//	bundleProposal, bundleProposalFound = s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
	//	Expect(bundleProposalFound).To(BeTrue())
	//
	//	Expect(bundleProposal.PoolId).To(Equal(uint64(0)))
	//	Expect(bundleProposal.StorageId).To(Equal("y62A3tfbSNcNYDGoL-eXwzyV-Zc9Q0OVtDvR1biJmNI"))
	//	Expect(bundleProposal.Uploader).To(Equal(i.STAKER_0))
	//	Expect(bundleProposal.NextUploader).To(Equal(i.STAKER_1))
	//	Expect(bundleProposal.ByteSize).To(Equal(uint64(100)))
	//	Expect(bundleProposal.ToHeight).To(Equal(uint64(100)))
	//	Expect(bundleProposal.ToKey).To(Equal("99"))
	//	Expect(bundleProposal.ToValue).To(Equal("test_value"))
	//	Expect(bundleProposal.BundleHash).To(Equal("test_hash"))
	//	Expect(bundleProposal.CreatedAt).NotTo(BeZero())
	//	Expect(bundleProposal.VotersValid).To(ContainElement(i.STAKER_0))
	//	Expect(bundleProposal.VotersInvalid).To(BeEmpty())
	//	Expect(bundleProposal.VotersAbstain).To(BeEmpty())
	//
	//	// check uploader status
	//	valaccountUploader, _ := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)
	//	Expect(valaccountUploader.Points).To(BeZero())
	//
	//	balanceValaddress := s.GetBalanceFromAddress(valaccountUploader.Valaddress)
	//	Expect(balanceValaddress).To(Equal(initialBalanceValaddress0))
	//
	//	balanceStaker := s.GetBalanceFromAddress(valaccountUploader.Staker)
	//
	//	Expect(balanceStaker).To(Equal(initialBalanceStaker0))
	//
	//	staker, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), valaccountUploader.Staker)
	//	Expect(stakerFound).To(BeTrue())
	//
	//	// Timeout Slash
	//	fraction, _ := sdk.NewDecFromStr(s.App().StakersKeeper.TimeoutSlash(s.Ctx()))
	//	expectedBalance := sdk.NewDec(int64(100 * i.KYVE)).Mul(sdk.NewDec(1).Sub(fraction)).TruncateInt().Uint64()
	//	Expect(expectedBalance).To(Equal(staker.Amount))
	//	Expect(s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)).To(Equal(expectedBalance))
	//
	//	// check pool funds
	//	pool, _ = s.App().PoolKeeper.GetPool(s.Ctx(), 0)
	//	funder, _ := pool.GetFunder(i.ALICE)
	//
	//	Expect(pool.Funders).To(HaveLen(1))
	//	Expect(funder.Amount).To(Equal(100 * i.KYVE))
	//})
})
