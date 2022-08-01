package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyUploadTimeout            = []byte("UploadTimeout")
	DefaultUploadTimeout uint64 = 600
)

var (
	KeyStorageCost            = []byte("StorageCost")
	DefaultStorageCost uint64 = 100000
)

var (
	KeyNetworkFee            = []byte("NetworkFee")
	DefaultNetworkFee string = "0.01"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(

	uploadTimeout uint64,
	storageCost uint64,
	networkFee string,

) Params {
	return Params{
		UploadTimeout: uploadTimeout,
		StorageCost:   storageCost,
		NetworkFee:    networkFee,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultUploadTimeout,
		DefaultStorageCost,
		DefaultNetworkFee,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyUploadTimeout, &p.UploadTimeout, validateUploadTimeout),
		paramtypes.NewParamSetPair(KeyStorageCost, &p.StorageCost, validateStorageCost),
		paramtypes.NewParamSetPair(KeyNetworkFee, &p.NetworkFee, validateNetworkFee),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {

	if err := validateUploadTimeout(p.UploadTimeout); err != nil {
		return err
	}

	if err := validateStorageCost(p.StorageCost); err != nil {
		return err
	}

	if err := validateNetworkFee(p.NetworkFee); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateUploadTimeout validates the uploadTimeout param
func validateUploadTimeout(v interface{}) error {
	uploadTimeout, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = uploadTimeout

	return nil
}

// validateStorageCost validates the StorageCost param
func validateStorageCost(v interface{}) error {
	storageCost, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = storageCost

	return nil
}

// validateNetworkFee validates the NetworkFee param
func validateNetworkFee(v interface{}) error {
	return validatePercentage(v)
}

// TODO maybe outsource to util
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
