package stakers

import (
	"github.com/KYVENetwork/chain/x/stakers/keeper"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, staker := range genState.StakerList {
		k.AppendStaker(ctx, staker)
	}

	for _, entry := range genState.ValaccountList {
		k.SetValaccount(ctx, entry)
		staker, _ := k.GetStaker(ctx, entry.Staker)
		k.AddToTotalStake(ctx, entry.PoolId, staker.Amount)
		k.AddOneToCount(ctx, entry.PoolId)
	}

	for _, entry := range genState.CommissionChangeEntries {
		k.SetCommissionChangeEntry(ctx, entry)
	}

	for _, entry := range genState.UnbondingStakeEntries {
		k.SetUnbondingStakeEntry(ctx, entry)
	}

	for _, entry := range genState.LeavePoolEntries {
		k.SetLeavePoolEntry(ctx, entry)
	}

	k.SetQueueState(ctx, types.QUEUE_IDENTIFIER_UNSTAKING, genState.QueueStateUnstaking)
	k.SetQueueState(ctx, types.QUEUE_IDENTIFIER_COMMISSION, genState.QueueStateCommission)
	k.SetQueueState(ctx, types.QUEUE_IDENTIFIER_LEAVE, genState.QueueStateLeave)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StakerList = k.GetAllStakers(ctx)

	genesis.ValaccountList = k.GetAllValaccounts(ctx)

	genesis.CommissionChangeEntries = k.GetAllCommissionChangeEntries(ctx)

	genesis.UnbondingStakeEntries = k.GetAllUnbondingStakeEntries(ctx)

	genesis.LeavePoolEntries = k.GetAllLeavePoolEntries(ctx)

	genesis.QueueStateUnstaking = k.GetQueueState(ctx, types.QUEUE_IDENTIFIER_UNSTAKING)

	genesis.QueueStateCommission = k.GetQueueState(ctx, types.QUEUE_IDENTIFIER_COMMISSION)

	genesis.QueueStateLeave = k.GetQueueState(ctx, types.QUEUE_IDENTIFIER_LEAVE)

	return genesis
}
