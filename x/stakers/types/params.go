package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyVoteSlash            = []byte("VoteSlash")
	DefaultVoteSlash string = "0.1"
)

var (
	KeyUploadSlash            = []byte("UploadSlash")
	DefaultUploadSlash string = "0.2"
)

var (
	KeyTimeoutSlash            = []byte("TimeoutSlash")
	DefaultTimeoutSlash string = "0.02"
)

var (
	KeyMaxPoints            = []byte("MaxPoints")
	DefaultMaxPoints uint64 = 5
)

var (
	KeyUnbondingStakingTime            = []byte("UnbondingStakingTime")
	DefaultUnbondingStakingTime uint64 = 60 * 60 * 24 * 5
)

var (
	KeyCommissionChangeTime            = []byte("KeyCommissionChangeTime")
	DefaultCommissionChangeTime uint64 = 60 * 60 * 24 * 5
)

var (
	KeyLeavePoolTime            = []byte("KeyLeavePoolTime")
	DefaultLeavePoolTime uint64 = 60 * 60 * 24 * 5
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	voteSlash string,
	uploadSlash string,
	timeoutSlash string,
	maxPoints uint64,
	unbondingStakingTime uint64,
	commissionChangeTime uint64,
	leavePoolTime uint64,
) Params {
	return Params{
		VoteSlash:            voteSlash,
		UploadSlash:          uploadSlash,
		TimeoutSlash:         timeoutSlash,
		MaxPoints:            maxPoints,
		UnbondingStakingTime: unbondingStakingTime,
		CommissionChangeTime: commissionChangeTime,
		LeavePoolTime:        leavePoolTime,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultVoteSlash,
		DefaultUploadSlash,
		DefaultTimeoutSlash,
		DefaultMaxPoints,
		DefaultUnbondingStakingTime,
		DefaultCommissionChangeTime,
		DefaultLeavePoolTime,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVoteSlash, &p.VoteSlash, validateVoteSlash),
		paramtypes.NewParamSetPair(KeyUploadSlash, &p.UploadSlash, validateUploadSlash),
		paramtypes.NewParamSetPair(KeyTimeoutSlash, &p.TimeoutSlash, validateTimeoutSlash),
		paramtypes.NewParamSetPair(KeyMaxPoints, &p.MaxPoints, validateMaxPoints),
		paramtypes.NewParamSetPair(KeyUnbondingStakingTime, &p.UnbondingStakingTime, validateUnbondingStakingTime),
		paramtypes.NewParamSetPair(KeyCommissionChangeTime, &p.CommissionChangeTime, validateTrue),
		paramtypes.NewParamSetPair(KeyLeavePoolTime, &p.LeavePoolTime, validateTrue),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateTrue(v interface{}) error {
	return nil
}

// validateVoteSlash validates the VoteSlash param
func validateVoteSlash(v interface{}) error {
	return validatePercentage(v)
}

// validateUploadSlash validates the UploadSlash param
func validateUploadSlash(v interface{}) error {
	return validatePercentage(v)
}

// validateTimeoutSlash validates the TimeoutSlash param
func validateTimeoutSlash(v interface{}) error {
	return validatePercentage(v)
}

// validateMaxPoints validates the MaxPoints param
func validateMaxPoints(v interface{}) error {
	maxPoints, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxPoints

	return nil
}

// validateMaxPoints validates the MaxPoints param
func validateUnbondingStakingTime(v interface{}) error {
	_, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}

// validatePercentage ...
func validatePercentage(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	parsedVal, err := sdk.NewDecFromStr(val)
	if err != nil {
		return fmt.Errorf("invalid decimal representation: %T", v)
	}

	if parsedVal.LT(sdk.NewDec(0)) {
		return fmt.Errorf("percentage should be greater than or equal to 0")
	}
	if parsedVal.GT(sdk.NewDec(1)) {
		return fmt.Errorf("percentage should be less than or equal to 1")
	}

	return nil
}
