package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) StakePool(goCtx context.Context, msg *types.MsgStakePool) (*types.MsgStakePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) ReactivateStaker(goCtx context.Context, msg *types.MsgReactivateStaker) (*types.MsgReactivateStakerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) UnstakePool(goCtx context.Context, msg *types.MsgUnstakePool) (*types.MsgUnstakePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) UpdateMetadata(goCtx context.Context, msg *types.MsgUpdateMetadata) (*types.MsgUpdateMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) UpdateCommission(ctx context.Context, commission *types.MsgUpdateCommission) (*types.MsgUpdateCommissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
