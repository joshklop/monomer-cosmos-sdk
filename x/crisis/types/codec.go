package types

import (
	"github.com/joshklop/monomer-cosmos-sdk/codec"
	"github.com/joshklop/monomer-cosmos-sdk/codec/legacy"
	codectypes "github.com/joshklop/monomer-cosmos-sdk/codec/types"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/crisis interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgVerifyInvariant{}, "cosmos-sdk/MsgVerifyInvariant")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/crisis/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the Interface Registry.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifyInvariant{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
