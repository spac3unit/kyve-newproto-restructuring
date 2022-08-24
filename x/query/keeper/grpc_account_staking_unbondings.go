package keeper

import (
	"context"
	"encoding/binary"
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/query/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountStakingUnbondings(goCtx context.Context, req *types.QueryAccountStakingUnbondingsRequest) (*types.QueryAccountStakingUnbondingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var stakingUnbondings []types.StakingUnbonding
	// TODO maybe move this function to the stakersModule and provide a function as an argument
	store := prefix.NewStore(ctx.KVStore(k.stakerKeeper.StoreKey()), util.GetByteKey(stakerstypes.UnbondingStakingEntryKeyPrefixIndex2, req.Address))
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			index := binary.BigEndian.Uint64(key[0:8])
			unbondingEntry, _ := k.stakerKeeper.GetUnbondingStakeEntry(ctx, index)

			stakingUnbondings = append(stakingUnbondings, types.StakingUnbonding{
				Amount:       unbondingEntry.Amount,
				CreationTime: unbondingEntry.CreationDate,
			})
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAccountStakingUnbondingsResponse{
		Unbondings: stakingUnbondings,
		Staker:     k.GetFullStaker(ctx, req.Address),
		Pagination: pageRes,
	}, nil
}
