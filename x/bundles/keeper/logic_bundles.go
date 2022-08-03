package keeper

import (
	"math"
	"math/rand"
	"sort"
	"strings"

	"github.com/KYVENetwork/chain/x/bundles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func containsElement(array []string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

func (k Keeper) validateSubmitBundleArgs(ctx sdk.Context, bundleProposal *types.BundleProposal, msg *types.MsgSubmitBundleProposal) (error) {
	pool, _ := k.poolKeeper.GetPool(ctx, msg.PoolId)
	

	current_height := bundleProposal.ToHeight
	current_key := bundleProposal.ToKey

	

	// Validate storage id
	if msg.StorageId == "" {
		return types.ErrInvalidArgs
	}

	// Check if the sender is the designated uploader.
	if bundleProposal.NextUploader != msg.Creator {
		return types.ErrNotDesignatedUploader
	}

	// Validate upload interval has been surpassed
	if uint64(ctx.BlockTime().Unix()) < (bundleProposal.CreatedAt + pool.UploadInterval) {
		return types.ErrUploadInterval
	}

	// Validate if bundle is not too big
	if msg.ToHeight-current_height > pool.MaxBundleSize {
		return types.ErrMaxBundleSize
	}

	// Validate from height
	if msg.FromHeight != current_height {
		return types.ErrFromHeight
	}

	// Validate to height
	if msg.ToHeight < current_height {
		return types.ErrToHeight
	}

	// Validate from key
	if msg.FromKey != current_key {
		return types.ErrFromKey
	}

	// check if bundle is of type no data bundle
	if strings.HasPrefix(msg.StorageId, types.KYVE_NO_DATA_BUNDLE) {
		// Validate bundle args
		if msg.ToHeight != current_height || msg.ByteSize != 0 {
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
		if msg.ToHeight <= current_height || msg.ByteSize == 0 {
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

func (k Keeper) registerBundleProposal(ctx sdk.Context, bundleProposal types.BundleProposal, msg *types.MsgSubmitBundleProposal, nextUploader string) {
	bundleProposal = types.BundleProposal{
		PoolId:       msg.PoolId,
		Uploader:     msg.Staker,
		NextUploader: nextUploader,
		StorageId:    msg.StorageId,
		ByteSize:     msg.ByteSize,
		ToHeight:     msg.ToHeight,
		CreatedAt:    uint64(ctx.BlockTime().Unix()),
		VotersValid:  append(make([]string, 0), msg.Staker),
		ToKey:        msg.ToKey,
		ToValue:      msg.ToValue,
		BundleHash:   msg.BundleHash,
	}

	k.SetBundleProposal(ctx, bundleProposal)
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

func (k Keeper) getNextUploaderFromAddresses(ctx sdk.Context, poolId uint64, addresses []string) (nextUploader string) {
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

	return k.getWeightedRandomChoice(_candidates, uint64(ctx.BlockHeight()+ctx.BlockTime().Unix()))
}

func (k Keeper) chooseNextUploaderFromAllStakers(ctx sdk.Context, poolId uint64) (nextUploader string) {
	stakers := k.stakerKeeper.GetAllStakerAddressesOfPool(ctx, poolId)
	return k.getNextUploaderFromAddresses(ctx, poolId, stakers)
}

// getVoteDistribution is an internal function evaulates the quorum status of a bundle proposal.
func (k Keeper) getVoteDistribution(ctx sdk.Context, poolId uint64) (valid uint64, invalid uint64, abstain uint64, total uint64) {
	bundleProposal, found := k.GetBundleProposal(ctx, poolId)
	if !found {
		return
	}

	// get $KYVE voted for valid
	for _, voter := range bundleProposal.VotersValid {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		valid += stake + delegation
	}

	// get $KYVE voted for invalid
	for _, voter := range bundleProposal.VotersInvalid {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		invalid += stake + delegation
	}

	// get $KYVE voted for abstain
	for _, voter := range bundleProposal.VotersAbstain {
		stake := k.stakerKeeper.GetStakeInPool(ctx, poolId, voter)
		delegation := k.delegationKeeper.GetDelegationAmount(ctx, voter)
		abstain += stake + delegation
	}

	// subtract uploader stake because he can not vote
	// TODO get voting power
	//total = k.stakerKeeper.GetTotalStake(ctx, pool.Id) - k.stakerKeeper.GetActiveStake(ctx, pool.Id, pool.BundleProposal.Uploader)

	return
}

// getQuorumStatus is an internal function evaulates if quorum was reached on a bundle proposal.
func (k Keeper) getQuorumStatus(valid uint64, invalid uint64, abstain uint64, total uint64) (quorum types.BundleStatus) {
	if valid*2 > total {
		return types.BUNDLE_STATUS_VALID
	}

	if invalid*2 >= total {
		return types.BUNDLE_STATUS_INVALID
	}

	return types.BUNDLE_STATUS_NO_QUORUM
}
