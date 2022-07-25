package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgStakePool{}, "registry/StakePool", nil)
	cdc.RegisterConcrete(&MsgUnstakePool{}, "registry/UnstakePool", nil)
	cdc.RegisterConcrete(&MsgUpdateMetadata{}, "registry/UpdateMetadata", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateMetadata{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStakePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnstakePool{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
