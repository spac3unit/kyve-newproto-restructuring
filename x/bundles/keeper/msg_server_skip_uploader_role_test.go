package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	. "github.com/onsi/ginkgo/v2"
)

/*

TEST CASES - msg_server_skip_uploader_role.go

// TODO: implement
* Call can propose if pool does not exist
* Call can propose if pool is currently upgrading
* Call can propose if pool is paused
* Call can propose if pool is out of funds
* Call can propose if pool has not reached the minimum stake
* Call can propose with a valaccount which does not exist
* Call can propose as a staker who is not the next uploader
* Call can propose before the upload interval passed
* Call can propose with an invalid from height

// TODO: implement
* Skip uploader role if staker is not next uploader
* Skip uploader role if from height does not match with current one
* Skip uploader role if staker is next uploader
* Skip uploader role if staker is the only staker in pool
* Skip uploader role after uploader role has already been skipped

*/

var _ = Describe("msg_server_skip_uploader_role.go", Ordered, func() {
	s := i.NewCleanChain()

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
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})
})
