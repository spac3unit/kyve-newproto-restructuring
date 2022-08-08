package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
)

var _ = Describe("Bundle data flow integration tests", Ordered, func() {
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
		})
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyStakersModuleAssetsIntegrity()
	})

	// TODO: add test cases where pool is not created
})
