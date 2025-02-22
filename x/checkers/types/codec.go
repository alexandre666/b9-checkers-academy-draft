package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGame{}, "checkers/CreateGame", nil)
	cdc.RegisterConcrete(&MsgPlayMove{}, "checkers/PlayMove", nil)
	cdc.RegisterConcrete(&MsgCreateEnergy{}, "checkers/CreateEnergy", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGame{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPlayMove{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEnergy{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
