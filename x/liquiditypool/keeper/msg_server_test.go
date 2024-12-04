package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquiditypool/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquiditypool/testutil"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquiditypool/types"
)

func setupMsgServer(t *testing.T) (keeper.Keeper, *testutil.MockBankKeeper, types.MsgServer, context.Context) {
	k, bk, ctx := keepertest.LiquiditypoolKeeper(t)
	return k, bk, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, _, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
