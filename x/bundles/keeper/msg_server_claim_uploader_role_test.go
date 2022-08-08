package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	bundletypes "github.com/KYVENetwork/chain/x/bundles/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
)

var _ = Describe("Claim Uploader Role", Ordered, func() {
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

		// fund pool
		s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
			Creator: i.ALICE,
			Id: 0,
			Amount: 100*i.KYVE,
		})
	})

	AfterEach(func() {
		// s.VerifyPoolModuleAssetsIntegrity()
	})

	It("Try to claim uploader role without being a staker", func () {
		// ACT
		s.RunTxBundlesError(&bundletypes.MsgClaimUploaderRole{
			Creator: i.BOB,
			Staker: i.ALICE,
			PoolId: 0,
		})

		// ASSERT
		_, found := s.App().BundlesKeeper.GetBundleProposal(s.Ctx(), 0)
		Expect(found).To(BeFalse())
	})

})