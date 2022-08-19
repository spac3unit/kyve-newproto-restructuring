package keeper

import (
	"math"
	"math/rand"
	"sort"
	"strings"

	"github.com/KYVENetwork/chain/x/bundles/types"
	poolmoduletypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AssertPoolCanRun ...
func (k Keeper) AssertPoolCanRun(ctx sdk.Context, poolId uint64) error {

	pool, poolErr := k.poolKeeper.GetPoolWithError(ctx, poolId)
	if poolErr != nil {
		return poolErr
	}

	// Error if the pool is upgrading.
	if pool.UpgradePlan.ScheduledAt > 0 && uint64(ctx.BlockTime().Unix()) >= pool.UpgradePlan.ScheduledAt {
		return sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrPoolCurrentlyUpgrading.Error())
	}

	// Error if the pool is paused.
	if pool.Paused {
		return sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrPoolPaused.Error())
	}

	// Error if the pool has no funds.
	if len(pool.Funders) == 0 {
		return sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, types.ErrPoolOutOfFunds.Error())
	}

	// Error if min stake is not reached
	if k.stakerKeeper.GetTotalStake(ctx, pool.Id) < pool.MinStake {
		return sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, types.ErrMinStakeNotReached.Error())
	}

	return nil
}

func (k Keeper) validateSubmitBundleArgs(ctx sdk.Context, bundleProposal *types.BundleProposal, msg *types.MsgSubmitBundleProposal) error {
	pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)

	currentHeight := bundleProposal.ToHeight
	currentKey := bundleProposal.ToKey

	// Validate storage id
	if msg.StorageId == "" {
		return types.ErrInvalidArgs
	}

	// Check if the sender is the designated uploader.
	if bundleProposal.NextUploader != msg.Staker {
		return sdkErrors.Wrapf(types.ErrNotDesignatedUploader, "expected %v received %v", bundleProposal.NextUploader, msg.Staker)
	}

	// Validate upload interval has been surpassed
	if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval) {
		return sdkErrors.Wrapf(types.ErrUploadInterval, "expected %v < %v", ctx.BlockTime().Unix(), bundleProposal.CreatedAt+pool.UploadInterval)
	}

	// Validate if bundle is not too big
	if msg.ToHeight-currentHeight > pool.MaxBundleSize {
		return sdkErrors.Wrapf(types.ErrMaxBundleSize, "expected %v received %v", pool.MaxBundleSize, msg.ToHeight-currentHeight)
	}

	// Validate from height
	if msg.FromHeight != currentHeight {
		return sdkErrors.Wrapf(types.ErrFromHeight, "expected %v received %v", currentHeight, msg.FromHeight)
	}

	// Validate to height
	if msg.ToHeight < currentHeight {
		return sdkErrors.Wrapf(types.ErrToHeight, "expected %v < %v", msg.ToHeight, currentHeight)
	}

	// Validate from key
	if currentKey != "" && msg.FromKey != currentKey {
		return types.ErrFromKey
	}

	// check if bundle is of type no data bundle
	if strings.HasPrefix(msg.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		// Validate bundle args
		if msg.ToHeight != currentHeight || msg.ByteSize != 0 {
			return types.ErrInvalidArgs
		}

		// Validate key values
		if msg.ToKey != "" || msg.ToValue != "" {
			return types.ErrInvalidArgs
		}

		// Validate bundle hash
		if msg.BundleHash != "" {
			return types.ErrInvalidArgs
		}
	} else {
		if msg.ToHeight <= currentHeight || msg.ByteSize == 0 {
			return types.ErrInvalidArgs
		}

		// Validate key values
		if msg.ToKey == "" || msg.ToValue == "" {
			return types.ErrInvalidArgs
		}

		// Validate bundle hash
		if msg.BundleHash == "" {
			return types.ErrInvalidArgs
		}
	}

	return nil
}

