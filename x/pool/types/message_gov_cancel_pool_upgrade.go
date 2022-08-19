package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovCancelPoolUpgrade = "gov_cancel_pool_upgrade"

var _ sdk.Msg = &GovMsgCancelPoolUpgrade{}

func NewGovMsgCancelPoolUpgrade(
	creator string,
	id uint64,
	payload string,
) *GovMsgUpdatePool {
	return &GovMsgUpdatePool{
		Creator: creator,
		Id:      id,
		Payload: payload,
	}
}

func (msg *GovMsgCancelPoolUpgrade) Route() string {
	return RouterKey
}

func (msg *GovMsgCancelPoolUpgrade) Type() string {
	return TypeMsgGovCancelPoolUpgrade
}

func (msg *GovMsgCancelPoolUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgCancelPoolUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgCancelPoolUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
