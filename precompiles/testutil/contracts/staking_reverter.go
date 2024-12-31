// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)

package contracts

import (
	contractutils "github.com/lyfeblocnetwork/lyfebloc/contracts/utils"
	evmtypes "github.com/lyfeblocnetwork/lyfebloc/x/evm/types"
)

func LoadStakingReverterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("StakingReverter.json")
}