func (k Keeper) registerBundleProposalFromUploader(ctx sdk.Context, pool poolmoduletypes.Pool, bundleProposal types.BundleProposal, msg *types.MsgSubmitBundleProposal, nextUploader string) error {
	validVoters := make([]string, 0)

	if !strings.HasPrefix(msg.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		validVoters = append(validVoters, msg.Staker)
	}

	bundleProposal = types.BundleProposal{
		PoolId:       msg.PoolId,
		Uploader:     msg.Staker,
		NextUploader: nextUploader,
		StorageId:    msg.StorageId,
		ByteSize:     msg.ByteSize,
		ToHeight:     msg.ToHeight,
		CreatedAt:    uint64(ctx.BlockTime().Unix()),
		VotersValid:  validVoters,
		ToKey:        msg.ToKey,
		ToValue:      msg.ToValue,
		BundleHash:   msg.BundleHash,
	}

	k.SetBundleProposal(ctx, bundleProposal)

	err := ctx.EventManager().EmitTypedEvent(&types.EventBundleProposed{
		PoolId:     bundleProposal.PoolId,
		Id:         pool.TotalBundles,
		StorageId:  bundleProposal.StorageId,
		Uploader:   bundleProposal.Uploader,
		ByteSize:   bundleProposal.ByteSize,
		FromHeight: pool.CurrentHeight,
		ToHeight:   bundleProposal.ToHeight,
		FromKey:    pool.CurrentKey,
		ToKey:      bundleProposal.ToKey,
		Value:      bundleProposal.ToValue,
		BundleHash: bundleProposal.BundleHash,
		CreatedAt:  bundleProposal.CreatedAt,
	})

	return err
}

// updateLowestFunder is an internal function that updates the lowest funder entry in a given pool.
func (k Keeper) handleNonVoters(ctx sdk.Context, poolId uint64) {
	voters := map[string]bool{}
	bundleProposal, _ := k.GetBundleProposal(ctx, poolId)

	for _, address := range bundleProposal.VotersValid {
		voters[address] = true
	}

	for _, address := range bundleProposal.VotersInvalid {
		voters[address] = true
	}

	for _, address := range bundleProposal.VotersAbstain {
		voters[address] = true
	}

	for _, staker := range k.stakerKeeper.GetAllStakerAddressesOfPool(ctx, poolId) {
		if !voters[staker] {
			k.stakerKeeper.AddPoint(ctx, poolId, staker)
		}
	}
}

func (k Keeper) calculatePayouts(ctx sdk.Context, poolId uint64) (bundleReward types.BundleReward) {
	pool, _ := k.poolKeeper.GetPool(ctx, poolId)
	bundleProposal, _ := k.GetBundleProposal(ctx, poolId)

	bundleReward.Total = pool.OperatingCost + (bundleProposal.ByteSize * k.StorageCost(ctx))

	// load and parse network fee TODO: how to test network fee?
	networkFee, err := sdk.NewDecFromStr(k.NetworkFee(ctx))
	if err != nil {
		// TODO: panic halt?
		// k.PanicHalt(ctx, "Invalid value for params: "+err.Error())
	}

	_, stakerFound := k.stakerKeeper.GetStaker(ctx, bundleProposal.Uploader)

	if !stakerFound {
		bundleReward.Treasury = bundleReward.Total

		return
	}

	// TODO: check if staker has delegations

	bundleReward.Treasury = uint64(sdk.NewDec(int64(bundleReward.Total)).Mul(networkFee).RoundInt64())
	totalNodeReward := bundleReward.Total - bundleReward.Treasury

	// TODO: check if staker has delegations
	uploaderCommission, err := sdk.NewDecFromStr( /*staker.Commission*/ "1")
	if err != nil {
		// TODO: panic halt?
		// k.PanicHalt(ctx, "Invalid value for params: "+err.Error())
	}

	bundleReward.Uploader = uint64(sdk.NewDec(int64(totalNodeReward)).Mul(uploaderCommission).RoundInt64())
	bundleReward.Delegation = totalNodeReward - bundleReward.Uploader

	return
}

func (k Keeper) finalizeCurrentBundleProposal(ctx sdk.Context, pool poolmoduletypes.Pool, bundleProposal types.BundleProposal, voteDistribution types.VoteDistribution, bundleReward types.BundleReward) error {
	// save finalized bundle
	finalizedBundle := types.FinalizedBundle{
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
	}

	k.SetFinalizedBundle(ctx, finalizedBundle)

	err := ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalized{
		PoolId:           finalizedBundle.PoolId,
		Id:               finalizedBundle.Id,
		Valid:            voteDistribution.Valid,
		Invalid:          voteDistribution.Invalid,
		Abstain:          voteDistribution.Abstain,
		Total:            voteDistribution.Total,
		Status:           voteDistribution.Status,
		RewardTreasury:   bundleReward.Treasury,
		RewardUploader:   bundleReward.Uploader,
		RewardDelegation: bundleReward.Delegation,
		RewardTotal:      bundleReward.Total,
	})

	if err != nil {
		return err
	}

	// Finalize the proposal, saving useful information.
	// eventFromHeight := pool.CurrentHeight
	pool.CurrentHeight = bundleProposal.ToHeight
	pool.TotalBundles = pool.TotalBundles + 1
	pool.CurrentKey = bundleProposal.ToKey
	pool.CurrentValue = bundleProposal.ToValue

	k.poolKeeper.SetPool(ctx, pool)

	return nil
}

