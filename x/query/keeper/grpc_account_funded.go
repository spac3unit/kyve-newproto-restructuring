package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountFundedList(goCtx context.Context, req *types.QueryAccountFundedListRequest) (*types.QueryAccountFundedListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var funded []types.Funded

	// TODO implement way in pool module to query funders
	_ = ctx

	//store := ctx.KVStore(k.storeKey)
	//// Build prefix. Store is already indexed in an optimal way
	//prefixBuilder := types.KeyPrefixBuilder{Key: types.KeyPrefix(types.FunderKeyPrefix)}.AString(req.Address).Key
	//funderStore := prefix.NewStore(store, prefixBuilder)
	//
	//pageRes, err := query.FilteredPaginate(funderStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
	//
	//	if accumulate {
	//
	//		var funder types.Funder
	//		if err := k.cdc.Unmarshal(value, &funder); err != nil {
	//			return false, err
	//		}
	//
	//		pool, _ := k.GetPool(ctx, funder.PoolId)
	//
	//		funded = append(funded, types.Funded{
	//			Account: funder.Account,
	//			Amount:  funder.Amount,
	//			Pool:    &pool,
	//		})
	//	}
	//
	//	return true, nil
	//})
	//
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	return &types.QueryAccountFundedListResponse{
		Funded: funded,
		//Pagination: pageRes,
	}, nil
}
