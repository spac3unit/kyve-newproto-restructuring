package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"
	delegationtypes "github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/KYVENetwork/chain/x/query/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AccountAssets returns an overview of the balances of the given user regarding the protocol nodes
// This includes the current balance, funding, staking, and delegation.
func (k Keeper) AccountAssets(goCtx context.Context, req *types.QueryAccountAssetsRequest) (*types.QueryAccountAssetsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	response := types.QueryAccountAssetsResponse{
		ProtocolStaking:          0, // TODO
		ProtocolStakingUnbonding: 0, // TODO
		ProtocolFunding:          0, // TODO
	}

	// =======
	// Balance
	// =======
	account, _ := sdk.AccAddressFromBech32(req.Address)
	balance := k.bankKeeper.GetBalance(ctx, account, "tkyve")
	response.Balance = balance.Amount.Uint64()

	// ==========================================
	// ProtocolStaking + ProtocolStakingUnbonding
	// ==========================================

	// TODO implement

	//// Iterate all Staker entries
	//// Fetches the total delegation and calculates the outstanding rewards
	//stakerStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StakerKeyPrefix))
	//var stakerPrefix []byte
	//stakerPrefix = append(stakerPrefix, []byte(req.Address)...)
	//stakerPrefix = append(stakerPrefix, []byte("/")...)
	//stakerIterator := sdk.KVStorePrefixIterator(stakerStore, stakerPrefix)
	//
	//defer stakerIterator.Close()
	//
	//for ; stakerIterator.Valid(); stakerIterator.Next() {
	//	var val types.Staker
	//	k.cdc.MustUnmarshal(stakerIterator.Value(), &val)
	//
	//	response.ProtocolStaking += val.Amount
	//}
	//
	//// Unbondings
	//// Iterate all UnbondingStaker entries to get total unbonding amount
	//unbondingStaker := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondingStakerKeyPrefix)
	//unbondingStakerIterator := sdk.KVStorePrefixIterator(unbondingStaker, types.KeyPrefixBuilder{}.AString(req.Address).Key)
	//
	//defer unbondingStakerIterator.Close()
	//
	//for ; unbondingStakerIterator.Valid(); unbondingStakerIterator.Next() {
	//	var val types.UnbondingStaker
	//	k.cdc.MustUnmarshal(unbondingStakerIterator.Value(), &val)
	//
	//	response.ProtocolStakingUnbonding += val.UnbondingAmount
	//}

	// ================================================
	// ProtocolDelegation + ProtocolDelegationUnbonding
	// ================================================

	// Iterate all Delegator entries
	delegatorStore := prefix.NewStore(ctx.KVStore(k.delegationKeeper.StoreKey()), util.GetByteKey(delegationtypes.DelegatorKeyPrefixIndex2, req.Address))
	delegatorIterator := sdk.KVStorePrefixIterator(delegatorStore, nil)
	defer delegatorIterator.Close()

	for ; delegatorIterator.Valid(); delegatorIterator.Next() {

		staker := string(delegatorIterator.Key()[0:43])

		response.ProtocolDelegation += k.delegationKeeper.GetDelegationAmountOfDelegator(ctx, staker, req.Address)
		response.ProtocolRewards += k.delegationKeeper.GetOutstandingRewards(ctx, staker, req.Address)
	}

	// ====================
	// Delegation Unbonding
	// ====================

	// Iterate all UnbondingDelegation entries to get total delegation unbonding amount
	for _, entry := range k.delegationKeeper.GetAllUnbondingDelegationQueueEntriesOfDelegator(ctx, req.Address) {
		response.ProtocolDelegationUnbonding += entry.Amount
	}

	// ===============
	// ProtocolFunding
	// ===============

	// TODO implement

	//// Iterate all funding entries
	//funderStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FunderKeyPrefix))
	//funderIterator := sdk.KVStorePrefixIterator(funderStore, []byte(req.Address))
	//
	//defer funderIterator.Close()
	//
	//for ; funderIterator.Valid(); funderIterator.Next() {
	//	var val types.Funder
	//	k.cdc.MustUnmarshal(funderIterator.Value(), &val)
	//
	//	response.ProtocolFunding += val.Amount
	//}

	return &types.QueryAccountAssetsResponse{}, nil
}
