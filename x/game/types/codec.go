package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	/*	"github.com/cosmos/cosmos-sdk/codec/types"
		cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
		sdk "github.com/cosmos/cosmos-sdk/types"
		"github.com/cosmos/cosmos-sdk/types/msgservice"*/)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBetAmountRequest{}, "game/MsgBetAmountRequest", nil)
}
func RegisterInterfaces (registry types.InterfaceRegistry) {
	
}

/*func init() {

}*/
