package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

var _ = Describe("Undelegation", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
			Name: "Moontest",
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
			Amount:  50 * i.KYVE,
		})

		s.CommitAfterSeconds(7)

		_, poolFound := s.App().PoolKeeper.GetPool(s.Ctx(), 0)
		Expect(poolFound).To(BeTrue())

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.ALICE,
			Amount:  100 * i.KYVE,
		})

		s.RunTxStakersSuccess(&stakerstypes.MsgStake{
			Creator: i.BOB,
			Amount:  100 * i.KYVE,
		})

		_, aliceFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
		Expect(aliceFound).To(BeTrue())

		_, bobFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.BOB)
		Expect(bobFound).To(BeTrue())

		s.CommitAfterSeconds(7)
	})

	It("Undelegate more KYVE than allowed", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))

		// Act
		res, err := s.RunTxDelegator(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  11 * i.KYVE,
		})

		// Assert
		Expect(res).To(BeNil())
		Expect(err).ToNot(BeNil())
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntriesOfDelegator(s.Ctx(), i.DUMMY[0])).To(BeEmpty())
	})

	It("Undelegate start unbonding", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})
		s.CommitAfterSeconds(1)

		// Assert

		//Delegation amount stays the same (due to unbonding)
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		unbondingEntries := s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntriesOfDelegator(s.Ctx(), i.DUMMY[0])
		Expect(unbondingEntries).To(HaveLen(1))
		Expect(unbondingEntries[0].Staker).To(Equal(i.ALICE))
		Expect(unbondingEntries[0].Delegator).To(Equal(i.DUMMY[0]))
		Expect(unbondingEntries[0].Amount).To(Equal(5 * i.KYVE))
		Expect(unbondingEntries[0].CreationTime).To(Equal(uint64(s.Ctx().BlockTime().Unix() - 1)))
	})

	It("Undelegate await unbonding", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})
		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(995 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(5 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(5 * i.KYVE))

		unbondingEntries := s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntriesOfDelegator(s.Ctx(), i.DUMMY[0])
		Expect(unbondingEntries).To(BeEmpty())
	})

	It("Redelegation during undelegation unbonding", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})
		s.RunTxDelegatorSuccess(&types.MsgRedelegate{
			Creator:    i.DUMMY[0],
			FromStaker: i.ALICE,
			ToStaker:   i.BOB,
			Amount:     10 * i.KYVE,
		})
		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		// Unbonding should have had no effect
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(0 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(0 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.BOB, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		unbondingEntries := s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntriesOfDelegator(s.Ctx(), i.DUMMY[0])
		Expect(unbondingEntries).To(BeEmpty())
	})

	It("Undelegate Slashed Amount", func() {
		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		params := s.App().StakersKeeper.GetParams(s.Ctx())
		params.UploadSlash = "0.1"
		s.App().StakersKeeper.SetParams(s.Ctx(), params)
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(9 * i.KYVE))

		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(999 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(0 * i.KYVE))
	})

	It("Delegate twice and undelegate twice", func() {
		// Assert
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		s.CommitAfterSeconds(10)
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})
		s.CommitAfterSeconds(10)
		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(980 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(30 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(20 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})
		s.CommitAfterSeconds(10)

		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  8 * i.KYVE,
		})
		s.CommitAfterSeconds(10)

		// Assert
		unbondingEntries := s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntries(s.Ctx())
		Expect(unbondingEntries).To(HaveLen(2))
		Expect(unbondingEntries[0].Staker).To(Equal(i.ALICE))
		Expect(unbondingEntries[0].Delegator).To(Equal(i.DUMMY[0]))
		Expect(unbondingEntries[0].Amount).To(Equal(5 * i.KYVE))
		Expect(unbondingEntries[0].CreationTime).To(Equal(uint64(s.Ctx().BlockTime().Unix() - 20)))

		Expect(unbondingEntries[1].Staker).To(Equal(i.ALICE))
		Expect(unbondingEntries[1].Delegator).To(Equal(i.DUMMY[1]))
		Expect(unbondingEntries[1].Amount).To(Equal(8 * i.KYVE))
		Expect(unbondingEntries[1].CreationTime).To(Equal(uint64(s.Ctx().BlockTime().Unix() - 10)))
	})

	It("Delegate twice and undelegate twice and await unbonding", func() {
		// Assert
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		s.CommitAfterSeconds(10)
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})
		s.CommitAfterSeconds(10)
		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(980 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(30 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(20 * i.KYVE))

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  5 * i.KYVE,
		})

		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  8 * i.KYVE,
		})

		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		unbondingEntries := s.App().DelegationKeeper.GetAllUnbondingDelegationQueueEntries(s.Ctx())
		Expect(unbondingEntries).To(BeEmpty())

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(995 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(5 * i.KYVE))

		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(988 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(12 * i.KYVE))
	})

	//It("Undelegate all after rewards and slashing and check state", func() {
	//
	//	slashes0 := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
	//	_ = slashes0
	//
	//	// Assert
	//	s.RunTxDelegatorSuccess(&types.MsgDelegate{
	//		Creator: i.DUMMY[0],
	//		Staker:  i.ALICE,
	//		Amount:  10 * i.KYVE,
	//	})
	//	s.CommitAfterSeconds(10)
	//	Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
	//	Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
	//	Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))
	//
	//	s.RunTxDelegatorSuccess(&types.MsgDelegate{
	//		Creator: i.DUMMY[1],
	//		Staker:  i.ALICE,
	//		Amount:  20 * i.KYVE,
	//	})
	//	s.CommitAfterSeconds(10)
	//	Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(980 * i.KYVE))
	//	Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(30 * i.KYVE))
	//	Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(20 * i.KYVE))
	//
	//	// Payout rewards
	//	s.App().DelegationKeeper.PayoutRewards(s.Ctx(), i.ALICE, 10*i.KYVE, pooltypes.ModuleName)
	//
	//	// Collect
	//	s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
	//		Creator: i.DUMMY[1],
	//		Staker:  i.ALICE,
	//	})
	//
	//	slashes1 := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
	//	_ = slashes1
	//
	//	// Slash 10%
	//	params := s.App().StakersKeeper.GetParams(s.Ctx())
	//	params.UploadSlash = "0.1"
	//	s.App().StakersKeeper.SetParams(s.Ctx(), params)
	//	s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)
	//
	//	Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(9 * i.KYVE))
	//	Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(18 * i.KYVE))
	//
	//	slashes2 := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
	//	_ = slashes2
	//
	//	// Act
	//	s.RunTxDelegatorSuccess(&types.MsgUndelegate{
	//		Creator: i.DUMMY[0],
	//		Staker:  i.ALICE,
	//		Amount:  9 * i.KYVE,
	//	})
	//	s.CommitAfterSeconds(10)
	//
	//	s.RunTxDelegatorSuccess(&types.MsgUndelegate{
	//		Creator: i.DUMMY[1],
	//		Staker:  i.ALICE,
	//		Amount:  18 * i.KYVE,
	//	})
	//	s.CommitAfterSeconds(10)
	//
	//	s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
	//	s.CommitAfterSeconds(1)
	//
	//	b0 := s.GetBalanceFromAddress(i.DUMMY[0])
	//	_ = b0
	//	b1 := s.GetBalanceFromAddress(i.DUMMY[1])
	//	_ = b1
	//
	//	delegationEntries := s.App().DelegationKeeper.GetAllDelegationEntries(s.Ctx())
	//	_ = delegationEntries
	//	delegators := s.App().DelegationKeeper.GetAllDelegators(s.Ctx())
	//	_ = delegators
	//	slashes := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
	//	_ = slashes
	//
	//	//Expect(len(slashes)).To(Equal(len(delegationEntries)))
	//	//Expect(delegators).To(HaveLen(0))
	//})

	// TODO test undelegate with multiple slashes

	// TODO test rewards retrieval on undelegation

})
