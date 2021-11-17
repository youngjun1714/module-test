package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BettingInfoParams []*BettingInfo

func NewBetting(from string, betting string, amount sdk.Coins) Betting {

	return Betting{
		FromAddress: from,
		Betting:     betting,
		Amount:      amount,
	}
}

func NewTotalBettingInfo(info BettingInfoParams) TotalBettingInfo {
	return TotalBettingInfo{
		Info: info,
	}
}

func NewBettingInfo(amount int64, from string) *BettingInfo {
	return &BettingInfo{
		Amount:      amount,
		FromAddress: from,
	}
}

func NewResInfo(total int64) TotalResInfo {
	return TotalResInfo{
		Total: total,
	}
}

func GetBettingInfoParams(totalInfo TotalBettingInfo, amount int64, from string) TotalBettingInfo {

	info := append(totalInfo.Info, NewBettingInfo(amount, from))

	return NewTotalBettingInfo(info)
}
