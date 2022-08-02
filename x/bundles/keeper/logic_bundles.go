package keeper

import (
	"github.com/KYVENetwork/chain/x/bundles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math"
	"math/rand"
	"sort"
)

func containsElement(array []string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

// updateLowestFunder is an internal function that updates the lowest funder entry in a given pool.
func (k Keeper) handleNonVoters(ctx sdk.Context, poolId uint64) {
	//nonVoters := make([]string, 0)
	//
	//for _, staker := range pool.Stakers {
	//	if staker == pool.BundleProposal.Uploader {
	//		continue
	//	}
	//
	//	// TODO improve runtime
	//	valid := containsElement(pool.BundleProposal.VotersValid, staker)
	//	invalid := containsElement(pool.BundleProposal.VotersInvalid, staker)
	//	abstain := containsElement(pool.BundleProposal.VotersAbstain, staker)
	//
	//	if !valid && !invalid && !abstain {
	//		nonVoters = append(nonVoters, staker)
	//	}
	//}
	//
	//for _, voter := range nonVoters {
	//	staker, foundStaker := k.GetStaker(ctx, voter, pool.Id)
	//
	//	// skip timeout slash if staker is not found
	//	if foundStaker {
	//
	//		// TODO think about if slash should happen in staker module or in
	//		if staker.Points < k.MaxPoints(ctx) {
	//			// Increase points
	//			staker.Points += 1
	//			k.SetStaker(ctx, staker)
	//		} else {
	//			// slash nonVoter for not voting in time
	//			slashAmount := k.slashStaker(ctx, pool, staker.Account, k.TimeoutSlash(ctx))
	//
	//			// emit slashing event
	//			ctx.EventManager().EmitTypedEvent(&types.EventSlash{
	//				PoolId:    pool.Id,
	//				Address:   staker.Account,
	//				Amount:    slashAmount,
	//				SlashType: types.SLASH_TYPE_TIMEOUT,
	//			})
	//
	//			// Check if staker is still in stakers list and remove staker.
	//			staker, foundStaker = k.GetStaker(ctx, voter, pool.Id)
	//
	//			// check if next uploader is still there or already removed
	//			if foundStaker {
	//				deactivateStaker(pool, &staker)
	//				k.SetStaker(ctx, staker)
	//
	//				ctx.EventManager().EmitTypedEvent(&types.EventStakerStatusChanged{
	//					PoolId:  pool.Id,
	//					Address: staker.Account,
	//					Status:  types.STAKER_STATUS_INACTIVE,
	//				})
	//			}
	//
	//			// Update current lowest staker
	//			k.updateLowestStaker(ctx, pool)
	//		}
	//	}
	//}
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

func (k Keeper) GetUploadProbability(ctx sdk.Context, stakerAddress string, poolId uint64) sdk.Dec {

	//pool, poolFound := k.GetPool(ctx, poolId)
	//if !poolFound {
	//	return sdk.NewDec(0)
	//}
	//
	//totalWeight := uint64(0)
	//userWeight := uint64(0)
	//
	//for _, s := range pool.Stakers {
	//	staker, _ := k.GetStaker(ctx, s, pool.Id)
	//	delegation, _ := k.GetDelegationPoolData(ctx, pool.Id, s)
	//
	//	totalWeight += staker.Amount + getDelegationWeight(delegation.TotalDelegation)
	//	if staker.Account == stakerAddress {
	//		userWeight = staker.Amount + getDelegationWeight(delegation.TotalDelegation)
	//	}
	//}
	//
	//return sdk.NewDec(int64(userWeight)).Quo(sdk.NewDec(int64(totalWeight)))
	return sdk.Dec{}
}

// Calculate Delegation weight to influnce the upload probability
// formula:
// A = 10000, dec = 10**9
// weight = dec * (sqrt(A * (A + x/dec)) - A)
func getDelegationWeight(delegation uint64) uint64 {

	const A uint64 = 10000

	number := A * (A + (delegation / 1_000_000_000))

	// Deterministic sqrt using only int
	// Uses the babylon recursive formula:
	// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
	var x uint64 = 14142 // expected value for 10000 $KYVE as input
	var xn uint64
	var epsilon uint64 = 100
	for epsilon > 2 {

		xn = (x + number/x) / 2

		if xn > x {
			epsilon = xn - x
		} else {
			epsilon = x - xn
		}
		x = xn
	}

	return (x - A) * 1_000_000_000
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

func (k Keeper) getNextUploader(ctx sdk.Context, poolId uint64) (nextUploader string) {
	stakers := k.stakerKeeper.GetStakerAddressesOfPool(ctx, poolId)
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
