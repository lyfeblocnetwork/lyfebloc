package keeper

import (
	"github.com/lyfeblocnetwork/lyfebloc/x/fee/types"
)

var _ types.QueryServer = Keeper{}
