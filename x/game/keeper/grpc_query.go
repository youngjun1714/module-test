package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjun1714/module-test/x/game/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) TotalBetting(ctx context.Context, _ *types.QueryTotalBettingRequest) (*types.QueryTotalBettingResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	oddinfo := k.GetOddsAmount(sdkCtx)
	evensinfo := k.GetEvensAmount(sdkCtx)

	var totalAmount int64 = 0

	for _, oddbettinginfo := range oddinfo.Info {
		totalAmount += oddbettinginfo.Amount
	}

	for _, evensbettinginfo := range evensinfo.Info {
		totalAmount += evensbettinginfo.Amount
	}

	return &types.QueryTotalBettingResponse{ResInfo: types.NewResInfo(totalAmount)}, nil
}
