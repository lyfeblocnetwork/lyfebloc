package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquidityincentive/types"
)

func (k msgServer) CollectVoteRewards(goCtx context.Context, msg *types.MsgCollectVoteRewards) (*types.MsgCollectVoteRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCollectVoteRewardsResponse{}, nil
}
