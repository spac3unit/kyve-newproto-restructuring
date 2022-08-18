package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovCreatePool = "gov_create_pool"

var _ sdk.Msg = &GovMsgCreatePool{}

func NewGovMsgCreatePool(
	creator string,
	name string,
	runtime string,
	logo string,
	config string,
	startKey string,
	uploadInterval uint64,
	operatingCost uint64,
	minStake uint64,
	maxBundleSize uint64,
	version string,
	binaries string,
) *GovMsgCreatePool {
	return &GovMsgCreatePool{
		Creator:        creator,
		Name:           name,
		Runtime:        runtime,
		Logo:           logo,
		Config:         config,
		StartKey:       startKey,
		UploadInterval: uploadInterval,
		OperatingCost:  operatingCost,
		MinStake:       minStake,
		MaxBundleSize:  maxBundleSize,
		Version:        version,
		Binaries:       binaries,
	}
}

func (msg *GovMsgCreatePool) Route() string {
	return RouterKey
}

func (msg *GovMsgCreatePool) Type() string {
	return TypeMsgGovCreatePool
}

func (msg *GovMsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
