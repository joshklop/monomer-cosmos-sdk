package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/joshklop/monomer-cosmos-sdk/crypto/types"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/staking/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(tb testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	tb.Helper()
	v, err := types.NewValidator(operator.String(), pubKey, types.Description{})
	require.NoError(tb, err)
	return v
}
