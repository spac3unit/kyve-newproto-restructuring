package keeper

import (
	"context"
	"encoding/binary"
	"github.com/KYVENetwork/chain/util"
	delegationtypes "github.com/KYVENetwork/chain/x/delegation/types"
	"github.com/KYVENetwork/chain/x/query/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountDelegationUnbondings(goCtx context.Context, req *types.QueryAccountDelegationUnbondingsRequest) (*types.QueryAccountDelegationUnbondingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var delegationUnbondings []types.DelegationUnbonding

	store := prefix.NewStore(ctx.KVStore(k.delegationKeeper.StoreKey()), util.GetByteKey(delegationtypes.UndelegationQueueKeyPrefixIndex2, req.Address))
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			index := binary.BigEndian.Uint64(key[0:8])
			unbondingEntry, _ := k.delegationKeeper.GetUndelegationQueueEntry(ctx, index)

			//staker, _ := k.GetStaker(ctx, unbondingEntry.Staker, unbondingEntry.PoolId)
			//// Load unbondingStaker
			//unbondingStaker, _ := k.GetUnbondingStaker(ctx, unbondingEntry.PoolId, unbondingEntry.Staker)
			//
			//// Fetch total delegation for staker, as it is stored in DelegationPoolData
			//poolDelegationData, _ := k.GetDelegationPoolData(ctx, staker.PoolId, staker.Account)

			//stakerResponse := types.StakerResponse{
			//	Staker:          staker.Account,
			//	PoolId:          staker.PoolId,
			//	Account:         staker.Account,
			//	Amount:          staker.Amount,
			//	TotalDelegation: poolDelegationData.TotalDelegation,
			//	Commission:      staker.Commission,
			//	Moniker:         staker.Moniker,
			//	Website:         staker.Website,
			//	Logo:            staker.Logo,
			//	Points:          staker.Points,
			//	UnbondingAmount: unbondingStaker.UnbondingAmount,
			//}

			delegationUnbondings = append(delegationUnbondings, types.DelegationUnbonding{
				Amount:       unbondingEntry.Amount,
				CreationTime: unbondingEntry.CreationTime,
				Staker:       nil,
			})
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAccountDelegationUnbondingsResponse{
		Unbondings: delegationUnbondings,
		Pagination: pageRes,
	}, nil
}
