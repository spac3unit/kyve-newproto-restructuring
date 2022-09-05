package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/bundles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SkipUploaderRole handles the logic of an SDK message that allows protocol nodes to skip an upload.
func (k msgServer) SkipUploaderRole(
	goCtx context.Context, msg *types.MsgSkipUploaderRole,
) (*types.MsgSkipUploaderRoleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.AssertCanPropose(ctx, msg.PoolId, msg.Staker, msg.Creator, msg.FromHeight); err != nil {
		return nil, err
	}

	pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)
	bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

	// reset points of uploader as node has proven to be active
	k.stakerKeeper.ResetPoints(ctx, msg.PoolId, msg.Staker)

	// Get next uploader from stakers voted
	voters := make([]string, 0)
	nextUploader := ""

	// exclude the staker who skips the uploader
	for _, voter := range bundleProposal.VotersValid {
		if voter != msg.Staker {
			voters = append(voters, voter)
		}
	}

	for _, voter := range bundleProposal.VotersInvalid {
		if voter != msg.Staker {
			voters = append(voters, voter)
		}
	}

	if len(voters) > 0 {
		nextUploader = k.chooseNextUploaderFromSelectedStakers(ctx, msg.PoolId, voters)
	} else {
		nextUploader = k.chooseNextUploaderFromAllStakers(ctx, msg.PoolId)
	}

	bundleProposal.NextUploader = nextUploader
	bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())

	k.SetBundleProposal(ctx, bundleProposal)

	if errEmit := ctx.EventManager().EmitTypedEvent(&types.EventSkippedUploaderRole{
		PoolId:           msg.PoolId,
		Id:               pool.TotalBundles,
		PreviousUploader: msg.Staker,
		NewUploader:      nextUploader,
	}); errEmit != nil {
		return nil, errEmit
	}

	return &types.MsgSkipUploaderRoleResponse{}, nil
}
