package query

import (
	"github.com/KYVENetwork/chain/x/query/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		return nil, nil
	}
}
