package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDelegatePool{}, "registry/DelegatePool", nil)
	cdc.RegisterConcrete(&MsgWithdrawPool{}, "registry/WithdrawPool", nil)
	cdc.RegisterConcrete(&MsgUndelegatePool{}, "registry/UndelegatePool", nil)
	cdc.RegisterConcrete(&MsgRedelegatePool{}, "registry/RedelegatePool", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelegatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndelegatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedelegatePool{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
