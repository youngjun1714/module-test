package game

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/youngjun1714/module-test/x/game/keeper"
	"github.com/youngjun1714/module-test/x/game/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {

	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgBetAmountRequest:
			res, err := msgServer.BetAmount(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}

}
