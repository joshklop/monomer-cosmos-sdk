package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/joshklop/monomer-cosmos-sdk/types/module"
	moduletestutil "github.com/joshklop/monomer-cosmos-sdk/types/module/testutil"
	simtypes "github.com/joshklop/monomer-cosmos-sdk/types/simulation"
	"github.com/joshklop/monomer-cosmos-sdk/x/authz"
	authzmodule "github.com/joshklop/monomer-cosmos-sdk/x/authz/module"
	"github.com/joshklop/monomer-cosmos-sdk/x/authz/simulation"
	banktypes "github.com/joshklop/monomer-cosmos-sdk/x/bank/types"
)

func TestRandomizedGenState(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(authzmodule.AppModuleBasic{})
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          encCfg.Codec,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var authzGenesis authz.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[authz.ModuleName], &authzGenesis)

	require.Len(t, authzGenesis.Authorization, len(simState.Accounts)-1)
}
