package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNotADelegator                   = sdkerrors.Register(ModuleName, 1000, "not a delegator")
	ErrNotEnoughDelegation             = sdkerrors.Register(ModuleName, 1001, "undelegate-amount is larger than current delegation")
	ErrRedelegationOnCooldown          = sdkerrors.Register(ModuleName, 1002, "all redelegation slots are on cooldown")
	ErrMultipleRedelegationInSameBlock = sdkerrors.Register(ModuleName, 1003, "only one redelegation per delegator per block")
	ErrSelfDelegation                  = sdkerrors.Register(ModuleName, 1117, "self delegation not allowed")
)
