package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGovPausePool = "gov_pause_pool_upgrade"

var _ sdk.Msg = &GovMsgPausePool{}

func NewGovPausePool(
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

func (msg *GovMsgPausePool) Route() string {
	return RouterKey
}

func (msg *GovMsgPausePool) Type() string {
	return TypeMsgGovPausePool
}

func (msg *GovMsgPausePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *GovMsgPausePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *GovMsgPausePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
