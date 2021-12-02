package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjun1714/module-test/x/game/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, lefacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)
		fmt.Println("========= NewQuerier =========")
		switch path[0] {
		case types.QueryBettingAmount:
			res, err = queryTotalBetting(ctx, k, lefacyQuerierCdc)
		}

		return res, err
	}
}

func queryTotalBetting(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	//	amt := k.GetEvensAmount(ctx)
	//	amt := k.GetOddsAmount(ctx)
	var total sdk.Coins
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, types.NewResInfo(total))

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
