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
	cdc.RegisterConcrete(&MsgFundPool{}, "registry/FundPool", nil)
	cdc.RegisterConcrete(&MsgDefundPool{}, "registry/DefundPool", nil)

	cdc.RegisterConcrete(&GovMsgCreatePoolResponse{}, "kyve/CreatePool", nil)
	cdc.RegisterConcrete(&GovMsgUpdatePoolResponse{}, "kyve/UpdatePool", nil)
	//cdc.RegisterConcrete(&UpdatePoolProposal{}, "kyve/UpdatePoolProposal", nil)
	//cdc.RegisterConcrete(&PausePoolProposal{}, "kyve/PausePoolProposal", nil)
	//cdc.RegisterConcrete(&UnpausePoolProposal{}, "kyve/UnpausePoolProposal", nil)
	//cdc.RegisterConcrete(&SchedulePoolUpgradeProposal{}, "kyve/SchedulePoolUpgradeProposal", nil)
	//cdc.RegisterConcrete(&CancelPoolUpgradeProposal{}, "kyve/CancelPoolUpgradeProposal", nil)
	//cdc.RegisterConcrete(&ResetPoolProposal{}, "kyve/ResetPoolProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFundPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDefundPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&GovMsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&GovMsgUpdatePool{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
