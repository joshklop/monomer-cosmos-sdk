package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"

	"github.com/joshklop/monomer-cosmos-sdk/client"
	"github.com/joshklop/monomer-cosmos-sdk/client/flags"
	"github.com/joshklop/monomer-cosmos-sdk/server"
	"github.com/joshklop/monomer-cosmos-sdk/types/module"
	moduletestutil "github.com/joshklop/monomer-cosmos-sdk/types/module/testutil"
	"github.com/joshklop/monomer-cosmos-sdk/x/auth"
	"github.com/joshklop/monomer-cosmos-sdk/x/bank"
	banktypes "github.com/joshklop/monomer-cosmos-sdk/x/bank/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/consensus"
	"github.com/joshklop/monomer-cosmos-sdk/x/distribution"
	"github.com/joshklop/monomer-cosmos-sdk/x/genutil"
	genutiltest "github.com/joshklop/monomer-cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/joshklop/monomer-cosmos-sdk/x/genutil/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/mint"
	"github.com/joshklop/monomer-cosmos-sdk/x/params"
	"github.com/joshklop/monomer-cosmos-sdk/x/staking"
)

func Test_TestnetCmd(t *testing.T) {
	moduleBasic := module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		params.AppModuleBasic{},
		consensus.AppModuleBasic{},
	)

	home := t.TempDir()
	encodingConfig := moduletestutil.MakeTestEncodingConfig(auth.AppModuleBasic{}, staking.AppModuleBasic{})
	logger := log.NewNopLogger()
	cfg, err := genutiltest.CreateDefaultCometConfig(home)
	require.NoError(t, err)

	err = genutiltest.ExecInitCmd(moduleBasic, home, encodingConfig.Codec)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithHomeDir(home).
		WithTxConfig(encodingConfig.TxConfig)

	ctx := context.Background()
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := testnetInitFilesCmd(moduleBasic, banktypes.GenesisBalancesIterator{})
	cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend), fmt.Sprintf("--output-dir=%s", home)})
	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)

	genFile := cfg.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)

	bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
	require.NotEmpty(t, bankGenState.Supply.String())
}
