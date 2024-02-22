package types

import (
	codectypes "github.com/joshklop/monomer-cosmos-sdk/codec/types"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
)

func (m *QueryAccountResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var account sdk.AccountI
	return unpacker.UnpackAny(m.Account, &account)
}

var _ codectypes.UnpackInterfacesMessage = &QueryAccountResponse{}
