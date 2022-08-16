package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/util"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// WithdrawRewards handles the logic of an SDK message that allows delegators to collect all rewards from a specified pool.
func (k msgServer) WithdrawRewards(goCtx context.Context, msg *types.MsgWithdrawRewards) (*types.MsgWithdrawRewardsResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is a delegator in this pool.
	_, isDelegator := k.GetDelegator(ctx, msg.Staker, msg.Creator)
	if !isDelegator {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNotADelegator.Error())
	}

	// Withdraw all rewards for the sender.
	reward := k.f1WithdrawRewards(ctx, msg.Staker, msg.Creator)

	// Transfer tokens from this module to sender.
	err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, msg.Creator, reward)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawRewardsResponse{}, nil
}
