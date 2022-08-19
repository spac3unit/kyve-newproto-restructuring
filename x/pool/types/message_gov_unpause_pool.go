package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovUnpausePool = "gov_pause_pool_upgrade"

var _ sdk.Msg = &GovMsgUnpausePool{}

func NewGovUnpausePool(
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

func (msg *GovMsgUnpausePool) Route() string {
	return RouterKey
}

func (msg *GovMsgUnpausePool) Type() string {
	return TypeMsgGovUnpausePool
}

func (msg *GovMsgUnpausePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgUnpausePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgUnpausePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
