package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// UpdateCommission ...
func (k msgServer) UpdateCommission(
	goCtx context.Context, msg *types.MsgUpdateCommission,
) (*types.MsgUpdateCommissionResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is a protocol node (aka has staked into this pool).
	_, isStaker := k.GetStaker(ctx, msg.Creator)
	if !isStaker {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrNoStaker.Error())
	}

	// Validate commission.
	commission, err := sdk.NewDecFromStr(msg.Commission)
	if err != nil {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidCommission.Error(), msg.Commission)
	}

	if commission.LT(sdk.NewDec(int64(0))) || commission.GT(sdk.NewDec(int64(1))) {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrInvalidCommission.Error(), msg.Commission)
	}

	k.orderNewCommissionChange(ctx, msg.Creator, msg.Commission)

	return &types.MsgUpdateCommissionResponse{}, nil
}

func (k msgServer) UpdateMetadata(
	goCtx context.Context, msg *types.MsgUpdateMetadata,
) (*types.MsgUpdateMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is a protocol node (aka has staked into this pool).
	_, isStaker := k.GetStaker(ctx, msg.Creator)
	if !isStaker {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrNoStaker.Error())
	}

	k.UpdateStakerMetadata(ctx, msg.Creator, msg.Moniker, msg.Website, msg.Logo)

	// Event an event.
	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventUpdateMetadata{
		Address: msg.Creator,
		Moniker: msg.Moniker,
		Website: msg.Website,
		Logo:    msg.Logo,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgUpdateMetadataResponse{}, nil
}
