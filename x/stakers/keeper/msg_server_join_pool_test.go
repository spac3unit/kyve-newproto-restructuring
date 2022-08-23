package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Join Pool", Ordered, func() {
	s := i.NewCleanChain()

	initialBalanceStaker0 := uint64(0)
	initialBalanceValaddress0 := uint64(0)

	BeforeEach(func() {
		// init new clean chain
		s = i.NewCleanChain()

		// create pool
		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// create staker
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  100 * i.KYVE,
		})

		initialBalanceStaker0 = s.GetBalanceFromAddress(i.STAKER_0)
		initialBalanceValaddress0 = s.GetBalanceFromAddress(i.VALADDRESS_0)
	})

	AfterEach(func() {
		s.PerformValidityChecks()
	})

	It("Staker was just created", func() {
		// ASSERT
		valaccounts := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)
		Expect(valaccounts).To(HaveLen(0))
	})

	It("Join a pool", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		balanceAfterStaker0 := s.GetBalanceFromAddress(i.STAKER_0)
		balanceAfterValaddress0 := s.GetBalanceFromAddress(i.VALADDRESS_0)

		Expect(initialBalanceStaker0 - balanceAfterStaker0).To(Equal(100 * i.KYVE))
		Expect(balanceAfterValaddress0 - initialBalanceValaddress0).To(Equal(100 * i.KYVE))

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.STAKER_0))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal(i.VALADDRESS_0))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeFalse())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Stake after joining a pool", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)
		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.STAKER_0,
			Amount:  50 * i.KYVE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.STAKER_0))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal(i.VALADDRESS_0))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeFalse())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		totalStakeOfPool = s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(150 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Try to join the same pool with same valaddress again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join the same pool with a different valaddress again", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_1,
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join another pool with same valaddress", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest2",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     1,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)
		Expect(valaccountsOfStaker).To(HaveLen(1))
	})

	It("Try to join another pool with another valaddress", func() {
		// ARRANGE
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: i.VALADDRESS_0,
			Amount:     100 * i.KYVE,
		})

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest2",
			Protocol: &pooltypes.Protocol{
				Version:     "0.0.0",
				Binaries:    "{}",
				LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
			},
			UpgradePlan: &pooltypes.UpgradePlan{},
		})

		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     1,
			Valaddress: i.VALADDRESS_1,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)
		Expect(valaccountsOfStaker).To(HaveLen(2))
	})

	It("Join a pool with valaddress which does not exist on chain yet", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: "kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x",
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		balanceAfterStaker0 := s.GetBalanceFromAddress(i.STAKER_0)
		balanceAfterUnknown := s.GetBalanceFromAddress("kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x")

		Expect(initialBalanceStaker0 - balanceAfterStaker0).To(Equal(100 * i.KYVE))
		Expect(balanceAfterUnknown).To(Equal(100 * i.KYVE))

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.STAKER_0))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal("kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x"))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeFalse())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Join a pool with valaddress which does not exist on chain yet and send 0 funds", func() {
		// ACT
		s.RunTxStakersSuccess(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: "kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x",
			Amount:     0 * i.KYVE,
		})

		// ASSERT
		balanceAfterStaker0 := s.GetBalanceFromAddress(i.STAKER_0)
		balanceAfterUnknown := s.GetBalanceFromAddress("kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x")

		Expect(initialBalanceStaker0 - balanceAfterStaker0).To(BeZero())
		Expect(balanceAfterUnknown).To(BeZero())

		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(HaveLen(1))

		valaccount, found := s.App().StakersKeeper.GetValaccount(s.Ctx(), 0, i.STAKER_0)

		Expect(found).To(BeTrue())

		Expect(valaccount.Staker).To(Equal(i.STAKER_0))
		Expect(valaccount.PoolId).To(BeZero())
		Expect(valaccount.Valaddress).To(Equal("kyve1dx0nvx7y9d44jvr2dr6r2p636jea3f9827rn0x"))
		Expect(valaccount.Points).To(BeZero())
		Expect(valaccount.IsLeaving).To(BeFalse())

		valaccountsOfPool := s.App().StakersKeeper.GetAllValaccountsOfPool(s.Ctx(), 0)

		Expect(valaccountsOfPool).To(HaveLen(1))

		staker, _ := s.App().StakersKeeper.GetStaker(s.Ctx(), i.STAKER_0)
		totalStakeOfPool := s.App().StakersKeeper.GetTotalStake(s.Ctx(), 0)

		Expect(totalStakeOfPool).To(Equal(100 * i.KYVE))
		Expect(totalStakeOfPool).To(Equal(staker.Amount))
	})

	It("Join a pool with an invalid valaddress", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: "invalid_valaddress",
			Amount:     100 * i.KYVE,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.STAKER_0)

		Expect(valaccountsOfStaker).To(BeEmpty())
	})

	It("Join a pool with more amount than balance", func() {
		// ACT
		s.RunTxStakersError(&stakerstypes.MsgJoinPool{
			Creator:    i.STAKER_0,
			PoolId:     0,
			Valaddress: "invalid_valaddress",
			Amount:     initialBalanceStaker0 + 1,
		})

		// ASSERT
		valaccountsOfStaker := s.App().StakersKeeper.GetValaccountsFromStaker(s.Ctx(), i.ALICE)

		Expect(valaccountsOfStaker).To(BeEmpty())
	})
})
