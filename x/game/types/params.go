package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultBetting_enable = true
)

var (
	KeyBetting_enable = []byte("Battingenable")
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)
)

func NewParams(betting_enable bool) Params {
	return Params{
		Bettings: betting_enable,
	}
}

func DefaultParams() Params {
	return Params{
		Bettings: true,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBetting_enable, &p.Bettings, validateBettings),
	}
}

func validateBettings(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

/*func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}*/
