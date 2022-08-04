package keeper

import (
	"context"
	"strings"

	"github.com/KYVENetwork/chain/x/bundles/types"
	stakersmoduletypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// HandleUploadTimeout is an end block hook that triggers an upload timeout for every pool (if applicable).
func (k Keeper) HandleUploadTimeout(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Iterate over all pool Ids.
	for _, poolId := range []uint64{1, 2} /*TODO fetch pool ids*/ {
		// Set pool status

		err := k.poolKeeper.AssertPoolCanRun(ctx, poolId)

		pool, _ := k.poolKeeper.GetPool(ctx, poolId)
		bundleProposal, _ := k.GetBundleProposal(ctx, poolId)

		// Remove next uploader if pool is not active
		if err != nil {
			bundleProposal.NextUploader = ""
		}

		// Skip if we haven't reached the upload interval.
		if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval) {
			continue
		}

		// Check if bundle needs to be dropped
		if bundleProposal.StorageId != "" && !strings.HasPrefix(bundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
			// check if the quorum was actually reached
			voteDistribution := k.getVoteDistribution(ctx, poolId)

			if voteDistribution.Status == types.BUNDLE_STATUS_NO_QUORUM {
				// handle stakers who did not vote at all
				k.handleNonVoters(ctx, poolId)

				// Get next uploader
				voters := append(bundleProposal.VotersValid, bundleProposal.VotersInvalid...)
				nextUploader := ""

				if len(voters) > 0 {
					nextUploader = k.chooseNextUploaderFromSelectedStakers(ctx, poolId, voters)
				} else {
					nextUploader = k.chooseNextUploaderFromAllStakers(ctx, poolId)
				}

				// If consensus wasn't reached, we drop the bundle and emit an event.
				ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalized{
					PoolId:       pool.Id,
					Id: pool.TotalBundles,
					Valid: voteDistribution.Valid,
					Invalid: voteDistribution.Invalid,
					Abstain: voteDistribution.Abstain,
					Total: voteDistribution.Total,
					Status: voteDistribution.Status,
				})

				bundleProposal = types.BundleProposal{
					NextUploader: nextUploader,
					CreatedAt:    uint64(ctx.BlockTime().Unix()),
				}

				k.SetBundleProposal(ctx, bundleProposal)
				continue
			}
		}

		// Skip if we haven't reached the upload timeout.
		if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval + k.UploadTimeout(ctx)) {
			continue
		}

		// We now know that the pool is active and the upload timeout has been reached.
		// Now we slash and remove the current next_uploader and select a new one.

		k.stakerKeeper.Slash(ctx, poolId, bundleProposal.NextUploader, stakersmoduletypes.SLASH_TYPE_TIMEOUT)

		k.stakerKeeper.RemoveValaccountFromPool(ctx, poolId, bundleProposal.NextUploader)


		// Update bundle proposal
		bundleProposal.NextUploader = k.chooseNextUploaderFromAllStakers(ctx, poolId)
		bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())

		k.SetBundleProposal(ctx, bundleProposal)
	}
}
