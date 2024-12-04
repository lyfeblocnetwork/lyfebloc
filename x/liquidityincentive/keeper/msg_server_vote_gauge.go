package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lyfeblocnetwork/lyfebloc/x/liquidityincentive/types"
	liquiditypooltypes "github.com/lyfeblocnetwork/lyfebloc/x/liquiditypool/types"
)

func (k msgServer) VoteGauge(goCtx context.Context, msg *types.MsgVoteGauge) (*types.MsgVoteGaugeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, weight := range msg.Weights {
		if _, found := k.liquidityPoolKeeper.GetPool(goCtx, weight.PoolId); !found {
			return nil, liquiditypooltypes.ErrPoolNotFound
		}
	}

	k.SetVote(ctx, types.Vote{
		Sender:  msg.Sender,
		Weights: msg.Weights,
	})

	return &types.MsgVoteGaugeResponse{}, nil
}
