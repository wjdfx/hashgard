package faucet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = MsgFaucetSend{}

type MsgFaucetSend struct {
	Receiver sdk.AccAddress `json:"receiver"`
}

func NewMsgFaucetSend(receiver sdk.AccAddress) MsgFaucetSend {
	return MsgFaucetSend{
		Receiver: receiver,
	}
}

// implement Msg interface
func (msg MsgFaucetSend) Route() string {
	return RouterKey
}

func (msg MsgFaucetSend) Type() string {
	return "faucet"
}

func (msg MsgFaucetSend) ValidateBasic() sdk.Error {
	if msg.Receiver.Empty() {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "receiver address is nil")
	}

	return nil
}

func (msg MsgFaucetSend) GetSignBytes() []byte {
	bz := msgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgFaucetSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Receiver}
}
