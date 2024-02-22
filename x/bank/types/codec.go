package types

import (
	"github.com/joshklop/monomer-cosmos-sdk/codec"
	"github.com/joshklop/monomer-cosmos-sdk/codec/legacy"
	"github.com/joshklop/monomer-cosmos-sdk/codec/types"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/types/msgservice"
	"github.com/joshklop/monomer-cosmos-sdk/x/authz"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgSend{}, "cosmos-sdk/MsgSend")
	legacy.RegisterAminoMsg(cdc, &MsgMultiSend{}, "cosmos-sdk/MsgMultiSend")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/bank/MsgUpdateParams")
	legacy.RegisterAminoMsg(cdc, &MsgSetSendEnabled{}, "cosmos-sdk/MsgSetSendEnabled")

	cdc.RegisterConcrete(&SendAuthorization{}, "cosmos-sdk/SendAuthorization", nil)
	cdc.RegisterConcrete(&Params{}, "cosmos-sdk/x/bank/Params", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
		&MsgUpdateParams{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SendAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
