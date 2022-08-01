package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/KYVENetwork/chain/x/bundles/types"
)

// ClaimUploaderRole handles the logic of an SDK message that allows protocol nodes to claim the uploader role.
// Note that this function can only be called while the specified pool is in "genesis state".
// This function obeys "first come, first serve" mentality.
func (k msgServer) ClaimUploaderRole(
	goCtx context.Context, msg *types.MsgClaimUploaderRole,
) (*types.MsgClaimUploaderRoleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check pool reached min stake
	if poolErr := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId); poolErr != nil {
		return nil, poolErr
	}

	if err := k.stakerKeeper.AssertAuthorized(ctx, msg.Staker, msg.Creator, msg.PoolId); err != nil {
		return nil, err
	}

	bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

	// Error if the pool isn't in "genesis state".
	if bundleProposal.NextUploader != "" {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrUploaderAlreadyClaimed.Error())
	}

	bundleProposal.NextUploader = msg.Staker
	bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())

	k.SetBundleProposal(ctx, bundleProposal)

	return &types.MsgClaimUploaderRoleResponse{}, nil
}
