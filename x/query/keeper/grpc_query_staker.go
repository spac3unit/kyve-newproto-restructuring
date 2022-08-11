package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Stakers(c context.Context, req *types.QueryStakersRequest) (*types.QueryStakersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	stakers, pageRes, err := k.stakerKeeper.GetPaginatedStakerQuery(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	data := make([]types.StakerResponse, 0)

	for _, staker := range stakers {
		valaccounts := k.stakerKeeper.GetValaccountsFromStaker(ctx, staker.Address)

		data = append(data, types.StakerResponse{
			Staker: staker,
			Valaccounts: valaccounts,
		})
	}

	return &types.QueryStakersResponse{Stakers: data, Pagination: pageRes}, nil
}

func (k Keeper) Staker(c context.Context, req *types.QueryStakerRequest) (*types.QueryStakerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	staker, found := k.stakerKeeper.GetStaker(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	valaccounts := k.stakerKeeper.GetValaccountsFromStaker(ctx, staker.Address)

	return &types.QueryStakerResponse{Staker: types.StakerResponse{
		Staker: staker,
		Valaccounts: valaccounts,
	}}, nil
}
