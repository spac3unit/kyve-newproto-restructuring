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
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	if poolErr := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId); poolErr != nil {
		return nil, poolErr
	}

	bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

	// Error if the pool isn't in "genesis state".
	if bundleProposal.NextUploader != "" {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrUploaderAlreadyClaimed.Error())
	}

	// TODO check staker is in pool

	// TODO pool has at least two stakers (is this still necessary)

	// TODO check pool reached min stake

	bundleProposal.NextUploader = msg.Creator
	bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())

	k.SetBundleProposal(ctx, bundleProposal)

	return &types.MsgClaimUploaderRoleResponse{}, nil
}
