package authz

import (
	"github.com/joshklop/monomer-cosmos-sdk/client"
	addresscodec "github.com/joshklop/monomer-cosmos-sdk/codec/address"
	"github.com/joshklop/monomer-cosmos-sdk/testutil"
	clitestutil "github.com/joshklop/monomer-cosmos-sdk/testutil/cli"
	"github.com/joshklop/monomer-cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization(addresscodec.NewBech32Codec("cosmos"))
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
