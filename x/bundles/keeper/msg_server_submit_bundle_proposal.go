package keeper

import (
	"context"
	"strings"

	"github.com/KYVENetwork/chain/x/bundles/types"
	stakersmoduletypes "github.com/KYVENetwork/chain/x/stakers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SubmitBundleProposal handles the logic of an SDK message that allows protocol nodes to submit a new bundle proposal.
func (k msgServer) SubmitBundleProposal(
	goCtx context.Context, msg *types.MsgSubmitBundleProposal,
) (*types.MsgSubmitBundleProposalResponse, error) {

	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO check min stake+delegation
	poolErr := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId)
	if poolErr != nil {
		return nil, poolErr
	}

	if err := k.stakerKeeper.AssertValaccountAuthorized(ctx, msg.PoolId, msg.Staker, msg.Creator); err != nil {
		return nil, err
	}

	// TODO BEGIN BUNDLE LOGIC
	bundleProposal, found := k.GetBundleProposal(ctx, msg.PoolId)

	if !found {
		return nil, sdkErrors.ErrNotFound
	}

	// Validate submit bundle args.
	err := k.validateSubmitBundleArgs(ctx, &bundleProposal, msg)
	if err != nil {
		return nil, err
	}

	// If bundle was dropped or is of type KYVE_NO_DATA_BUNDLE just register new bundle.
	if bundleProposal.StorageId == "" || strings.HasPrefix(bundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		nextUploader := k.chooseNextUploaderFromAllStakers(ctx, msg.PoolId)
		k.registerBundleProposalFromUploader(ctx, bundleProposal, msg, nextUploader)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	}
	
	// handle stakers who did not vote at all
	k.handleNonVoters(ctx, msg.PoolId)
	
	// Get next uploader
	voters := append(bundleProposal.VotersValid, bundleProposal.VotersInvalid...)
	nextUploader := ""

	if len(voters) > 0 {
		nextUploader = k.chooseNextUploaderFromSelectedStakers(ctx, msg.PoolId, voters)
	} else {
		nextUploader = k.chooseNextUploaderFromAllStakers(ctx, msg.PoolId)
	}

	// check if the quorum was actually reached
	valid, invalid, abstain, total := k.getVoteDistribution(ctx, msg.PoolId)
	quorum := k.getQuorumStatus(valid, invalid, abstain, total)
	
	// handle valid proposal
	if quorum == types.BUNDLE_STATUS_VALID {
		// Calculate the total reward for the bundle, and individual payouts.
		bundleReward, treasuryPayout, uploaderPayout, delegationPayout := k.calculatePayouts(ctx, msg.PoolId)

		if err := k.poolKeeper.ChargeFundersOfPool(ctx, msg.PoolId, bundleReward); err != nil {
			return nil, err
		}

		pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)
		bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

		if len(pool.Funders) == 0 {
			// drop bundle because pool ran out of funds
			bundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())
			k.SetBundleProposal(ctx, bundleProposal)
			// TODO: emit event

			return &types.MsgSubmitBundleProposalResponse{}, nil
		}

		// TODO: payout treasury
		// TODO: payout uploader
		// TODO: payout delegators

		// slash stakers who voted incorrectly
		for _, voter := range bundleProposal.VotersInvalid {
			k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakersmoduletypes.SLASH_TYPE_VOTE)
		}

		// save finalized bundle
		k.SetFinalizedBundle(ctx, types.FinalizedBundle{
			StorageId:   bundleProposal.StorageId,
			PoolId:      pool.Id,
			Id:          pool.TotalBundles,
			Uploader:    bundleProposal.Uploader,
			FromHeight:  pool.CurrentHeight,
			ToHeight:    bundleProposal.ToHeight,
			FinalizedAt: uint64(ctx.BlockHeight()),
			Key:         bundleProposal.ToKey,
			Value:       bundleProposal.ToValue,
			BundleHash:  bundleProposal.BundleHash,
		})

		// Finalize the proposal, saving useful information.
		// eventFromHeight := pool.CurrentHeight
		pool.CurrentHeight = bundleProposal.ToHeight
		pool.TotalBytes = pool.TotalBytes + bundleProposal.ByteSize
		pool.TotalBundles = pool.TotalBundles + 1
		pool.TotalBundleRewards = pool.TotalBundleRewards + bundleReward
		pool.CurrentKey = bundleProposal.ToKey
		pool.CurrentValue = bundleProposal.ToValue

		// TODO: emit event

		k.registerBundleProposalFromUploader(ctx, bundleProposal, msg, nextUploader)

		k.poolKeeper.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else if quorum == types.BUNDLE_STATUS_INVALID {
		// slash stakers who voted incorrectly - uploader received upload slash
		for _, voter := range bundleProposal.VotersValid {
			if voter == bundleProposal.Uploader {
				k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakersmoduletypes.SLASH_TYPE_UPLOAD)
			} else {
				k.stakerKeeper.Slash(ctx, msg.PoolId, voter, stakersmoduletypes.SLASH_TYPE_VOTE)
			}
		}

		bundleProposal = types.BundleProposal{
			NextUploader: bundleProposal.NextUploader,
			CreatedAt:    uint64(ctx.BlockTime().Unix()),
		}

		// TODO: emit event

		k.SetBundleProposal(ctx, bundleProposal)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else {
		return nil, types.ErrQuorumNotReached
	}
	//
	//		// Calculate the individual cost for each pool funder.
	//		// NOTE: Because of integer division, it is possible that there is a small remainder.
	//		// This remainder is in worst case MaxFundersAmount(tkyve) and is charged to the lowest funder.
	//		fundersCost := bundleReward / uint64(len(pool.Funders))
	//		fundersCostRemainder := bundleReward - (uint64(len(pool.Funders)) * fundersCost)
	//
	//		// TODO use pool payout funders method
	//
	//		// Fetch the lowest funder, and find a new one if the current one isn't found.
	//		lowestFunder, foundLowestFunder := k.GetFunder(ctx, pool.LowestFunder, pool.Id)
	//
	//		if !foundLowestFunder {
	//			k.updateLowestFunder(ctx, &pool)
	//			lowestFunder, _ = k.GetFunder(ctx, pool.LowestFunder, pool.Id)
	//		}
	//
	//		slashedFunds := uint64(0)
	//
	//		// Remove every funder who can't afford the funder cost.
	//		for fundersCost+fundersCostRemainder > lowestFunder.Amount {
	//			// Now, let's remove all other funders who have run out of funds.
	//			for _, account := range pool.Funders {
	//				funder, _ := k.GetFunder(ctx, account, pool.Id)
	//
	//				if funder.Amount < fundersCost {
	//					// remove funder
	//					k.removeFunder(ctx, &pool, &funder)
	//
	//					// transfer amount to treasury
	//					slashedFunds += funder.Amount
	//
	//					// Emit a defund event.
	//					errEmit := ctx.EventManager().EmitTypedEvent(&types.EventDefundPool{
	//						PoolId:  msg.Id,
	//						Address: funder.Account,
	//						Amount:  funder.Amount,
	//					})
	//					if errEmit != nil {
	//						return nil, errEmit
	//					}
	//				}
	//			}
	//
	//			if pool.TotalFunds > 0 {
	//				fundersCost = bundleReward / uint64(len(pool.Funders))
	//				fundersCostRemainder = bundleReward - (uint64(len(pool.Funders)) * fundersCost)
	//
	//				k.updateLowestFunder(ctx, &pool)
	//				lowestFunder, _ = k.GetFunder(ctx, pool.LowestFunder, pool.Id)
	//			} else {
	//				// Recalculate the lowest funder, update, and return.
	//				k.updateLowestFunder(ctx, &pool)
	//
	//				if slashedFunds > 0 {
	//					// transfer slashed funds to treasury
	//					err := k.transferToTreasury(ctx, slashedFunds)
	//					if err != nil {
	//						return nil, err
	//					}
	//				}
	//
	//				pool.BundleProposal = &types.BundleProposal{
	//					Uploader:      pool.BundleProposal.Uploader,
	//					NextUploader:  pool.BundleProposal.NextUploader,
	//					StorageId:     pool.BundleProposal.StorageId,
	//					ByteSize:      pool.BundleProposal.ByteSize,
	//					ToHeight:      pool.BundleProposal.ToHeight,
	//					CreatedAt:     uint64(ctx.BlockTime().Unix()),
	//					VotersValid:   pool.BundleProposal.VotersValid,
	//					VotersInvalid: pool.BundleProposal.VotersInvalid,
	//					ToKey:         pool.BundleProposal.ToKey,
	//					ToValue:       pool.BundleProposal.ToValue,
	//					BundleHash:    pool.BundleProposal.BundleHash,
	//				}
	//
	//				k.SetPool(ctx, pool)
	//
	//				// Emit a bundle dropped event because of insufficient funds.
	//				errEmit := ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalised{
	//					PoolId:       pool.Id,
	//					StorageId:    pool.BundleProposal.StorageId,
	//					ByteSize:     pool.BundleProposal.ByteSize,
	//					Uploader:     pool.BundleProposal.Uploader,
	//					NextUploader: pool.BundleProposal.NextUploader,
	//					Reward:       0,
	//					Valid:        valid,
	//					Invalid:      invalid,
	//					FromHeight:   pool.CurrentHeight,
	//					ToHeight:     pool.BundleProposal.ToHeight,
	//					Status:       types.BUNDLE_STATUS_NO_FUNDS,
	//					ToKey:        pool.BundleProposal.ToKey,
	//					ToValue:      pool.BundleProposal.ToValue,
	//					Id:           0,
	//					BundleHash:   pool.BundleProposal.BundleHash,
	//					Abstain:      abstain,
	//					Total:        total,
	//				})
	//				if errEmit != nil {
	//					return nil, errEmit
	//				}
	//
	//				return &types.MsgSubmitBundleProposalResponse{}, nil
	//			}
	//		}
	//
	//		if slashedFunds > 0 {
	//			// transfer slashed funds to treasury
	//			err := k.transferToTreasury(ctx, slashedFunds)
	//			if err != nil {
	//				return nil, err
	//			}
	//		}
	//
	//		// TODO still use pool.PayoutFunders etc.
	//		// Charge every funder equally.
	//		for _, account := range pool.Funders {
	//			funder, _ := k.GetFunder(ctx, account, pool.Id)
	//
	//			if funder.Amount >= fundersCost {
	//				funder.Amount -= fundersCost
	//			}
	//
	//			k.SetFunder(ctx, funder)
	//		}
	//
	//		// Remove any remainder cost from the lowest funder.
	//		lowestFunder, _ = k.GetFunder(ctx, pool.LowestFunder, pool.Id)
	//
	//		if lowestFunder.Amount >= fundersCostRemainder {
	//			lowestFunder.Amount -= fundersCostRemainder
	//		}
	//
	//		k.SetFunder(ctx, lowestFunder)
	//
	//		// Subtract bundle reward from the pool's total funds.
	//		pool.TotalFunds -= bundleReward
	//
	//		// Use stakers.Slash
	//		// Partially slash all nodes who voted incorrectly.
	//		for _, voter := range pool.BundleProposal.VotersInvalid {
	//			slashAmount := k.slashStaker(ctx, &pool, voter, k.VoteSlash(ctx))
	//
	//			errEmit := ctx.EventManager().EmitTypedEvent(&types.EventSlash{
	//				PoolId:    pool.Id,
	//				Address:   voter,
	//				Amount:    slashAmount,
	//				SlashType: types.SLASH_TYPE_VOTE,
	//			})
	//			if errEmit != nil {
	//				return nil, errEmit
	//			}
	//		}
	//
	//		// Send payout to treasury.
	//		errTreasury := k.transferToTreasury(ctx, treasuryPayout)
	//		if errTreasury != nil {
	//			return nil, errTreasury
	//		}
	//
	//		// Send payout to uploader.
	//		errTransfer := k.TransferToAddress(ctx, pool.BundleProposal.Uploader, uploaderPayout)
	//		if errTransfer != nil {
	//			return nil, errTransfer
	//		}
	//


	//		// save valid bundle
	//		k.SetProposal(ctx, types.Proposal{
	//			StorageId:   pool.BundleProposal.StorageId,
	//			PoolId:      pool.Id,
	//			Id:          pool.TotalBundles,
	//			Uploader:    pool.BundleProposal.Uploader,
	//			FromHeight:  pool.CurrentHeight,
	//			ToHeight:    pool.BundleProposal.ToHeight,
	//			FinalizedAt: uint64(ctx.BlockHeight()),
	//			Key:         pool.BundleProposal.ToKey,
	//			Value:       pool.BundleProposal.ToValue,
	//			BundleHash:  pool.BundleProposal.BundleHash,
	//		})
	//
	//		// Finalise the proposal, saving useful information.
	//		eventFromHeight := pool.CurrentHeight
	//		pool.CurrentHeight = pool.BundleProposal.ToHeight
	//		pool.TotalBytes = pool.TotalBytes + pool.BundleProposal.ByteSize
	//		pool.TotalBundles = pool.TotalBundles + 1
	//		pool.TotalBundleRewards = pool.TotalBundleRewards + bundleReward
	//		pool.CurrentKey = pool.BundleProposal.ToKey
	//		pool.CurrentValue = pool.BundleProposal.ToValue
	//
	//		// Emit a valid bundle event.
	//		errEmit := ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalised{
	//			PoolId:       pool.Id,
	//			StorageId:    pool.BundleProposal.StorageId,
	//			ByteSize:     pool.BundleProposal.ByteSize,
	//			Uploader:     pool.BundleProposal.Uploader,
	//			NextUploader: pool.BundleProposal.NextUploader,
	//			Reward:       bundleReward,
	//			Valid:        valid,
	//			Invalid:      invalid,
	//			FromHeight:   eventFromHeight,
	//			ToHeight:     pool.BundleProposal.ToHeight,
	//			Status:       types.BUNDLE_STATUS_VALID,
	//			ToKey:        pool.BundleProposal.ToKey,
	//			ToValue:      pool.BundleProposal.ToValue,
	//			Id:           pool.TotalBundles - 1,
	//			BundleHash:   pool.BundleProposal.BundleHash,
	//			Abstain:      abstain,
	//			Total:        total,
	//		})
	//		if errEmit != nil {
	//			return nil, errEmit
	//		}
	//
	//		// Set submitted bundle as new bundle proposal and select new next_uploader
	//		pool.BundleProposal = &types.BundleProposal{
	//			Uploader:     msg.Creator,
	//			NextUploader: nextUploader,
	//			StorageId:    msg.StorageId,
	//			ByteSize:     msg.ByteSize,
	//			ToHeight:     msg.ToHeight,
	//			CreatedAt:    uint64(ctx.BlockTime().Unix()),
	//			ToKey:        msg.ToKey,
	//			ToValue:      msg.ToValue,
	//			BundleHash:   msg.BundleHash,
	//		}
	//
	//		k.SetPool(ctx, pool)
	//
	//		return &types.MsgSubmitBundleProposalResponse{}, nil
	//	} else if quorum == types.BUNDLE_STATUS_INVALID {
	//		// Partially slash all nodes who voted incorrectly.
	//		for _, voter := range pool.BundleProposal.VotersValid {
	//			slashAmount := k.slashStaker(ctx, &pool, voter, k.VoteSlash(ctx))
	//
	//			errEmit := ctx.EventManager().EmitTypedEvent(&types.EventSlash{
	//				PoolId:    pool.Id,
	//				Address:   voter,
	//				Amount:    slashAmount,
	//				SlashType: types.SLASH_TYPE_VOTE,
	//			})
	//			if errEmit != nil {
	//				return nil, errEmit
	//			}
	//		}
	//
	//		// Partially slash the uploader.
	//		slashAmount := k.slashStaker(ctx, &pool, pool.BundleProposal.Uploader, k.UploadSlash(ctx))
	//
	//		// emit slash event
	//		errEmit := ctx.EventManager().EmitTypedEvent(&types.EventSlash{
	//			PoolId:    pool.Id,
	//			Address:   pool.BundleProposal.Uploader,
	//			Amount:    slashAmount,
	//			SlashType: types.SLASH_TYPE_UPLOAD,
	//		})
	//		if errEmit != nil {
	//			return nil, errEmit
	//		}
	//
	//		// Update the current lowest staker.
	//		k.updateLowestStaker(ctx, &pool)
	//
	//		// Emit an invalid bundle event.
	//		errEmit = ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalised{
	//			PoolId:       pool.Id,
	//			StorageId:    pool.BundleProposal.StorageId,
	//			ByteSize:     pool.BundleProposal.ByteSize,
	//			Uploader:     pool.BundleProposal.Uploader,
	//			NextUploader: pool.BundleProposal.NextUploader,
	//			Reward:       0,
	//			Valid:        valid,
	//			Invalid:      invalid,
	//			FromHeight:   pool.CurrentHeight,
	//			ToHeight:     pool.BundleProposal.ToHeight,
	//			Status:       types.BUNDLE_STATUS_INVALID,
	//			ToKey:        pool.BundleProposal.ToKey,
	//			ToValue:      pool.BundleProposal.ToValue,
	//			Id:           0,
	//			BundleHash:   pool.BundleProposal.BundleHash,
	//			Abstain:      abstain,
	//			Total:        total,
	//		})
	//		if errEmit != nil {
	//			return nil, errEmit
	//		}
	//
	//		// Update and return.
	//		pool.BundleProposal = &types.BundleProposal{
	//			NextUploader: pool.BundleProposal.NextUploader,
	//			CreatedAt:    uint64(ctx.BlockTime().Unix()),
	//		}
	//
	//		k.SetPool(ctx, pool)
	//
	//		return &types.MsgSubmitBundleProposalResponse{}, nil
	//	} else {
	//		return nil, types.ErrQuorumNotReached
	//	}

	return nil, nil

}
