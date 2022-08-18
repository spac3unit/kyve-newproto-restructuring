package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gov Pool", Ordered, func() {
	s := i.NewCleanChain()

	BeforeAll(func() {
		// init new clean chain
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyPoolTotalFunds()
	})

	It("Create Pool", func() {
		// ACT
		s.RunTxPoolSuccess(&pooltypes.GovMsgCreatePool{
			Creator:        "govAddress",
			Name:           "Moontest",
			Runtime:        "Runtime",
			Logo:           "logo",
			Config:         "{\"config\": \"test\"}",
			StartKey:       "0",
			UploadInterval: 1,
			OperatingCost:  2,
			MinStake:       3,
			MaxBundleSize:  4,
			Version:        "1",
			Binaries:       "{\"b1\": \"string\"}",
		})

		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(found).To(BeTrue())
		Expect(pool.Name).To(Equal("Moontest"))
		Expect(pool.Runtime).To(Equal("Runtime"))
		Expect(pool.Config).To(Equal("{\"config\": \"test\"}"))
		Expect(pool.StartKey).To(Equal("0"))
		Expect(pool.UploadInterval).To(Equal(uint64(1)))
		Expect(pool.OperatingCost).To(Equal(uint64(2)))
		Expect(pool.MinStake).To(Equal(uint64(3)))
		Expect(pool.MaxBundleSize).To(Equal(uint64(4)))

		// TODO version binaries?
	})

	It("Update Pool", func() {
		s.RunTxPoolSuccess(&pooltypes.GovMsgUpdatePool{
			Creator: "gov",
			Id:      0,
			Payload: "{\"name\": \"Bitcoin\"}",
		})

		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)

		Expect(found).To(BeTrue())
		Expect(pool.Name).To(Equal("Bitcoin"))
		Expect(pool.Runtime).To(Equal("Runtime"))
		Expect(pool.Config).To(Equal("{\"config\": \"test\"}"))
		Expect(pool.StartKey).To(Equal("0"))
		Expect(pool.UploadInterval).To(Equal(uint64(1)))
		Expect(pool.OperatingCost).To(Equal(uint64(2)))
		Expect(pool.MinStake).To(Equal(uint64(3)))
		Expect(pool.MaxBundleSize).To(Equal(uint64(4)))
	})

})
