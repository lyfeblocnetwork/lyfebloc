package da_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/nullify"
	da "github.com/lyfeblocnetwork/lyfebloc/x/da/module"
	"github.com/lyfeblocnetwork/lyfebloc/x/da/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, _, _, _, ctx := keepertest.DaKeeper(t)
	da.InitGenesis(ctx, k, genesisState)
	got := da.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
