package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjun1714/module-test/x/game/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) BetAmount(goCtx context.Context, msg *types.MsgBetAmountRequest) (*types.MsgBetAmountResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	newBetting := types.NewBetting(msg.FromAddress, msg.Betting, msg.Amount)

	err := k.Keeper.Betting(ctx, newBetting)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBetting,
			sdk.NewAttribute(types.AttributeKeySender, msg.FromAddress),
			sdk.NewAttribute(types.AtttibuteKeyBetting, msg.Betting),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.FromAddress),
		),
	})

	return &types.MsgBetAmountResponse{}, err
}

func (k msgServer) AllRewards(goCtx context.Context, msg *types.MsgAllRewardsRequest) (*types.MsgAllRewardsResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	rewards := types.NewRewards(msg.ToAddress, msg.Winners, msg.Amount)

	err := k.Keeper.Rewards(ctx, rewards)

	if err != nil {
		return nil, err
	}

	return &types.MsgAllRewardsResponse{}, nil
}
