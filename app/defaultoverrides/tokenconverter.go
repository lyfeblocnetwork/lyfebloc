package defaultoverrides

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	tokenconvertermodule "github.com/lyfeblocnetwork/lyfebloc/x/tokenconverter/module"
	tokenconvertertypes "github.com/lyfeblocnetwork/lyfebloc/x/tokenconverter/types"
)

// TokenConverterModuleBasic defines a custom wrapper around the x/tokenconverter module's AppModuleBasic
// implementation to provide custom default genesis state.
type TokenConverterModuleBasic struct {
	tokenconvertermodule.AppModuleBasic
}

// DefaultGenesis returns custom x/tokenconverter module genesis state.
func (m TokenConverterModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := tokenconvertertypes.DefaultGenesis()
	genState.Params.BondDenom = "uvbloc"
	genState.Params.FeeDenom = "ubloc"

	return cdc.MustMarshalJSON(genState)
}
