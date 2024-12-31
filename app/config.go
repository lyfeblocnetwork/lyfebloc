// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)

package app

import (
	"github.com/lyfeblocnetwork/lyfebloc/app/eips"
	evmconfig "github.com/lyfeblocnetwork/lyfebloc/x/evm/config"
	"github.com/lyfeblocnetwork/lyfebloc/x/evm/core/vm"
)

// The init function of the config file allows to setup the global
// configuration for the EVM, modifying the custom ones defined in Lyfebloc.
func init() {
	err := evmconfig.NewEVMConfigurator().
		WithExtendedEips(lyfeblocActivators).
		Configure()
	if err != nil {
		panic(err)
	}
}

// LyfeblocActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var lyfeblocActivators = map[string]func(*vm.JumpTable){
	"lyfebloc_0": eips.Enable0000,
	"lyfebloc_1": eips.Enable0001,
	"lyfebloc_2": eips.Enable0002,
}
