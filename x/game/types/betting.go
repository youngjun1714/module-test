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

func NewRewards(to string, winners string, amount int64) Rewards {
	return Rewards{
		ToAddress: to,
		Winners:   winners,
		Amount:    amount,
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

/*
func NewOddsInfo(info TotalBettingInfo) TotalBettingInfo {
	return TotalBettingInfo{
		Info: info,
	}
}
*/
func GetBettingInfoParams(totalInfo TotalBettingInfo, amount int64, from string) TotalBettingInfo {

	info := append(totalInfo.Info, NewBettingInfo(amount, from))
	//	info = info[:len(info)-1]
	return NewTotalBettingInfo(info)
}

func GetRewards(totalInfo TotalBettingInfo, to string) TotalBettingInfo {

	var index int
	for i, info := range totalInfo.Info {

		if info.FromAddress == to {
			index = i
			break
		}
	}

	temp := make(BettingInfoParams, 0)
	temp2 := append(temp, totalInfo.Info[:index]...)
	temp3 := append(temp2, totalInfo.Info[index+1:]...)
	return NewTotalBettingInfo(temp3)

}