func (k Keeper) dropCurrentBundleProposal(ctx sdk.Context, pool poolmoduletypes.Pool, bundleProposal types.BundleProposal, voteDistribution types.VoteDistribution) error {
	err := ctx.EventManager().EmitTypedEvent(&types.EventBundleFinalized{
		PoolId:  pool.Id,
		Id:      pool.TotalBundles,
		Valid:   voteDistribution.Valid,
		Invalid: voteDistribution.Invalid,
		Abstain: voteDistribution.Abstain,
		Total:   voteDistribution.Total,
		Status:  voteDistribution.Status,
	})

	// drop bundle
	bundleProposal = types.BundleProposal{
		NextUploader: bundleProposal.NextUploader,
		CreatedAt:    uint64(ctx.BlockTime().Unix()),
	}

	k.SetBundleProposal(ctx, bundleProposal)

	return err
}

// RandomChoiceCandidate ...
type RandomChoiceCandidate struct {
	Account string
	Amount  uint64
}

// getWeightedRandomChoice is an internal function that returns a random selection out of a list of candidates.
func (k Keeper) getWeightedRandomChoice(candidates []RandomChoiceCandidate, seed uint64) string {
	type WeightedRandomChoice struct {
		Elements    []string
		Weights     []uint64
		TotalWeight uint64
	}

	wrc := WeightedRandomChoice{}

	for _, candidate := range candidates {
		i := sort.Search(len(wrc.Weights), func(i int) bool { return wrc.Weights[i] > candidate.Amount })
		wrc.Weights = append(wrc.Weights, 0)
		wrc.Elements = append(wrc.Elements, "")
		copy(wrc.Weights[i+1:], wrc.Weights[i:])
		copy(wrc.Elements[i+1:], wrc.Elements[i:])
		wrc.Weights[i] = candidate.Amount
		wrc.Elements[i] = candidate.Account
		wrc.TotalWeight += candidate.Amount
	}

	rand.Seed(int64(seed))
	value := uint64(math.Floor(rand.Float64() * float64(wrc.TotalWeight)))

	for key, weight := range wrc.Weights {
		if weight > value {
			return wrc.Elements[key]
		}

		value -= weight
	}

	return ""
}

func (k Keeper) chooseNextUploaderFromSelectedStakers(ctx sdk.Context, poolId uint64, addresses []string) (nextUploader string) {
	var _candidates []RandomChoiceCandidate

	if len(addresses) == 0 {
		return ""
	}

	for _, s := range addresses {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, s)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, s)

		_candidates = append(_candidates, RandomChoiceCandidate{
			Account: s,
			Amount:  stake + delegation,
		})
	}

	return k.getWeightedRandomChoice(_candidates, uint64(ctx.BlockHeight()))
}

func (k Keeper) chooseNextUploaderFromAllStakers(ctx sdk.Context, poolId uint64) (nextUploader string) {
	stakers := k.stakerKeeper.GetAllStakerAddressesOfPool(ctx, poolId)
	return k.chooseNextUploaderFromSelectedStakers(ctx, poolId, stakers)
}

// getVoteDistribution is an internal function evaulates the quorum status of a bundle proposal.
func (k Keeper) getVoteDistribution(ctx sdk.Context, poolId uint64) (voteDistribution types.VoteDistribution) {
	bundleProposal, found := k.GetBundleProposal(ctx, poolId)
	if !found {
		return
	}

	// get $KYVE voted for valid
	for _, voter := range bundleProposal.VotersValid {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		voteDistribution.Valid += stake + delegation
	}

	// get $KYVE voted for invalid
	for _, voter := range bundleProposal.VotersInvalid {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		voteDistribution.Invalid += stake + delegation
	}

	// get $KYVE voted for abstain
	for _, voter := range bundleProposal.VotersAbstain {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		voteDistribution.Abstain += stake + delegation
	}

	// TODO: get total delegation of pool
	voteDistribution.Total = k.stakerKeeper.GetTotalStake(ctx, poolId) + 0

	if voteDistribution.Valid*2 > voteDistribution.Total {
		voteDistribution.Status = types.BUNDLE_STATUS_VALID
	} else if voteDistribution.Invalid*2 >= voteDistribution.Total {
		voteDistribution.Status = types.BUNDLE_STATUS_INVALID
	} else {
		voteDistribution.Status = types.BUNDLE_STATUS_NO_QUORUM
	}

	return
}
