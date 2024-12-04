package keeper

import (
	"github.com/lyfeblocnetwork/lyfebloc/x/blobstream/types"
)

var _ types.QueryServer = Keeper{}
