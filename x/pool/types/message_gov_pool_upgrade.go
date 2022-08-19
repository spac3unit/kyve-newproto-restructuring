package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovPoolUpgrade = "gov_pool_upgrade"

var _ sdk.Msg = &GovMsgPoolUpgrade{}

func NewGovMsgPoolUpgrade(
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

func (msg *GovMsgPoolUpgrade) Route() string {
	return RouterKey
}

func (msg *GovMsgPoolUpgrade) Type() string {
	return TypeMsgGovPoolUpgrade
}

func (msg *GovMsgPoolUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgPoolUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgPoolUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
