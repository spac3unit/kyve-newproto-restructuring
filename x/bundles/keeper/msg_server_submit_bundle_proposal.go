package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/x/bundles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SubmitBundleProposal handles the logic of an SDK message that allows protocol nodes to submit a new bundle proposal.
func (k msgServer) SubmitBundleProposal(
	goCtx context.Context, msg *types.MsgSubmitBundleProposal,
) (*types.MsgSubmitBundleProposalResponse, error) {

	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO function should check minstake
	poolErr := k.poolKeeper.AssertPoolCanRun(ctx, msg.PoolId)
	if poolErr != nil {
		return nil, poolErr
	}

	if err := k.stakerKeeper.AssertAuthorized(ctx, msg.Staker, msg.Creator, msg.PoolId); err != nil {
		return nil, err
	}

	// TODO BEGIN BUNDLE LOGIC

	// Validate bundle id.
	if msg.StorageId == "" {
		return nil, types.ErrInvalidArgs
	}

	// Get current height from where the bundle proposal should resume
	// TODO where to handle current Height? pool module or bundle

	current_height := uint64(0)
	_ = current_height

	bundleProposal, _ := k.GetBundleProposal(ctx, msg.PoolId)

	// TODO outsource checks as they all just check for errors

	if bundleProposal.ToHeight != 0 {
		current_height = bundleProposal.ToHeight
	}

	// Validate from height
	if msg.FromHeight != current_height {
		return nil, types.ErrFromHeight
	}

	// Validate to height
	if msg.ToHeight < current_height {
		return nil, types.ErrToHeight
	}

	// TODO GET max bundle size from pool settings ?
	//if msg.ToHeight-current_height > pool.MaxBundleSize {
	//	return nil, types.ErrMaxBundleSize
	//}

	// TODO fetch current key from bundle module? I don't think this belongs to pool module
	//current_key := pool.CurrentKey
	current_key := ""

	// TODO also outsource checks

	if bundleProposal.ToKey != "" {
		current_key = bundleProposal.ToKey
	}

	// Validate from key
	if msg.FromKey != current_key {
		return nil, types.ErrFromKey
	}

	// Check if the sender is the designated uploader.
	if bundleProposal.NextUploader != msg.Creator {
		return nil, types.ErrNotDesignatedUploader
	}

	// Check if upload_interval has been surpassed
	//if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval) {
	if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + 0 /* TODO fetch upload interval */) {
		return nil, types.ErrUploadInterval
	}

	// TODO the entire process of evaluating the round could probably also be outsourced to a different file
	// TODO and then be reused by the endblock logic as well.

	//
	//	// EVALUATE PREVIOUS ROUND
	//
	//	// Check args of bundle types
	//	if strings.HasPrefix(msg.StorageId, types.KYVE_NO_DATA_BUNDLE) {
	//		// Validate bundle args
	//		if msg.ToHeight != current_height || msg.ByteSize != 0 {
	//			return nil, types.ErrInvalidArgs
	//		}
	//
	//		// Validate key values
	//		if msg.ToKey != "" || msg.ToValue != "" {
	//			return nil, types.ErrInvalidArgs
	//		}
	//
	//		// Validate bundle hash
	//		if msg.BundleHash != "" {
	//			return nil, types.ErrInvalidArgs
	//		}
	//	} else {
	//		if msg.ToHeight <= current_height || msg.ByteSize == 0 {
	//			return nil, types.ErrInvalidArgs
	//		}
	//
	//		// Validate key values
	//		if msg.ToKey == "" || msg.ToValue == "" {
	//			return nil, types.ErrInvalidArgs
	//		}
	//
	//		// Validate bundle hash
	//		if msg.BundleHash == "" {
	//			return nil, types.ErrInvalidArgs
	//		}
	//	}
	//
	//	// If bundle was dropped or is of type KYVE_NO_DATA_BUNDLE just register new bundle.
	//	if pool.BundleProposal.StorageId == "" || strings.HasPrefix(pool.BundleProposal.StorageId, types.KYVE_NO_DATA_BUNDLE) {
	//		pool.BundleProposal = &types.BundleProposal{
	//			Uploader:     msg.Creator,
	//			NextUploader: k.getNextUploaderByRandom(ctx, &pool, pool.Stakers),
	//			StorageId:    msg.StorageId,
	//			ByteSize:     msg.ByteSize,
	//			ToHeight:     msg.ToHeight,
	//			CreatedAt:    uint64(ctx.BlockTime().Unix()),
	//			ToKey:        msg.ToKey,
	//			ToValue:      msg.ToValue,
	//			BundleHash:   msg.BundleHash,
	//		}
	//
	//		// TODO replace with bundle (set bundle) maybe not event set it here
	//		k.SetPool(ctx, pool)
	//
	//		return &types.MsgSubmitBundleProposalResponse{}, nil
	//	}
	//
	//	// handle stakers who did not vote at all
	//	k.handleNonVoters(ctx, &pool)
	//
	//	// Get next uploader
	//	voters := append(pool.BundleProposal.VotersValid, pool.BundleProposal.VotersInvalid...)
	//	nextUploader := ""
	//
	//	if len(voters) > 0 {
	//		nextUploader = k.getNextUploaderByRandom(ctx, &pool, voters)
	//	} else {
	//		nextUploader = k.getNextUploaderByRandom(ctx, &pool, pool.Stakers)
	//	}
	//
	//	// check if the quorum was actually reached
	//	valid, invalid, abstain, total := k.getVoteDistribution(ctx, &pool)
	//	quorum := k.getQuorumStatus(valid, invalid, abstain, total)
	//
	//	// handle valid proposal
	//	if quorum == types.BUNDLE_STATUS_VALID {
	//		// Calculate the total reward for the bundle, and individual payouts.
	//		bundleReward := pool.OperatingCost + (pool.BundleProposal.ByteSize * k.StorageCost(ctx))
	//
	//		// load and parse network fee
	//		networkFee, err := sdk.NewDecFromStr(k.NetworkFee(ctx))
	//		if err != nil {
	//			k.PanicHalt(ctx, "Invalid value for params: "+err.Error())
	//		}
	//
	//		treasuryPayout := uint64(sdk.NewDec(int64(bundleReward)).Mul(networkFee).RoundInt64())
	//		uploaderPayout := bundleReward - treasuryPayout
	//
	//		// Calculate the delegation rewards for the uploader.
	//		uploader, foundUploader := k.GetStaker(ctx, pool.BundleProposal.Uploader, pool.Id)
	//		uploaderDelegation, foundUploaderDelegation := k.GetDelegationPoolData(ctx, pool.Id, pool.BundleProposal.Uploader)
	//
	//		if foundUploader && foundUploaderDelegation {
	//			// If the uploader has no delegators, it keeps the delegation reward.
	//
	//			if uploaderDelegation.DelegatorCount > 0 {
	//				// Calculate the reward, factoring in the node commission, and subtract from the uploader payout.
	//				commission, _ := sdk.NewDecFromStr(uploader.Commission)
	//				delegationReward := uint64(
	//					sdk.NewDec(int64(uploaderPayout)).Mul(sdk.NewDec(1).Sub(commission)).RoundInt64(),
	//				)
	//
	//				uploaderPayout -= delegationReward
	//				uploaderDelegation.CurrentRewards += delegationReward
	//
	//				k.SetDelegationPoolData(ctx, uploaderDelegation)
	//			}
	//		}
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
