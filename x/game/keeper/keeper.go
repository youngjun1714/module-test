/*
	1. 전체 amount 저장.
	2. 누가 얼마큼 어디에 betting 했는지 저장.
*/

package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/youngjun1714/module-test/x/game/types"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	params     paramtypes.Subspace
	bankKeeper types.BankKeeper
}

func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	params paramtypes.Subspace,
	bankKeeper types.BankKeeper,
) Keeper {

	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		params:     params.WithKeyTable(types.ParamKeyTable()),
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
