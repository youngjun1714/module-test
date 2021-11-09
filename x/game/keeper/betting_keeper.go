package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/youngjun1714/module-test/x/game/types"
)

func (k Keeper) Betting(ctx sdk.Context, msg types.Betting) error {

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	fmt.Println("========== Betting ============")
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
		amt := k.GetOddsAmount(ctx)
		amt += msg.Amount[0].Amount.Int64()
		k.SetOddsAmount(ctx, types.NewBettingAmount(amt))
	case "Evens":
		amt := k.GetEvensAmount(ctx)
		amt += msg.Amount[0].Amount.Int64()
		k.SetEvensAmount(ctx, types.NewBettingAmount(amt))
	default:
		fmt.Println("================ default ===============")
		return fmt.Errorf("msg type error")
	}

	return nil
}

func (k Keeper) SetOddsAmount(ctx sdk.Context, amt types.BettingAmount) {

	store := k.Store(ctx)
	bz := k.cdc.MustMarshalBinaryBare(&amt)
	store.Set([]byte(types.KeyOddsAmount), bz)

}
func (k Keeper) GetOddsAmount(ctx sdk.Context) int64 {

	store := k.Store(ctx)

	var amt types.BettingAmount

	bz := store.Get(types.KeyOddsAmount)

	k.cdc.UnmarshalBinaryBare(bz, &amt)

	return amt.Amount
}
func (k Keeper) SetEvensAmount(ctx sdk.Context, amt types.BettingAmount) {

	store := k.Store(ctx)
	bz := k.cdc.MustMarshalBinaryBare(&amt)
	store.Set([]byte(types.KeyEvensAmount), bz)
}

func (k Keeper) GetEvensAmount(ctx sdk.Context) int64 {

	store := k.Store(ctx)

	var amt types.BettingAmount

	bz := store.Get(types.KeyEvensAmount)

	k.cdc.UnmarshalBinaryBare(bz, &amt)

	return amt.Amount
}
