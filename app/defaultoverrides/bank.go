package defaultoverrides

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankModuleBasic defines a custom wrapper around the x/bank module's AppModuleBasic
// implementation to provide custom default genesis state.
type BankModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns custom x/bank module genesis state.
func (BankModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	metadataFee := banktypes.Metadata{
		Description: "The native token of the Lyfebloc network for fees.",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "ubloc",
				Exponent: 0,
				Aliases:  []string{"microbloc"},
			},
			{
				Denom:    "bloc",
				Exponent: 6,
			},
		},
		Base:    "ubloc",
		Display: "bloc",
		Name:    "Lyfebloc BLOC",
		Symbol:  "BLOC",
	}
	metadataBond := banktypes.Metadata{
		Description: "The native token of the Lyfebloc network for staking. This token is non transferrable. This token can be retrieved by providing liquidity.",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uvbloc",
				Exponent: 0,
				Aliases:  []string{"microvbloc"},
			},
			{
				Denom:    "vbloc",
				Exponent: 6,
			},
		},
		Base:    "uvbloc",
		Display: "vbloc",
		Name:    "Lyfebloc vBLOC",
		Symbol:  "vBLOC",
	}

	sendEnabledVbloc := banktypes.SendEnabled{
		Denom:   "uvbloc",
		Enabled: false,
	}

	genState := banktypes.DefaultGenesisState()
	genState.DenomMetadata = append(genState.DenomMetadata, metadataFee)
	genState.DenomMetadata = append(genState.DenomMetadata, metadataBond)

	genState.SendEnabled = append(genState.SendEnabled, sendEnabledVbloc)

	return cdc.MustMarshalJSON(genState)
}
