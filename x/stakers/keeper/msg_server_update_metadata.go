package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UpdateMetadata ...
func (k msgServer) UpdateMetadata(goCtx context.Context, msg *types.MsgUpdateMetadata) (*types.MsgUpdateMetadataResponse, error) {
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
