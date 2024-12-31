// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoLyfebloc defines the default coin denomination used in Lyfebloc in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Lyfebloc.
	AttoLyfebloc string = "alyfe"

	// BaseDenomUnit defines the base denomination unit for Lyfebloc.
	// 1 lyfebloc = 1x10^{BaseDenomUnit} alyfe
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewLyfeblocCoin is a utility function that returns an "alyfe" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewLyfeblocCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoLyfebloc, amount)
}

// NewLyfeblocDecCoin is a utility function that returns an "alyfe" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewLyfeblocDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoLyfebloc, amount)
}

// NewLyfeblocCoinInt64 is a utility function that returns an "alyfe" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewLyfeblocCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoLyfebloc, amount)
}
