// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global erc20 module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to modules/erc20 and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	AminoCdc = codec.NewAminoCodec(amino) //nolint:staticcheck
)

const (
	// Amino names
	convertERC20Name = "lyfebloc/MsgConvertERC20"
	convertCoinName  = "lyfebloc/MsgConvertCoin" // keep it for backwards compatibility when querying txs
	updateParams     = "lyfebloc/erc20/MsgUpdateParams"
	registerERC20    = "lyfebloc/erc20/MsgRegisterERC20"
	toggleConversion = "lyfebloc/erc20/MsgToggleConversion"
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConvertCoin{}, // keep it for backwards compatibility when querying txs
		&MsgConvertERC20{},
		&MsgUpdateParams{},
		&MsgRegisterERC20{},
		&MsgToggleConversion{},
	)
	registry.RegisterImplementations(
		(*govv1beta1.Content)(nil),
		&RegisterCoinProposal{}, // Keep interface for backwards compatibility on proposals query
		&RegisterERC20Proposal{},
		&ToggleTokenConversionProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the necessary x/erc20 interfaces and
// concrete types on the provided LegacyAmino codec. These types are used for
// Amino JSON serialization and EIP-712 compatibility.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, updateParams, nil)
	cdc.RegisterConcrete(&MsgConvertERC20{}, convertERC20Name, nil)
	cdc.RegisterConcrete(&MsgConvertCoin{}, convertCoinName, nil)
	cdc.RegisterConcrete(&MsgRegisterERC20{}, registerERC20, nil)
	cdc.RegisterConcrete(&MsgToggleConversion{}, toggleConversion, nil)
}