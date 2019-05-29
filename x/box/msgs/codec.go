package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hashgard/hashgard/x/box/types"
)

var MsgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgLockBox{}, "box/MsgLockBox", nil)
	cdc.RegisterConcrete(MsgDepositBox{}, "box/MsgDepositBox", nil)
	cdc.RegisterConcrete(MsgFutureBox{}, "box/MsgFutureBox", nil)
	cdc.RegisterConcrete(MsgBoxInterestInjection{}, "box/MsgBoxInterestInjection", nil)
	cdc.RegisterConcrete(MsgBoxInterestFetch{}, "box/MsgBoxInterestFetch", nil)
	cdc.RegisterConcrete(MsgBoxDepositTo{}, "box/MsgBoxDepositTo", nil)
	cdc.RegisterConcrete(MsgBoxDepositFetch{}, "box/MsgBoxDepositFetch", nil)
	cdc.RegisterConcrete(MsgBoxWithdraw{}, "box/MsgBoxWithdraw", nil)
	cdc.RegisterConcrete(MsgBoxDescription{}, "box/MsgBoxDescription", nil)
	cdc.RegisterConcrete(MsgBoxDisableFeature{}, "box/MsgBoxDisableFeature", nil)

	cdc.RegisterInterface((*types.Box)(nil), nil)
	cdc.RegisterConcrete(&types.BoxInfo{}, "box/BoxInfo", nil)
}

//nolint
func init() {
	RegisterCodec(MsgCdc)
}
