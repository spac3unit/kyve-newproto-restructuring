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
		s = i.NewCleanChain()
	})

	AfterEach(func() {
		s.VerifyPoolModuleAssetsIntegrity()
		s.VerifyPoolTotalFunds()
	})

	It("Create Pool", func() {
		// Arrange
		_, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeFalse())

		// Act
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

		// Assert
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

	It("Pause Pool", func() {
		// Arrange
		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
		Expect(pool.Paused).To(BeFalse())

		// Act
		s.RunTxPoolSuccess(&pooltypes.GovMsgPausePool{
			Creator: i.GOV,
			Id:      0,
		})

		// Assert
		poolAfter, foundAfter := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(foundAfter).To(BeTrue())
		Expect(poolAfter.Paused).To(BeTrue())

		poolAfter.Paused = false
		Expect(pool).To(Equal(poolAfter))
	})

	It("Pause Pool when already paused", func() {
		// Arrange
		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
		Expect(pool.Paused).To(BeTrue())

		// Act
		s.RunTxPoolError(&pooltypes.GovMsgPausePool{
			Creator: i.GOV,
			Id:      0,
		})

		// Assert
		poolAfter, foundAfter := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(foundAfter).To(BeTrue())
		Expect(pool).To(Equal(poolAfter))
	})

	It("Unpause Pool", func() {
		// Arrange
		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
		Expect(pool.Paused).To(BeTrue())

		// Act
		s.RunTxPoolSuccess(&pooltypes.GovMsgUnpausePool{
			Creator: i.GOV,
			Id:      0,
		})

		// Assert
		poolAfter, foundAfter := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(foundAfter).To(BeTrue())
		Expect(poolAfter.Paused).To(BeFalse())

		poolAfter.Paused = true
		Expect(pool).To(Equal(poolAfter))
	})

	It("Unpause Pool when already unpaused", func() {
		// Arrange
		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
		Expect(pool.Paused).To(BeFalse())

		// Act
		s.RunTxPoolError(&pooltypes.GovMsgUnpausePool{
			Creator: i.GOV,
			Id:      0,
		})

		// Assert
		poolAfter, foundAfter := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(foundAfter).To(BeTrue())
		Expect(pool).To(Equal(poolAfter))
	})

	It("Create Upgrade Pool proposal", func() {
		// Arrange
		pool, found := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(found).To(BeTrue())
		Expect(pool.Paused).To(BeFalse())

		// Act
		s.RunTxPoolSuccess(&pooltypes.GovMsgPoolUpgrade{
			Creator:     i.GOV,
			Runtime:     "Runtime",
			Version:     "new version",
			ScheduledAt: uint64(s.Ctx().BlockTime().Unix() + 1000),
			Duration:    60,
			Binaries:    "{\"test\": \"link.com\"}",
		})

		// Assert
		poolAfter, foundAfter := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(foundAfter).To(BeTrue())

		Expect(poolAfter.UpgradePlan.Version).To(Equal("new version"))
		Expect(poolAfter.UpgradePlan.ScheduledAt).To(Equal(uint64(s.Ctx().BlockTime().Unix() + 1000)))
		Expect(poolAfter.UpgradePlan.Duration).To(Equal(uint64(60)))
		Expect(poolAfter.UpgradePlan.Binaries).To(Equal("{\"test\": \"link.com\"}"))

		// Fast-forward
		s.CommitAfterSeconds(2000)
		s.CommitAfterSeconds(1)

		//TODO check
		//Expect(poolAfter.UpgradePlan).To(Equal(pooltypes.UpgradePlan{}))
	})
})
