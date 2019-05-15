package utils

import (
	"fmt"
	"reflect"

	"github.com/hashgard/hashgard/x/box/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCliContext(cdc *codec.Codec) (authtxb.TxBuilder, context.CLIContext, auth.Account, error) {
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
	cliCtx := context.NewCLIContext().
		WithCodec(cdc).
		WithAccountDecoder(cdc)
	from := cliCtx.GetFromAddress()
	account, err := cliCtx.GetAccount(from)

	return txBldr, cliCtx, account, err
}
func GetBoxInfo(cdc *codec.Codec, cliCtx context.CLIContext, box types.BoxInfo) fmt.Stringer {
	switch box.BoxType {
	case types.Lock:
		var clientBox LockBoxInfo
		StructCopy(&clientBox, &box)
		//clientBox.TotalAmount.Token = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, box.TotalAmount.Token)
		return clientBox
	case types.Deposit:
		return processDepositBoxInfo(cdc, cliCtx, box)
	case types.Future:
		var clientBox FutureBoxInfo
		StructCopy(&clientBox, &box)
		//clientBox.TotalAmount.Token = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, box.TotalAmount.Token)
		return clientBox
	default:
		return box
	}

}

func processDepositBoxInfo(cdc *codec.Codec, cliCtx context.CLIContext, box types.BoxInfo) DepositBoxInfo {
	var clientBox DepositBoxInfo
	StructCopy(&clientBox, &box)

	//clientBox.Deposit.Price = boxutils.GetBoxCoinByDecimal(cdc, cliCtx,
	//	sdk.NewCoin(clientBox.Deposit.Interest.Token.Denom, clientBox.Deposit.Price)).Amount
	//
	//clientBox.Deposit.BottomLine = boxutils.GetBoxCoinByDecimal(cdc, cliCtx,
	//	sdk.NewCoin(clientBox.Deposit.Interest.Token.Denom, clientBox.Deposit.BottomLine)).Amount
	//
	//clientBox.Deposit.TotalDeposit = boxutils.GetBoxCoinByDecimal(cdc, cliCtx,
	//	sdk.NewCoin(clientBox.TotalAmount.Token.Denom, clientBox.Deposit.TotalDeposit)).Amount
	//
	//if clientBox.Deposit.InterestInjections != nil {
	//	for i, v := range clientBox.Deposit.InterestInjections {
	//		clientBox.Deposit.InterestInjections[i].Amount = boxutils.GetBoxCoinByDecimal(cdc, cliCtx,
	//			sdk.NewCoin(clientBox.Deposit.Interest.Token.Denom, v.Amount)).Amount
	//	}
	//}
	//
	//clientBox.TotalAmount.Token = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, box.TotalAmount.Token)
	//clientBox.Deposit.Interest.Token = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, clientBox.Deposit.Interest.Token)

	return clientBox
}
func GetBoxList(cdc *codec.Codec, cliCtx context.CLIContext, boxs types.BoxInfos, boxType string) fmt.Stringer {
	switch boxType {
	case types.Lock:
		var boxInfos = make(LockBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			var clientBox LockBoxInfo
			StructCopy(&clientBox, &box)
			//boxs[i].TotalAmount.Token = boxutils.GetBoxCoinByDecimal(cdc, cliCtx, box.TotalAmount.Token)
			boxInfos = append(boxInfos, clientBox)
		}

		return boxInfos
	case types.Deposit:
		var boxInfos = make(DepositBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			boxInfos = append(boxInfos, processDepositBoxInfo(cdc, cliCtx, box))
		}

		return boxInfos
	case types.Future:
		var boxInfos = make(FutureBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			var clientBox FutureBoxInfo
			StructCopy(&clientBox, &box)
			boxInfos = append(boxInfos, clientBox)
		}

		return boxInfos
	}
	return boxs
}
func deepFields(faceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < faceType.NumField(); i++ {
		v := faceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}

func StructCopy(destPtr interface{}, srcPtr interface{}) {
	srcv := reflect.ValueOf(srcPtr)
	dstv := reflect.ValueOf(destPtr)
	srct := reflect.TypeOf(srcPtr)
	dstt := reflect.TypeOf(destPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := deepFields(reflect.ValueOf(srcPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}
