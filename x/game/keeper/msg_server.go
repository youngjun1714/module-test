package keeper

import (
	"context"

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

	//	ctx := sdk.UnwrapSDKContext(goCtx)

	//	if err :=

	return &types.MsgBetAmountResponse{}, nil
}
