package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBetAmountRequest{}, "game/MsgBetAmountRequest", nil)
}
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBetAmountRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// air session social scrap grab laptop enroll shove educate unaware kingdom honey cloth enemy alarm fun arch output skate hungry target dolphin alien walnut
// game135rhl235p269jkukngw4r288t9uj32m9qnms79

// AtM23bswgQsmPNQKLC/JR3SBv4IfJm/fS2IPB/XusWbX
// gamepub1addwnpepqtfndhdmxzqskf3u6s9zct7fga6gr0uzrunxlh6tvg8s0a0wk9ndw8w7494
//game1hmnxk5d84sl33lmdyxcervdqrp9nfau0xtv9u0
