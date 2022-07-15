package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/KYVENetwork/chain/testutil/keeper"
	"github.com/KYVENetwork/chain/x/delegation/keeper"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DelegationKeeper(t)

	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
