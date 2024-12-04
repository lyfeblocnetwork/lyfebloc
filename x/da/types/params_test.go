package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/lyfeblocnetwork/lyfebloc/x/da/types"
)

func TestGenerateParams(t *testing.T) {
	provingKeyBase64, verifyingKeyBase64 := types.GenerateZkpKeys()
	require.NotNil(t, provingKeyBase64)
	require.NotNil(t, verifyingKeyBase64)
}
