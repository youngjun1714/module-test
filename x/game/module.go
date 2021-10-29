package game

import (
	"context"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/youngjun1714/module-test/x/game/types"
	gametypes "github.com/youngjun1714/module-test/x/game/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/spf13/cobra"
	"github.com/youngjun1714/module-test/x/game/client/cli"
	"github.com/youngjun1714/module-test/x/game/client/rest"
	"github.com/youngjun1714/module-test/x/game/keeper"

	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
)

type AppModuleBasic struct {
	cdc codec.Marshaler
}

func (AppModuleBasic) Name() string {
	return gametypes.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
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
	return cli.NewQueryCmd()
}

func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	rest.RegisterHandlers(clientCtx, rtr)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	bankKepper types.BankKeeper
}

func NewAppModule(
	cdc codec.Marshaler,
	keeper keeper.Keeper,
	bankKeeper types.BankKeeper,
) AppModule {

	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		bankKepper:     bankKeeper,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	// 구현 예정.
	return nil
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	// 구현 예정.
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	// 구현 예정.
	return nil
}
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the tokenswap module.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}
