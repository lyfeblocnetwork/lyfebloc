package keeper

import (
	"github.com/lyfeblocnetwork/lyfebloc/x/liquiditypool/types"
)

var _ types.QueryServer = Keeper{}
