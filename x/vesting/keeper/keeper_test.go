package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/network"
	testutiltx "github.com/lyfeblocnetwork/lyfebloc/testutil/tx"
	"github.com/lyfeblocnetwork/lyfebloc/x/vesting/keeper"
	vestingtypes "github.com/lyfeblocnetwork/lyfebloc/x/vesting/types"
	"github.com/stretchr/testify/require"
)

func TestNewKeeper(t *testing.T) {
	nw := network.NewUnitTestNetwork()
	encCfg := nw.GetEncodingConfig()
	cdc := encCfg.Codec
	storeKey := storetypes.NewKVStoreKey(vestingtypes.StoreKey)

	addr, _ := testutiltx.NewAccAddressAndKey()

	testcases := []struct {
		name      string
		authority sdk.AccAddress
		expPass   bool
	}{
		{
			name:      "valid authority format",
			authority: addr,
			expPass:   true,
		},
		{
			name:      "empty authority",
			authority: []byte{},
			expPass:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expPass {
				newKeeper := keeper.NewKeeper(
					storeKey,
					tc.authority,
					cdc,
					nw.App.AccountKeeper,
					nw.App.BankKeeper,
					nw.App.DistrKeeper,
					nw.App.StakingKeeper,
					nw.App.GovKeeper,
				)
				require.NotNil(t, newKeeper)
			} else {
				require.PanicsWithError(t, "addresses cannot be empty: unknown address", func() {
					_ = keeper.NewKeeper(
						storeKey,
						tc.authority,
						cdc,
						nw.App.AccountKeeper,
						nw.App.BankKeeper,
						nw.App.DistrKeeper,
						nw.App.StakingKeeper,
						nw.App.GovKeeper,
					)
				})
			}
		})
	}
}