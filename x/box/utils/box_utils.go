package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	boxqueriers "github.com/hashgard/hashgard/x/box/client/queriers"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"

	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

func MulDecimals(coin sdk.Coin, decimals uint) sdk.Int {
	if coin.Denom == types.Agard {
		return coin.Amount
	}
	return issueutils.MulDecimals(coin.Amount, decimals)
}
func QuoDecimals(coin sdk.Coin, decimals uint) sdk.Int {
	if coin.Denom == types.Agard {
		return coin.Amount
	}
	return issueutils.QuoDecimals(coin.Amount, decimals)
}
func ParseCoin(denom string, amount sdk.Int) sdk.Coin {
	if denom == types.Agard {
		denom = types.Gard
	}
	coin, _ := sdk.ParseCoin(fmt.Sprintf("%s%s", amount, denom))
	return coin
}
func CalcInterest(perCoupon sdk.Dec, share sdk.Int, interest types.BoxToken) sdk.Int {
	dec := perCoupon.MulInt(share)
	decimals := interest.Decimals
	if interest.Token.Denom == types.Agard {
		decimals = types.GardDecimal
	}
	dec = GetMaxPrecision(dec, decimals)
	return dec.MulInt(issueutils.GetDecimalsInt(decimals)).TruncateInt()
}

func IsBoxId(boxID string) bool {
	return strings.HasPrefix(boxID, types.IDPreStr)
}

func CheckBoxId(boxID string) sdk.Error {
	if !IsBoxId(boxID) {
		return errors.ErrBoxID(boxID)
	}
	return nil
}

func CalcInterestRate(totalAmount sdk.Int, price sdk.Int, interest sdk.Coin, decimals uint) sdk.Dec {
	totalCoupon := totalAmount.Quo(price)
	perCoupon := sdk.NewDecFromBigInt(interest.Amount.BigInt()).QuoInt(totalCoupon)
	if interest.Denom == types.Agard {
		decimals = types.GardDecimal
	}
	return quoMaxPrecisionByDecimal(perCoupon, decimals)
}

func quoMaxPrecisionByDecimal(dec sdk.Dec, decimals uint) sdk.Dec {
	dec = dec.QuoInt(issueutils.GetDecimalsInt(decimals))
	dec = GetMaxPrecision(dec, decimals)
	return dec
}

func GetBoxByID(cdc *codec.Codec, cliCtx context.CLIContext, boxID string) (types.Box, error) {
	var boxInfo types.Box
	// Query the box
	res, err := boxqueriers.QueryBoxByID(boxID, cliCtx)
	if err != nil {
		return nil, err
	}
	cdc.MustUnmarshalJSON(res, &boxInfo)
	return boxInfo, nil
}

func BoxOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, sender auth.Account, boxID string) (types.Box, error) {
	boxInfo, err := GetBoxByID(cdc, cliCtx, boxID)
	if err != nil {
		return nil, err
	}
	if !sender.GetAddress().Equals(boxInfo.GetOwner()) {
		return nil, errors.Errorf(errors.ErrOwnerMismatch(boxID))
	}
	return boxInfo, nil
}

//func GetBoxCoinByDecimal(cdc *codec.Codec, cliCtx context.CLIContext, coin sdk.Coin) sdk.Coin {
//
//	issueInfo, _ := issueutils.GetIssueByID(cdc, cliCtx, coin.Denom)
//
//	return sdk.Coin{fmt.Sprintf("%s(%s)", issueInfo.GetName(), coin.Denom), issueutils.QuoDecimals(coin.Amount, issueInfo.GetDecimals())}
//}
func GetBoxTypeByValue(value string) string {
	value = strings.ReplaceAll(value, types.IDPreStr, "")
	for k, v := range types.BoxType {
		if strings.HasPrefix(value, v) {
			return k
		}
	}
	return ""
}
func GetCoinDenomByFutureBoxSeq(boxID string, seq int) string {
	return fmt.Sprintf("%s%02d", boxID, seq)
}
func GetBoxIdFromBoxSeqID(boxIDSeq string) string {
	if len(boxIDSeq) > types.BoxIdLength {
		return boxIDSeq[:types.BoxIdLength]
	}
	return boxIDSeq
}
func GetSeqFromFutureBoxSeq(boxSeqStr string) int {
	seqStr := boxSeqStr[len(boxSeqStr)-2:]
	seq, _ := strconv.Atoi(seqStr)
	return seq
}
func GetMaxPrecision(dec sdk.Dec, decimals uint) sdk.Dec {
	precision := types.MaxPrecision
	if decimals < types.MaxPrecision {
		precision = decimals
	}
	decStr := dec.String()
	len := strings.Index(decStr, ".") + int(precision)
	str := decStr[0 : len+1]
	dec, _ = sdk.NewDecFromStr(str)
	return dec
}
