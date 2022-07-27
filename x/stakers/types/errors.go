package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// staking errors
var (
	ErrStakeTooLow    = sdkerrors.Register(ModuleName, 1103, "minimum staking amount of %vkyve not reached")
	ErrUnstakeTooHigh = sdkerrors.Register(ModuleName, 1104, "maximum unstaking amount of %vkyve surpassed")
	ErrNoStaker       = sdkerrors.Register(ModuleName, 1105, "sender is no staker")

	ErrInvalidCommission          = sdkerrors.Register(ModuleName, 1116, "invalid commission %v")
	ErrPoolLeaveAlreadyInProgress = sdkerrors.Register(ModuleName, 1117, "Pool leave is already in progress")
)
