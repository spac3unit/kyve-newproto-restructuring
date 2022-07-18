package pool_test

import (
	"github.com/KYVENetwork/chain/testutil/nullify"
	"testing"

	keepertest "github.com/KYVENetwork/chain/testutil/keeper"
	"github.com/KYVENetwork/chain/x/pool"
	"github.com/KYVENetwork/chain/x/pool/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoolKeeper(t)
	pool.InitGenesis(ctx, *k, genesisState)
	got := pool.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
