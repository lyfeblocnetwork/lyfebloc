package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/x/tokenconverter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TokenconverterKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
