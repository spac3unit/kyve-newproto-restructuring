package keeper_test

import (
	i "github.com/KYVENetwork/chain/testutil/integration"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKeeper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func initPoolWithStakersAliceAndBob(s *i.KeeperTestSuite) {

}
