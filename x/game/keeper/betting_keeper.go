package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjun1714/module-test/x/game/types"
)

func (k Keeper) Betting(ctx sdk.Context, msg types.Betting) error {

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		from,
		types.ModuleName,
		msg.Amount,
	)
	if err != nil {
		return err
	}

	switch msg.Betting {
	case "Odds":
		info := types.GetBettingInfoParams(k.GetOddsAmount(ctx), msg.Amount, msg.FromAddress)
		k.SetOddsAmount(ctx, info)
	case "Evens":
		info := types.GetBettingInfoParams(k.GetEvensAmount(ctx), msg.Amount, msg.FromAddress)
		k.SetEvensAmount(ctx, info)
	default:
		return fmt.Errorf("msg type error")
	}

	return nil
}

func (k Keeper) Rewards(ctx sdk.Context, msg types.Rewards) error {

	oddinfo := k.GetOddsAmount(ctx)
	evensinfo := k.GetEvensAmount(ctx)

	var (
		totalAmount, oddsAmount, evnesAmount sdk.Dec
	)

	for _, oddbettinginfo := range oddinfo.Info {
		totalAmount.Add(sdk.NewDecFromInt(oddbettinginfo.Amount[0].Amount))
		oddsAmount.Add(sdk.NewDecFromInt(oddbettinginfo.Amount[0].Amount))
	}

	for _, evensbettinginfo := range evensinfo.Info {
		totalAmount.Add(sdk.NewDecFromInt(evensbettinginfo.Amount[0].Amount))
		evnesAmount.Add(sdk.NewDecFromInt(evensbettinginfo.Amount[0].Amount))
	}
	/**/

	/**/
	switch msg.Winners {
	case "Odds":
		k.DelStore(ctx, types.KeyEvensAmount)

	case "Evens":
		k.DelStore(ctx, types.KeyOddsAmount)

	default:

	}

	return nil
}

func (k Keeper) SetOddsAmount(ctx sdk.Context, info types.TotalBettingInfo) {

	store := k.Store(ctx)
	bz := k.cdc.MustMarshalBinaryBare(&info)
	store.Set([]byte(types.KeyOddsAmount), bz)
}
func (k Keeper) GetOddsAmount(ctx sdk.Context) types.TotalBettingInfo {

	store := k.Store(ctx)
	var info types.TotalBettingInfo
	bz := store.Get(types.KeyOddsAmount)
	k.cdc.UnmarshalBinaryBare(bz, &info)
	return info
}
func (k Keeper) SetEvensAmount(ctx sdk.Context, info types.TotalBettingInfo) {
	store := k.Store(ctx)
	bz := k.cdc.MustMarshalBinaryBare(&info)
	store.Set([]byte(types.KeyEvensAmount), bz)
}

func (k Keeper) GetEvensAmount(ctx sdk.Context) types.TotalBettingInfo {

	store := k.Store(ctx)
	var info types.TotalBettingInfo
	bz := store.Get(types.KeyEvensAmount)
	k.cdc.UnmarshalBinaryBare(bz, &info)
	return info
}

func (k Keeper) DelStore(ctx sdk.Context, keyName []byte) {

	store := k.Store(ctx)
	store.Delete(keyName)
}
