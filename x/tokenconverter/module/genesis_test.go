package tokenconverter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/nullify"
	tokenconverter "github.com/lyfeblocnetwork/lyfebloc/x/tokenconverter/module"
	"github.com/lyfeblocnetwork/lyfebloc/x/tokenconverter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenconverterKeeper(t)
	tokenconverter.InitGenesis(ctx, k, genesisState)
	got := tokenconverter.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
