package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/registry/types"
)

type msgServer struct {
	Keeper
}

func (k msgServer) StakePool(ctx context.Context, pool *types.MsgStakePool) (*types.MsgStakePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) ReactivateStaker(ctx context.Context, staker *types.MsgReactivateStaker) (*types.MsgReactivateStakerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) UnstakePool(ctx context.Context, pool *types.MsgUnstakePool) (*types.MsgUnstakePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) UpdateMetadata(ctx context.Context, metadata *types.MsgUpdateMetadata) (*types.MsgUpdateMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k msgServer) UpdateCommission(ctx context.Context, commission *types.MsgUpdateCommission) (*types.MsgUpdateCommissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
