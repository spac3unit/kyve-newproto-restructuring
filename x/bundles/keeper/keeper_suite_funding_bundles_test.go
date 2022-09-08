package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	. "github.com/onsi/ginkgo/v2"
)

/*

TEST CASES - funding bundles

* TODO: Produce a valid bundle with only one funder
* TODO: Produce a valid bundle with multiple funders and same funding amounts
* TODO: Produce a valid bundle with multiple funders and different funding amounts

*/

var _ = Describe("funding bundles", Ordered, func() {
	s := i.NewCleanChain()

	//initialBalanceStaker0 := s.GetBalanceFromAddress(i.STAKER_0)
	//initialBalanceValaddress0 := s.GetBalanceFromAddress(i.VALADDRESS_0)
	//
	//initialBalanceStaker1 := s.GetBalanceFromAddress(i.STAKER_1)
	//initialBalanceValaddress1 := s.GetBalanceFromAddress(i.VALADDRESS_1)

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

		//initialBalanceStaker0 = s.GetBalanceFromAddress(i.STAKER_0)
		//initialBalanceValaddress0 = s.GetBalanceFromAddress(i.VALADDRESS_0)
		//
		//initialBalanceStaker1 = s.GetBalanceFromAddress(i.STAKER_1)
		//initialBalanceValaddress1 = s.GetBalanceFromAddress(i.VALADDRESS_1)

		s.CommitAfterSeconds(60)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})
})
