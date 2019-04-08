package exchange

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/exchange/msgs"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(msgs.MsgCreateOrder{}, "hashgard/MsgCreateOrder", nil)
	cdc.RegisterConcrete(msgs.MsgWithdrawalOrder{}, "hashgard/MsgWithdrawalOrder", nil)
	cdc.RegisterConcrete(msgs.MsgTakeOrder{}, "hashgard/MsgTakeOrder", nil)
}

// generic sealed codec to be used throughout sdk
var MsgCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	MsgCdc = cdc.Seal()
}