package integration

import (
	"github.com/KYVENetwork/chain/x/pool"
	"github.com/KYVENetwork/chain/x/stakers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func (suite *KeeperTestSuite) RunTxPool(msg sdk.Msg) (*sdk.Result, error) {
	cachedCtx, commit := suite.ctx.CacheContext()
	resp, err := pool.NewHandler(suite.app.PoolKeeper)(cachedCtx, msg)
	if err == nil {
		commit()
		return resp, nil
	}
	return nil, err
}

func (suite *KeeperTestSuite) RunTxStakers(msg sdk.Msg) (*sdk.Result, error) {
	cachedCtx, commit := suite.ctx.CacheContext()
	resp, err := stakers.NewHandler(suite.app.StakersKeeper)(cachedCtx, msg)
	if err == nil {
		commit()
		return resp, nil
	}
	return nil, err
}

func (suite *KeeperTestSuite) RunTxPoolSuccess(t *testing.T, msg sdk.Msg) {
	_, err := suite.RunTxPool(msg)
	require.NoError(t, err)
}

func (suite *KeeperTestSuite) RunTxStakersSuccess(t *testing.T, msg sdk.Msg) {
	_, err := suite.RunTxStakers(msg)
	require.NoError(t, err)
}
