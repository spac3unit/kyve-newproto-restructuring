package keeper

import (
	"github.com/KYVENetwork/chain/x/query/types"
)

var _ types.QueryServer = Keeper{}
var _ types.QueryDelegationServer = Keeper{}
