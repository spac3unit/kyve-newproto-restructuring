package delegation

import (
	"github.com/KYVENetwork/chain/x/delegation/keeper"
	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	k.SetParams(ctx, genState.Params)

	for _, delegator := range genState.DelegatorList {
		k.SetDelegator(ctx, delegator)
	}

	for _, entry := range genState.DelegationEntries {
		k.SetDelegationEntries(ctx, entry)
	}

	for _, entry := range genState.DelegationData {
		k.SetDelegationData(ctx, entry)
	}

	for _, entry := range genState.UndelegationQueueEntry {
		// TODO set undelegation queue entry
		_ = entry
	}

	// TODO set Undelegation Queue State

	// TODO set redelegation

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DelegatorList = k.GetAllDelegator(ctx)

	genesis.DelegationEntries = k.GetAllDelegationEntries(ctx)

	genesis.DelegationData = k.GetAllDelegationData(ctx)

	// TODO set undelegation queue entry

	// TODO set queue state

	// TODO set redelegation cooldown

	return genesis
}
