// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/lyfeblocnetwork/lyfebloc/x/evm/types"
)

var (
	//go:embed compiled_contracts/WLYFEBLOC.json
	WLYFEBLOCJSON []byte

	// WLYFEBLOCContract is the compiled contract of WLYFEBLOC
	WLYFEBLOCContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WLYFEBLOCJSON, &WLYFEBLOCContract)
	if err != nil {
		panic(err)
	}

	if len(WLYFEBLOCContract.Bin) == 0 {
		panic("failed to load WLYFEBLOC smart contract")
	}
}
