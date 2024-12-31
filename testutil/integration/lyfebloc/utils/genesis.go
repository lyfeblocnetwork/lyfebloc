// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)

package utils

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	testkeyring "github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/keyring"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/network"
	utiltx "github.com/lyfeblocnetwork/lyfebloc/testutil/tx"
	"github.com/lyfeblocnetwork/lyfebloc/utils"
	erc20types "github.com/lyfeblocnetwork/lyfebloc/x/erc20/types"
	evmtypes "github.com/lyfeblocnetwork/lyfebloc/x/evm/types"
)

// CreateGenesisWithTokenPairs creates a genesis that includes
// the WLYFEBLOC and the provided denoms.
// If no denoms provided, creates only one dynamic precompile with the 'xmpl' denom.
func CreateGenesisWithTokenPairs(keyring testkeyring.Keyring, denoms ...string) network.CustomGenesisState {
	// Add all keys from the keyring to the genesis accounts as well.
	//
	// NOTE: This is necessary to enable the account to send EVM transactions,
	// because the Mono ante handler checks the account balance by querying the
	// account from the account keeper first. If these accounts are not in the genesis
	// state, the ante handler finds a zero balance because of the missing account.

	// if denom not provided, defaults to create only one dynamic erc20
	// precompile with the 'xmpl' denom
	if len(denoms) == 0 {
		denoms = []string{"xmpl"}
	}
	accs := keyring.GetAllAccAddrs()
	genesisAccounts := make([]*authtypes.BaseAccount, len(accs))
	for i, addr := range accs {
		genesisAccounts[i] = &authtypes.BaseAccount{
			Address:       addr.String(),
			PubKey:        nil,
			AccountNumber: uint64(i + 1), //nolint:gosec // G115
			Sequence:      1,
		}
	}

	accGenesisState := authtypes.DefaultGenesisState()
	for _, genesisAccount := range genesisAccounts {
		// NOTE: This type requires to be packed into a *types.Any as seen on SDK tests,
		// e.g. https://github.com/lyfebloc/cosmos-sdk/blob/v0.47.5-lyfebloc.2/x/auth/keeper/keeper_test.go#L193-L223
		accGenesisState.Accounts = append(accGenesisState.Accounts, codectypes.UnsafePackAny(genesisAccount))
	}

	// get the common.Address to store the hex string addresses using EIP-55
	wlyfeblocAddr := common.HexToAddress(erc20types.WLYFEBLOCContractTestnet).Hex()

	// Add token pairs to genesis
	tokenPairs := make([]erc20types.TokenPair, 0, len(denoms)+1)
	// add token pair for fees token (wlyfebloc)
	tokenPairs = append(tokenPairs, erc20types.TokenPair{
		Erc20Address:  wlyfeblocAddr,
		Denom:         utils.BaseDenom,
		Enabled:       true,
		ContractOwner: erc20types.OWNER_MODULE, // NOTE: Owner is the module account since it's a native token and was registered through governance
	})

	dynPrecAddr := make([]string, 0, len(denoms))
	for _, denom := range denoms {
		addr := utiltx.GenerateAddress().Hex()
		tp := erc20types.TokenPair{
			Erc20Address:  addr,
			Denom:         denom,
			Enabled:       true,
			ContractOwner: erc20types.OWNER_MODULE, // NOTE: Owner is the module account since it's a native token and was registered through governance
		}
		tokenPairs = append(tokenPairs, tp)
		dynPrecAddr = append(dynPrecAddr, addr)
	}

	erc20GenesisState := erc20types.DefaultGenesisState()
	erc20GenesisState.TokenPairs = tokenPairs

	// STR v2: update the NativePrecompiles and DynamicPrecompiles
	// with the WLYFEBLOC (default is mainnet) and 'xmpl' tokens in the erc20 params
	erc20GenesisState.Params.NativePrecompiles = []string{wlyfeblocAddr}
	erc20GenesisState.Params.DynamicPrecompiles = dynPrecAddr

	// Add the smart contracts to the EVM genesis
	evmGenesisState := evmtypes.DefaultGenesisState()

	// Combine module genesis states
	return network.CustomGenesisState{
		authtypes.ModuleName:  accGenesisState,
		erc20types.ModuleName: erc20GenesisState,
		evmtypes.ModuleName:   evmGenesisState,
	}
}
