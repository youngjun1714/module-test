package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBetAmount = "bet-amount"

var _ sdk.Msg = &MsgBetAmountRequest{}

func NewMsgBetAmountRequest(fromAddr sdk.AccAddress, betting string, amount sdk.Coins) *MsgBetAmountRequest {
	return &MsgBetAmountRequest{FromAddress: fromAddr.String(), Betting: betting, Amount: amount}
}

func (msg MsgBetAmountRequest) Route() string { return RouterKey }

func (msg MsgBetAmountRequest) Type() string { return TypeMsgBetAmount }

func (msg MsgBetAmountRequest) ValidateBasic() error {

	_, err := sdk.AccAddressFromBech32(msg.FromAddress)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgBetAmountRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBetAmountRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
