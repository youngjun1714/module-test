package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewBetting(from string, betting string, amount sdk.Coins) Betting {
	return Betting{
		FromAddress: from,
		Betting:     betting,
		Amount:      amount,
	}
}

func NewBettingAmount(amount int64) BettingAmount {
	return BettingAmount{
		Amount: amount,
	}
}
