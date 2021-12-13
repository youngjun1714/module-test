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

	totalAmount := sdk.NewCoin("stake", sdk.NewInt(0))

	for _, oddbettinginfo := range oddinfo.Info {
		totalAmount = totalAmount.Add(oddbettinginfo.Amount[0])
	}

	for _, evensbettinginfo := range evensinfo.Info {
		totalAmount = totalAmount.Add(evensbettinginfo.Amount[0])
	}

	return &types.QueryTotalBettingResponse{ResInfo: types.NewResInfo(sdk.NewCoins(totalAmount))}, nil
}

func (k Querier) OddsBetting(ctx context.Context, _ *types.QueryOddsBettingRequest) (*types.QueryOddsBettingResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	oddsinfo := k.GetOddsAmount(sdkCtx)

	return &types.QueryOddsBettingResponse{OddsInfo: types.NewTotalBettingInfo(oddsinfo.Info)}, nil
}

func (k Querier) EvensBetting(ctx context.Context, _ *types.QueryEvensBettingRequest) (*types.QueryEvensBettingResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	evensinfo := k.GetOddsAmount(sdkCtx)

	return &types.QueryEvensBettingResponse{EvensInfo: types.NewTotalBettingInfo(evensinfo.Info)}, nil
}
