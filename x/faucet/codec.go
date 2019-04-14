package faucet

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgFaucetSend{}, "hashgard/MsgFaucetSend", nil)
}

// generic sealed codec to be used throughout sdk
var msgCdc *codec.Codec

func init() {
	cdc := codec.New()
	RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	msgCdc = cdc.Seal()
}