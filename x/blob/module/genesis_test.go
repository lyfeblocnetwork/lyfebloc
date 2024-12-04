package blob_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/lyfeblocnetwork/lyfebloc/testutil/keeper"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/nullify"
	blob "github.com/lyfeblocnetwork/lyfebloc/x/blob/module"
	"github.com/lyfeblocnetwork/lyfebloc/x/blob/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlobKeeper(t)
	blob.InitGenesis(ctx, k, genesisState)
	got := blob.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
