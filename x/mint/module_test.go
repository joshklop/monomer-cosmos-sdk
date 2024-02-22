package mint_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	simtestutil "github.com/joshklop/monomer-cosmos-sdk/testutil/sims"
	authkeeper "github.com/joshklop/monomer-cosmos-sdk/x/auth/keeper"
	authtypes "github.com/joshklop/monomer-cosmos-sdk/x/auth/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/mint/testutil"
	"github.com/joshklop/monomer-cosmos-sdk/x/mint/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(
		depinject.Configs(
			testutil.AppConfig,
			depinject.Supply(log.NewNopLogger()),
		), &accountKeeper)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false)
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
