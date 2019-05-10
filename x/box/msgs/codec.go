package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hashgard/hashgard/x/box/types"
)

var MsgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBox{}, "box/MsgBox", nil)
	cdc.RegisterConcrete(MsgBoxInterest{}, "box/MsgBoxInterest", nil)
	cdc.RegisterConcrete(MsgBoxDeposit{}, "box/MsgBoxDeposit", nil)
	cdc.RegisterConcrete(MsgBoxDescription{}, "box/MsgBoxDescription", nil)
	cdc.RegisterConcrete(MsgBoxDisableFeature{}, "box/MsgBoxDisableFeature", nil)

	cdc.RegisterInterface((*types.Box)(nil), nil)
	cdc.RegisterConcrete(&types.BoxInfo{}, "box/BoxInfo", nil)
}

//nolint
func init() {
	RegisterCodec(MsgCdc)
}
