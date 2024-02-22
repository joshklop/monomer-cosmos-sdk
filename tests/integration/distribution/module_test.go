package distribution_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	simtestutil "github.com/joshklop/monomer-cosmos-sdk/testutil/sims"
	authkeeper "github.com/joshklop/monomer-cosmos-sdk/x/auth/keeper"
	authtypes "github.com/joshklop/monomer-cosmos-sdk/x/auth/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/distribution/testutil"
	"github.com/joshklop/monomer-cosmos-sdk/x/distribution/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(
		depinject.Configs(
			testutil.AppConfig,
			depinject.Supply(log.NewNopLogger()),
		),
		&accountKeeper)
	assert.NilError(t, err)

	ctx := app.BaseApp.NewContext(false)
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
