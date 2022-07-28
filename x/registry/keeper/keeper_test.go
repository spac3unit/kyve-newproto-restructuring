package keeper_test

import (
	"github.com/KYVENetwork/chain/x/registry/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"
	"testing"

	"time"
)

// TODO maybe include in new tests?
func createGenesis(t *testing.T) {
	s = new(KeeperTestSuite)
	s.SetupTest()

	currentTime := s.ctx.BlockTime().Unix()
	s.CommitAfter(time.Second * 60)
	require.Equal(t, s.ctx.BlockTime().Unix(), currentTime+60)

	s.CommitAfter(time.Second * 60)
	require.Equal(t, s.ctx.BlockTime().Unix(), currentTime+2*60)

	pool := types.Pool{
		Creator:        govtypes.ModuleName,
		Name:           "Moontest",
		Runtime:        "@kyve/evm",
		Logo:           "9FJDam56yBbmvn8rlamEucATH5UcYqSBw468rlCXn8E",
		Config:         "{\"rpc\":\"https://rpc.api.moonbeam.network\",\"github\":\"https://github.com/KYVENetwork/evm\"}",
		UploadInterval: 60,
		OperatingCost:  100,
		BundleProposal: &types.BundleProposal{},
		MaxBundleSize:  100,
		Protocol: &types.Protocol{
			Version:     "1.3.0",
			LastUpgrade: uint64(s.ctx.BlockTime().Unix()),
			Binaries:    "{\"macos\":\"https://github.com/kyve-org/evm/releases/download/v1.0.5/kyve-evm-macos.zip\"}",
		},
		UpgradePlan: &types.UpgradePlan{},
		StartKey:    "0",
		Status:      types.POOL_STATUS_NOT_ENOUGH_VALIDATORS,
		MinStake:    0,
	}

	s.app.RegistryKeeper.AppendPool(s.ctx, pool)
	s.Commit()

}
