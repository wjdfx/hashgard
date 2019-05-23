package utils

import (
	"fmt"
	"reflect"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/hashgard/hashgard/x/box/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
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
func GetCoinDecimal(cdc *codec.Codec, cliCtx context.CLIContext, coin sdk.Coin) (uint, error) {
	if coin.Denom == types.Agard {
		return types.AgardDecimals, nil
	}
	issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, coin.Denom)
	if err != nil {
		return 0, err
	}
	return issueInfo.GetDecimals(), nil
}
func GetBoxInfo(box types.BoxInfo) fmt.Stringer {
	switch box.BoxType {
	case types.Lock:
		var clientBox LockBoxInfo
		StructCopy(&clientBox, &box)
		return clientBox
	case types.Deposit:
		return processDepositBoxInfo(box)
	case types.Future:
		var clientBox FutureBoxInfo
		StructCopy(&clientBox, &box)
		return clientBox
	default:
		return box
	}
}
func processDepositBoxInfo(box types.BoxInfo) DepositBoxInfo {
	var clientBox DepositBoxInfo
	StructCopy(&clientBox, &box)
	return clientBox
}
func GetBoxList(boxs types.BoxInfos, boxType string) fmt.Stringer {
	switch boxType {
	case types.Lock:
		var boxInfos = make(LockBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			var clientBox LockBoxInfo
			StructCopy(&clientBox, &box)
			boxInfos = append(boxInfos, clientBox)
		}

		return boxInfos
	case types.Deposit:
		var boxInfos = make(DepositBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			boxInfos = append(boxInfos, processDepositBoxInfo(box))
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
