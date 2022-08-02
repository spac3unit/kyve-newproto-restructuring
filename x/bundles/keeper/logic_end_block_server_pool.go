package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/bundles/types"
	stakersmoduletypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// HandleUploadTimeout is an end block hook that triggers an upload timeout for every pool (if applicable).
func (k Keeper) HandleUploadTimeout(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Iterate over all pool Ids.
	for _, poolId := range []uint64{1, 2} /*TODO fetch pool ids*/ {
		// Set pool status

		err := k.poolKeeper.AssertPoolCanRun(ctx, poolId)

		bundleProposal, _ := k.GetBundleProposal(ctx, poolId)

		// Remove next uploader if pool is not active
		if err != nil {
			bundleProposal.NextUploader = ""
		}

		// Skip if we haven't reached the upload interval.
		if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + 0 /* TODO upload interval */) {
			continue
		}

		// Check if bundle needs to be dropped
		if bundleProposal.StorageId != "" && !strings.HasPrefix(bundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
			// check if the quorum was actually reached
			valid, invalid, abstain, total := k.getVoteDistribution(ctx, poolId)
			quorum := k.getQuorumStatus(valid, invalid, abstain, total)

			if quorum == types.BUNDLE_STATUS_NO_QUORUM {
				// handle stakers who did not vote at all
				k.handleNonVoters(ctx, poolId)

				// Get next uploader
				voters := append(bundleProposal.VotersValid, bundleProposal.VotersInvalid...)
				nextUploader := ""

				if len(voters) > 0 {
					nextUploader = k.getNextUploaderFromAddresses(ctx, poolId, voters)
				} else {
					nextUploader = k.getNextUploader(ctx, poolId)
				}

				// If consensus wasn't reached, we drop the bundle and emit an event.
				ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalised{
					PoolId:       poolId,
					StorageId:    bundleProposal.StorageId,
					ByteSize:     bundleProposal.ByteSize,
					Uploader:     bundleProposal.Uploader,
					NextUploader: bundleProposal.NextUploader,
					Reward:       0,
					Valid:        valid,
					Invalid:      invalid,
					//FromHeight:   bundleProposal.FromHeight, // TODO whats up with fromHeight?
					ToHeight: bundleProposal.ToHeight,
					Status:   types.BUNDLE_STATUS_NO_QUORUM,
					ToKey:    bundleProposal.ToKey,
					ToValue:  bundleProposal.ToValue,
					Id:       0,
					Abstain:  abstain,
					Total:    total,
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
		if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + 0 /* TODO pool.UploadInterval */ + k.UploadTimeout(ctx)) {
			continue
		}

		// We now know that the pool is active and the upload timeout has been reached.
		// Now we slash and remove the current next_uploader and select a new one.

		k.stakerKeeper.Slash(ctx, poolId, bundleProposal.NextUploader, stakersmoduletypes.SLASH_TYPE_TIMEOUT)

		// Update bundle proposal
		bundleProposal.NextUploader = k.getNextUploader(ctx, poolId)
		bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())

		k.SetBundleProposal(ctx, bundleProposal)
	}
}
