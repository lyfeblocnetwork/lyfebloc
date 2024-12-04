package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquidityincentive/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquidityincentive/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, keepertest.LiquidityIncentiveMocks, types.MsgServer, context.Context) {
	k, mocks, ctx := keepertest.LiquidityincentiveKeeper(t)
	return k, mocks, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, _, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
