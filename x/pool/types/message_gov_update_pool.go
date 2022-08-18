package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovUpdatePool = "gov_update_pool"

var _ sdk.Msg = &GovMsgUpdatePool{}

func NewGovMsgUpdatePool(
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

func (msg *GovMsgUpdatePool) Route() string {
	return RouterKey
}

func (msg *GovMsgUpdatePool) Type() string {
	return TypeMsgGovUpdatePool
}

func (msg *GovMsgUpdatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgUpdatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
