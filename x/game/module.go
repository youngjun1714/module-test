package game

import (
	"encoding/json"

	gametypes "github.com/youngjun1714/module-test/x/game/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	"github.com/spf13/cobra"
)

var test = gametypes.ModuleName

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
	cdc codec.Marshaler
}

func (AppModuleBasic) Name() string {
	return gametypes.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {

}
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	/* genesis 정의 후 구현 예정. */
	return nil
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {

	return nil
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}
