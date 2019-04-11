package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateOrder{}, "hashgard/MsgCreateOrder", nil)
	cdc.RegisterConcrete(MsgWithdrawalOrder{}, "hashgard/MsgWithdrawalOrder", nil)
	cdc.RegisterConcrete(MsgTakeOrder{}, "hashgard/MsgTakeOrder", nil)
}

// generic sealed codec to be used throughout sdk
var MsgCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	MsgCdc = cdc.Seal()
}
