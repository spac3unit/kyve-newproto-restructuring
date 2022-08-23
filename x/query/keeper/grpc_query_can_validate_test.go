package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	querytypes "github.com/KYVENetwork/chain/x/query/types"
	"github.com/KYVENetwork/chain/x/registry/types"
	stakertypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Can Validate Tests", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest",
			MinStake:       200 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     0,
		})

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name:           "Moontest2",
			MinStake:       200 * i.KYVE,
			UploadInterval: 60,
			MaxBundleSize:  100,
			Protocol:       &pooltypes.Protocol{},
			UpgradePlan:    &pooltypes.UpgradePlan{},
		})

		s.RunTxStakersSuccess(&stakertypes.MsgStake{
			Creator: i.STAKER_1,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakertypes.MsgJoinPool{
			Creator:    i.STAKER_1,
			PoolId:     1,
			Valaddress: i.VALADDRESS_1,
			Amount:     0,
		})
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Can validate should fail if pool does not exist", func() {
		// ACT
		canValidate, err := s.App().QueryKeeper.CanValidate(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanValidateRequest{
			PoolId:     2,
			Valaddress: i.VALADDRESS_0,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canValidate.Possible).To(BeFalse())
		Expect(canValidate.Reason).To(Equal(sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), 2).Error()))
	})

	It("Can validate should fail if valaddress does not exist", func() {
		// ACT
		canValidate, err := s.App().QueryKeeper.CanValidate(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanValidateRequest{
			PoolId:     0,
			Valaddress: i.VALADDRESS_2,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canValidate.Possible).To(BeFalse())
		Expect(canValidate.Reason).To(Equal("no valaccount found"))
	})

	It("Can validate should fail if valaddress belongs to another pool", func() {
		// ACT
		canValidate, err := s.App().QueryKeeper.CanValidate(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanValidateRequest{
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canValidate.Possible).To(BeFalse())
		Expect(canValidate.Reason).To(Equal("no valaccount found"))
	})

	It("Can validate should succeed", func() {
		// ACT
		canValidate, err := s.App().QueryKeeper.CanValidate(sdk.WrapSDKContext(s.Ctx()), &querytypes.QueryCanValidateRequest{
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
		})

		// ASSERT
		Expect(err).To(BeNil())

		Expect(canValidate.Possible).To(BeTrue())
		Expect(canValidate.Reason).To(Equal(i.STAKER_0))
	})
})
