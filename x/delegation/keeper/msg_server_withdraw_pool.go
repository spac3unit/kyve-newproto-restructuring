package keeper

import (
	"context"
	"github.com/KYVENetwork/chain/util"

	"github.com/KYVENetwork/chain/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// WithdrawPool handles the logic of an SDK message that allows delegators to collect all rewards from a specified pool.
func (k msgServer) WithdrawPool(goCtx context.Context, msg *types.MsgWithdrawPool) (*types.MsgWithdrawPoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)

	if poolErr := k.poolKeeper.AssertPoolExists(ctx, msg.PoolId); poolErr != nil {
		return nil, poolErr
	}

	// Check if the sender is a delegator in this pool.
	_, isDelegator := k.GetDelegator(ctx, msg.PoolId, msg.Staker, msg.Creator)
	if !isDelegator {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrNotADelegator.Error())
	}

	// Create a new F1Distribution struct for interacting with delegations.
	f1Distribution := F1Distribution{
		k:                k.Keeper,
		ctx:              ctx,
		poolId:           msg.PoolId,
		stakerAddress:    msg.Staker,
		delegatorAddress: msg.Creator,
	}

	// Withdraw all rewards for the sender.
	reward := f1Distribution.Withdraw()

	// Transfer tokens from this module to sender.
	err := util.TransferToAddress(k.bankKeeper, ctx, types.ModuleName, msg.Creator, reward)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawPoolResponse{}, nil
}
