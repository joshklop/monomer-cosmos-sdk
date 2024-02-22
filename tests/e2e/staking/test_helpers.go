package testutil

import (
	"fmt"

	"cosmossdk.io/math"

	"github.com/joshklop/monomer-cosmos-sdk/client"
	"github.com/joshklop/monomer-cosmos-sdk/client/flags"
	addresscodec "github.com/joshklop/monomer-cosmos-sdk/codec/address"
	"github.com/joshklop/monomer-cosmos-sdk/testutil"
	clitestutil "github.com/joshklop/monomer-cosmos-sdk/testutil/cli"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	stakingcli "github.com/joshklop/monomer-cosmos-sdk/x/staking/client/cli"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(10))).String()),
}

// MsgRedelegateExec creates a redelegate message.
func MsgRedelegateExec(clientCtx client.Context, from, src, dst, amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		src.String(),
		dst.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
		fmt.Sprintf("--%s=%d", flags.FlagGas, 300000),
	}
	args = append(args, extraArgs...)

	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewRedelegateCmd(addresscodec.NewBech32Codec("cosmosvaloper"), addresscodec.NewBech32Codec("cosmos")), args)
}

// MsgUnbondExec creates a unbond message.
func MsgUnbondExec(clientCtx client.Context, from, valAddress,
	amount fmt.Stringer,
	extraArgs ...string,
) (testutil.BufferWriter, error) {
	args := []string{
		valAddress.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
	}

	args = append(args, commonArgs...)
	args = append(args, extraArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewUnbondCmd(addresscodec.NewBech32Codec("cosmosvaloper"), addresscodec.NewBech32Codec("cosmos")), args)
}
