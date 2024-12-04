package testfactory_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/lyfeblocnetwork/lyfebloc/app/encoding"
	"github.com/lyfeblocnetwork/lyfebloc/test/util/testfactory"
	"github.com/lyfeblocnetwork/lyfebloc/test/util/testnode"

	testencoding "github.com/lyfeblocnetwork/lyfebloc/test/util/encoding"
)

func TestTestAccount(t *testing.T) {
	enc := encoding.MakeConfig(testencoding.ModuleEncodingRegisters...)
	kr := testfactory.TestKeyring(enc.Codec)
	record, err := kr.Key(testfactory.TestAccName)
	require.NoError(t, err)
	addr, err := record.GetAddress()
	require.NoError(t, err)
	require.Equal(t, testfactory.TestAccAddr, addr.String())
	require.Equal(t, testnode.TestAddress(), addr)
}
