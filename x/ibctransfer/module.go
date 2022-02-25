package ibctransfer

import (
	"encoding/json"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"

	// "context"

	"github.com/cosmos/cosmos-sdk/client"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	sdktransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	sdktransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Sifchain/sifnode/x/ibctransfer/keeper"
	tokenregistrytypes "github.com/Sifchain/sifnode/x/tokenregistry/types"
)

// Type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	_ porttypes.IBCModule   = IBCModule{}
)

// AppModuleBasic defines the basic application module.
type AppModuleBasic struct {
	cosmosAppModule transfer.AppModule
}

func (am AppModuleBasic) Name() string {
	return am.cosmosAppModule.Name()
}

func (am AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) { //nolint
	am.cosmosAppModule.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (am AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	am.cosmosAppModule.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return am.cosmosAppModule.DefaultGenesis(cdc)
}

// ValidateGenesis performs genesis state validation for the module.
func (am AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return am.cosmosAppModule.ValidateGenesis(cdc, config, bz)
}

// RegisterRESTRoutes registers the REST routes for the module.
func (am AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
	am.cosmosAppModule.RegisterRESTRoutes(ctx, rtr)
}

func (am AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	am.cosmosAppModule.RegisterGRPCGatewayRoutes(clientCtx, mux)
}

// GetTxCmd returns the root tx command for the module.
func (am AppModuleBasic) GetTxCmd() *cobra.Command {
	// Append local TX cmd to this if required
	return am.cosmosAppModule.GetTxCmd()
}

// GetQueryCmd returns no root query command for the module.
func (am AppModuleBasic) GetQueryCmd() *cobra.Command {
	// Append local TX cmd to this if required
	return am.cosmosAppModule.GetQueryCmd()
}

//____________________________________________________________________________

// AppModule implements an application module for the ibctransfer module.
type AppModule struct {
	AppModuleBasic
	sdkTransferKeeper sdktransferkeeper.Keeper
	whitelistKeeper   tokenregistrytypes.Keeper
	bankKeeper        bankkeeper.Keeper
	cdc               codec.BinaryCodec
}

func NewAppModule(sdkTransferKeeper sdktransferkeeper.Keeper, whitelistKeeper tokenregistrytypes.Keeper, bankKeeper bankkeeper.Keeper, cdc codec.BinaryCodec) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{
			cosmosAppModule: transfer.NewAppModule(sdkTransferKeeper),
		},
		sdkTransferKeeper: sdkTransferKeeper,
		bankKeeper:        bankKeeper,
		whitelistKeeper:   whitelistKeeper,
		cdc:               cdc,
	}
}

// IBC does not support a legacy querier
func (am AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier { //nolint
	return am.cosmosAppModule.LegacyQuerierHandler(amino)
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	sdktransfertypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.sdkTransferKeeper, am.bankKeeper, am.whitelistKeeper))
	sdktransfertypes.RegisterQueryServer(cfg.QueryServer(), am.sdkTransferKeeper)
}

// Name returns the dispensation module's name.
func (am AppModule) Name() string {
	return am.cosmosAppModule.Name()
}

// RegisterInvariants registers the dispensation module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	am.cosmosAppModule.RegisterInvariants(ir)
}

// Route returns the message routing key for the dispensation module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(sdktransfertypes.RouterKey, nil)
}

// QuerierRoute returns the dispensation module's querier route name.
func (am AppModule) QuerierRoute() string {
	return am.cosmosAppModule.QuerierRoute()
}

// InitGenesis performs genesis initialization for the dispensation module. It returns
// no validator updates
func (am AppModule) InitGenesis(ctx sdk.Context, codec codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	return am.cosmosAppModule.InitGenesis(ctx, codec, data)
}

// ExportGenesis returns the exported genesis state as raw bytes for the dispensation
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, codec codec.JSONCodec) json.RawMessage {
	return am.cosmosAppModule.ExportGenesis(ctx, codec)
}

// BeginBlock returns the begin blocker for the dispensation module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	am.cosmosAppModule.BeginBlock(ctx, req)
}

// EndBlock returns the end blocker for the dispensation module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return am.cosmosAppModule.EndBlock(ctx, req)
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

type IBCModule struct {
	cosmosIBCModule   transfer.IBCModule
	sdkTransferKeeper sdktransferkeeper.Keeper
	whitelistKeeper   tokenregistrytypes.Keeper
	bankKeeper        bankkeeper.Keeper
	cdc               codec.BinaryCodec
}

func NewIBCModule(sdkTransferKeeper sdktransferkeeper.Keeper, whitelistKeeper tokenregistrytypes.Keeper, bankKeeper bankkeeper.Keeper, cdc codec.BinaryCodec) IBCModule {
	return IBCModule{
		cosmosIBCModule:   transfer.NewIBCModule(sdkTransferKeeper),
		sdkTransferKeeper: sdkTransferKeeper,
		bankKeeper:        bankKeeper,
		whitelistKeeper:   whitelistKeeper,
		cdc:               cdc,
	}
}

func (am IBCModule) OnChanOpenInit(ctx sdk.Context, order types.Order, connectionHops []string, portID string, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, version string) error {
	return am.cosmosIBCModule.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, version)
}

func (am IBCModule) OnChanOpenTry(ctx sdk.Context, order types.Order, connectionHops []string, portID, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, counterpartyVersion string) (string, error) {
	return am.cosmosIBCModule.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

func (am IBCModule) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyVersion string) error {
	return am.cosmosIBCModule.OnChanOpenAck(ctx, portID, channelID, counterpartyVersion)
}

func (am IBCModule) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return am.cosmosIBCModule.OnChanOpenConfirm(ctx, portID, channelID)
}

func (am IBCModule) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return am.cosmosIBCModule.OnChanCloseInit(ctx, portID, channelID)
}

func (am IBCModule) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return am.cosmosIBCModule.OnChanOpenConfirm(ctx, portID, channelID)
}

func (am IBCModule) OnRecvPacket(ctx sdk.Context, packet types.Packet, _ sdk.AccAddress) exported.Acknowledgement {
	return OnRecvPacketWhitelistConvert(ctx, am.sdkTransferKeeper, am.whitelistKeeper, am.bankKeeper, packet)
}

func (am IBCModule) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	return am.cosmosIBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

func (am IBCModule) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress) error {
	return am.cosmosIBCModule.OnTimeoutPacket(ctx, packet, relayer)
}
