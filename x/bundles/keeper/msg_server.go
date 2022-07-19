package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/bundles/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) SubmitBundleProposal(goCtx context.Context, msg *types.MsgSubmitBundleProposal) (*types.MsgSubmitBundleProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) VoteProposal(goCtx context.Context, msg *types.MsgVoteProposal) (*types.MsgVoteProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) ClaimUploaderRole(goCtx context.Context, msg *types.MsgClaimUploaderRole) (*types.MsgClaimUploaderRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
