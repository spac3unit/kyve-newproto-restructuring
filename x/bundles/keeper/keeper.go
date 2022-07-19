package keeper

import (
	"fmt"
	delegationkeeper "github.com/KYVENetwork/chain/x/delegation/keeper"
	poolkeeper "github.com/KYVENetwork/chain/x/pool/keeper"
	stakerskeeper "github.com/KYVENetwork/chain/x/stakers/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/KYVENetwork/chain/x/bundles/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper       bankkeeper.Keeper
		poolKeeper       poolkeeper.Keeper
		stakerKeeper     stakerskeeper.Keeper
		delegationKeeper delegationkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper bankkeeper.Keeper,
	poolKeeper poolkeeper.Keeper,
	stakerKeeper stakerskeeper.Keeper,
	delegationKeeper delegationkeeper.Keeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		bankKeeper:       bankKeeper,
		poolKeeper:       poolKeeper,
		stakerKeeper:     stakerKeeper,
		delegationKeeper: delegationKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
