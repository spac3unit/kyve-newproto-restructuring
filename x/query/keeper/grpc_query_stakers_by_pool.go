package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakersByPool(c context.Context, req *types.QueryStakersByPoolRequest) (*types.QueryStakersByPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	data := make([]types.StakerPoolResponse, 0)

	ctx := sdk.UnwrapSDKContext(c)

	_, found := k.poolKeeper.GetPool(ctx, req.PoolId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	valaccounts := k.stakerKeeper.GetAllValaccountsOfPool(ctx, req.PoolId)

	for _, valaccount := range valaccounts {
		staker, stakerFound := k.stakerKeeper.GetStaker(ctx, valaccount.Staker)

		if stakerFound {
			data = append(data, types.StakerPoolResponse{
				Staker: &staker,
				Valaccount: valaccount,
			})
		}
	}

	return &types.QueryStakersByPoolResponse{Stakers: data}, nil
}
