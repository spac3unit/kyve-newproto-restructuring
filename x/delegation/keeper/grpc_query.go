package keeper

import (
	"github.com/KYVENetwork/chain/x/delegation/types"
)

var _ types.QueryServer = Keeper{}
