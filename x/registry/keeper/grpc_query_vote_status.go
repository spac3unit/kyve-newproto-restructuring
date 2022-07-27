package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VoteStatus(c context.Context, req *types.QueryVoteStatusRequest) (*types.QueryVoteStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	//ctx := sdk.UnwrapSDKContext(c)
	//pool, found := k.GetPool(ctx, req.PoolId)
	//if !found {
	//	return nil, sdkerrors.ErrKeyNotFound
	//}
	//
	//valid, invalid, abstain, total := k.getVoteDistribution(ctx, &pool)
	//
	//return &types.QueryVoteStatusResponse{
	//	VoteStatus: &types.VoteStatusResponse{
	//		Valid:   valid,
	//		Invalid: invalid,
	//		Abstain: abstain,
	//		Total:   total,
	//	},
	//}, nil

	return &types.QueryVoteStatusResponse{}, nil
}
