package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	i "github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/delegation/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
)

/*

TEST CASES - msg_server_undelegate_test.go

* Undelegate more $KYVE than allowed
* Start undelegation; Check unbonding queue state
* Start undelegation and await unbonding
* Redelegation during undelegation unbonding
* Undelegate Slashed Amount
* Delegate twice and undelegate twice
* Delegate twice and undelegate twice and await unbonding
* Undelegate all after rewards and slashing
* JoinA, Slash, JoinB, PayoutReward
* Slash twice
* Start unbonding, slash twice, payout, await undelegation

TODO joinA slash joinB slash -> remaining delegation

*/

var _ = Describe("Delegation - Undelegation", Ordered, func() {
	s := i.NewCleanChain()

	BeforeEach(func() {
		s = i.NewCleanChain()

		CreateFundedPool(&s)

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

	AfterEach(func() {
		CheckAndContinueChainForOneMonth(&s)
	})

	It("Undelegate more $KYVE than allowed", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		s.PerformValidityChecks()

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

	It("Start undelegation; Check unbonding queue state", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))
		s.PerformValidityChecks()

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

	It("Start undelegation and await unbonding", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(990 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(10 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(10 * i.KYVE))
		s.PerformValidityChecks()

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
		s.PerformValidityChecks()

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
		s.PerformValidityChecks()

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
		s.PerformValidityChecks()

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

	It("Undelegate all after rewards and slashing", func() {
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
		s.PerformValidityChecks()

		// Payout rewards
		PayoutRewards(&s, i.ALICE, 10*i.KYVE)

		// Collect
		s.RunTxDelegatorSuccess(&types.MsgWithdrawRewards{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
		})

		// Slash 10%
		params := s.App().StakersKeeper.GetParams(s.Ctx())
		params.UploadSlash = "0.1"
		s.App().StakersKeeper.SetParams(s.Ctx(), params)
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)

		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(9 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(18 * i.KYVE))

		slashes2 := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())
		_ = slashes2

		// Act
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  9 * i.KYVE,
		})
		s.CommitAfterSeconds(10)

		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  18 * i.KYVE,
		})
		s.CommitAfterSeconds(10)

		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(uint64(1002_333_333_333)))
		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(uint64(1004_666_666_666)))

		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(BeZero())

		delegationEntries := s.App().DelegationKeeper.GetAllDelegationEntries(s.Ctx())
		delegators := s.App().DelegationKeeper.GetAllDelegators(s.Ctx())
		slashes := s.App().DelegationKeeper.GetAllDelegationSlashEntries(s.Ctx())

		// Only slash entries should be left
		Expect(len(slashes)).To(Equal(len(delegationEntries)))
		Expect(delegators).To(HaveLen(0))
	})

	It("JoinA, Slash, JoinB, PayoutReward", func() {

		// Arrange
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		params := s.App().StakersKeeper.GetParams(s.Ctx())
		params.UploadSlash = "0.5"
		s.App().StakersKeeper.SetParams(s.Ctx(), params)
		s.PerformValidityChecks()

		// Slash 50%
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})

		// Dummy0: 5$KYVE Dummy1: 20$KYVE
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(25 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(5 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(20 * i.KYVE))

		// Act
		PayoutRewards(&s, i.ALICE, 10*i.KYVE)

		// Assert
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(2 * i.KYVE))
		Expect(s.App().DelegationKeeper.GetOutstandingRewards(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(8 * i.KYVE))
	})

	It("Slash twice", func() {
		// Assert
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})

		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})
		s.PerformValidityChecks()

		// Act
		params := s.App().StakersKeeper.GetParams(s.Ctx())
		params.UploadSlash = "0.5"
		s.App().StakersKeeper.SetParams(s.Ctx(), params)

		// Slash 50% twice
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)

		// Assert
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(Equal(uint64(2_500_000_000 + 5_000_000_000)))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(Equal(uint64(2_500_000_000)))
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(Equal(uint64(5_000_000_000)))
	})

	It("Start unbonding, slash twice, payout, await undelegation", func() {
		// Assert
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		s.RunTxDelegatorSuccess(&types.MsgDelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})

		s.CommitAfterSeconds(10)
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[0],
			Staker:  i.ALICE,
			Amount:  10 * i.KYVE,
		})
		s.RunTxDelegatorSuccess(&types.MsgUndelegate{
			Creator: i.DUMMY[1],
			Staker:  i.ALICE,
			Amount:  20 * i.KYVE,
		})
		s.PerformValidityChecks()

		// Act
		params := s.App().StakersKeeper.GetParams(s.Ctx())
		params.UploadSlash = "0.5"
		s.App().StakersKeeper.SetParams(s.Ctx(), params)
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)
		s.App().DelegationKeeper.SlashDelegators(s.Ctx(), i.ALICE, stakerstypes.SLASH_TYPE_UPLOAD)

		PayoutRewards(&s, i.ALICE, 10*i.KYVE)

		s.CommitAfterSeconds(s.App().DelegationKeeper.UnbondingDelegationTime(s.Ctx()) + 1)
		s.CommitAfterSeconds(1)

		// Assert
		Expect(s.App().DelegationKeeper.GetDelegationAmount(s.Ctx(), i.ALICE)).To(BeZero())
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[0])).To(BeZero())
		Expect(s.App().DelegationKeeper.GetDelegationAmountOfDelegator(s.Ctx(), i.ALICE, i.DUMMY[1])).To(BeZero())

		Expect(s.GetBalanceFromAddress(i.DUMMY[0])).To(Equal(uint64(1000e9 - 7_500_000_000 + 3_333_333_333)))
		Expect(s.GetBalanceFromAddress(i.DUMMY[1])).To(Equal(uint64(1000e9 - 15_000_000_000 + 6_666_666_666)))
	})

})
