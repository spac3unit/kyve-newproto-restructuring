package keeper_test

import (
	"fmt"
	"github.com/KYVENetwork/chain/testutil/integration"
	"github.com/KYVENetwork/chain/x/pool/types"
	"testing"
)

func TestHello(t *testing.T) {

	fmt.Println("Hello Test")

	suite := integration.NewCleanChain()

	suite.RunTx(&types.MsgCreatePool{
		Creator: "kyve1rfg7q868r8x04y4pnacdcs8hpec6sqrkm28ra0",
		Name:    "test",
	})

	pool, found := suite.App().PoolKeeper.GetPool(suite.Ctx(), 0)
	fmt.Printf("Found %t", found)
	fmt.Printf("Found %v", pool)
}
