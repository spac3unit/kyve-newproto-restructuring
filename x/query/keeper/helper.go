package keeper

import (
	"github.com/KYVENetwork/chain/x/query/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getFullStaker(ctx sdk.Context, stakerAddress string) *types.FullStaker {
	staker, _ := k.stakerKeeper.GetStaker(ctx, stakerAddress)

	commissionChange, found := k.stakerKeeper.GetCommissionChangeEntryByIndex2(ctx, staker.Address)
	var commissionChangeEntry *types.CommissionChangeEntry = nil
	if found {
		commissionChangeEntry = &types.CommissionChangeEntry{
			Commission:   commissionChange.Commission,
			CreationDate: commissionChange.CreationDate,
		}
	}

	stakerMetadata := types.StakerMetadata{
		Commission:              staker.Commission,
		Moniker:                 staker.Moniker,
		Website:                 staker.Website,
		Logo:                    staker.Logo,
		PendingCommissionChange: commissionChangeEntry,
	}

	delegationData, _ := k.delegationKeeper.GetDelegationData(ctx, staker.Address)

	var poolMemberships []*types.PoolMembership

	for _, valaccount := range k.stakerKeeper.GetValaccountsFromStaker(ctx, staker.Address) {

		pool, _ := k.poolKeeper.GetPool(ctx, valaccount.PoolId)

		poolMemberships = append(poolMemberships, &types.PoolMembership{
			Pool: &types.BasicPool{
				Id:         pool.Id,
				Name:       pool.Name,
				Runtime:    pool.Runtime,
				Logo:       pool.Logo,
				TotalFunds: pool.TotalFunds,
			},
			Points:     valaccount.Points,
			IsLeaving:  valaccount.IsLeaving,
			Valaccount: valaccount.Valaddress,
		})
	}

	return &types.FullStaker{
		Address:         staker.Address,
		Metadata:        &stakerMetadata,
		Amount:          staker.Amount,
		UnbondingAmount: staker.UnbondingAmount,
		TotalDelegation: k.delegationKeeper.GetDelegationAmount(ctx, staker.Address),
		DelegatorCount:  delegationData.DelegatorCount,
		Pools:           poolMemberships,
	}
}
