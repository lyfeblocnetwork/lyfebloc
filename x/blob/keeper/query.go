package keeper

import (
	"github.com/lyfeblocnetwork/lyfebloc/x/blob/types"
)

var _ types.QueryServer = Keeper{}
