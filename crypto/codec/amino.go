package codec

import (
	"github.com/cometbft/cometbft/crypto/sr25519"

	"github.com/joshklop/monomer-cosmos-sdk/codec"
	"github.com/joshklop/monomer-cosmos-sdk/crypto/keys/ed25519"
	kmultisig "github.com/joshklop/monomer-cosmos-sdk/crypto/keys/multisig"
	"github.com/joshklop/monomer-cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/joshklop/monomer-cosmos-sdk/crypto/types"
)

// RegisterCrypto registers all crypto dependency types with the provided Amino
// codec.
func RegisterCrypto(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*cryptotypes.PubKey)(nil), nil)
	cdc.RegisterConcrete(sr25519.PubKey{},
		sr25519.PubKeyName, nil)
	cdc.RegisterConcrete(&ed25519.PubKey{},
		ed25519.PubKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PubKey{},
		secp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&kmultisig.LegacyAminoPubKey{},
		kmultisig.PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*cryptotypes.PrivKey)(nil), nil)
	cdc.RegisterConcrete(sr25519.PrivKey{},
		sr25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&ed25519.PrivKey{},
		ed25519.PrivKeyName, nil)
	cdc.RegisterConcrete(&secp256k1.PrivKey{},
		secp256k1.PrivKeyName, nil)
}
