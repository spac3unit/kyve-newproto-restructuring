package bundles_test

import (
	"testing"

	keepertest "github.com/KYVENetwork/chain/testutil/keeper"
	"github.com/KYVENetwork/chain/testutil/nullify"
	"github.com/KYVENetwork/chain/x/bundles"
	"github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BundlesKeeper(t)
	bundles.InitGenesis(ctx, *k, genesisState)
	got := bundles.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
