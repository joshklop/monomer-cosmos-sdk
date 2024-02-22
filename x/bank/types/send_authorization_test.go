package types_test

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/core/header"
	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"github.com/joshklop/monomer-cosmos-sdk/testutil"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/monomer-cosmos-sdk/x/bank/types"
)

var (
	coins1000   = sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(1000)))
	coins500    = sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(500)))
	fromAddr    = sdk.AccAddress("_____from _____")
	toAddr      = sdk.AccAddress("_______to________")
	unknownAddr = sdk.AccAddress("_____unknown_____")
)

func TestSendAuthorization(t *testing.T) {
	ctx := testutil.DefaultContextWithDB(t, storetypes.NewKVStoreKey(types.StoreKey), storetypes.NewTransientStoreKey("transient_test")).Ctx.WithHeaderInfo(header.Info{})
	allowList := make([]sdk.AccAddress, 1)
	allowList[0] = toAddr
	authorization := types.NewSendAuthorization(coins1000, nil)

	t.Log("verify authorization returns valid method name")
	require.Equal(t, authorization.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send := types.NewMsgSend(fromAddr, toAddr, coins1000)

	require.NoError(t, authorization.ValidateBasic())

	t.Log("verify updated authorization returns nil")
	resp, err := authorization.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Delete)
	require.Nil(t, resp.Updated)

	authorization = types.NewSendAuthorization(coins1000, nil)
	require.Equal(t, authorization.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddr, toAddr, coins500)
	require.NoError(t, authorization.ValidateBasic())
	resp, err = authorization.Accept(ctx, send)

	t.Log("verify updated authorization returns remaining spent limit")
	require.NoError(t, err)
	require.False(t, resp.Delete)
	require.NotNil(t, resp.Updated)
	sendAuth := types.NewSendAuthorization(coins500, nil)
	require.Equal(t, sendAuth.String(), resp.Updated.String())

	t.Log("expect updated authorization nil after spending remaining amount")
	resp, err = resp.Updated.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Delete)
	require.Nil(t, resp.Updated)

	t.Log("allow list and no address")
	authzWithAllowList := types.NewSendAuthorization(coins1000, allowList)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddr, unknownAddr, coins500)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.False(t, resp.Accept)
	require.False(t, resp.Delete)
	require.Nil(t, resp.Updated)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("cannot send to %s address", unknownAddr))

	t.Log("send to address in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddr, allowList[0], coins500)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Accept)
	require.NotNil(t, resp.Updated)
	// coins1000-coins500 = coins500
	require.Equal(t, types.NewSendAuthorization(coins500, allowList).String(), resp.Updated.String())

	t.Log("send everything to address not in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddr, unknownAddr, coins1000)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.Error(t, err)
	require.Contains(t, err.Error(), fmt.Sprintf("cannot send to %s address", unknownAddr))

	t.Log("send everything to address in allow list")
	authzWithAllowList = types.NewSendAuthorization(coins1000, allowList)
	require.Equal(t, authzWithAllowList.MsgTypeURL(), "/cosmos.bank.v1beta1.MsgSend")
	require.NoError(t, authorization.ValidateBasic())
	send = types.NewMsgSend(fromAddr, allowList[0], coins1000)
	require.NoError(t, authzWithAllowList.ValidateBasic())
	resp, err = authzWithAllowList.Accept(ctx, send)
	require.NoError(t, err)
	require.True(t, resp.Accept)
	require.Nil(t, resp.Updated)
}
