package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKeeper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func initPoolWithStakersAliceAndBob(s *i.KeeperTestSuite) {
	s.App().PoolKeeper.AppendPool(s.Ctx(), pooltypes.Pool{
		Name: "Moontest",
		Protocol: &pooltypes.Protocol{
			Version:     "0.0.0",
			Binaries:    "{}",
			LastUpgrade: uint64(s.Ctx().BlockTime().Unix()),
		},
		UpgradePlan: &pooltypes.UpgradePlan{},
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

	_, stakerFound := s.App().StakersKeeper.GetStaker(s.Ctx(), i.ALICE)
	Expect(stakerFound).To(BeTrue())

	_, stakerFound = s.App().StakersKeeper.GetStaker(s.Ctx(), i.BOB)
	Expect(stakerFound).To(BeTrue())

	s.RunTxPoolSuccess(&pooltypes.MsgFundPool{
		Creator: i.ALICE,
		Id:      0,
		Amount:  50 * i.KYVE,
	})

	s.CommitAfterSeconds(7)
}
