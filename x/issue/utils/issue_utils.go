package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/types"
)

var (
	randomBytes = []rune("abcdefghijklmnopqrstuvwxyz")
)

func GetRandomString(l int) string {
	result := make([]rune, l)
	length := len(randomBytes)
	for i := range result {
		result[i] = randomBytes[rand.Intn(length)]
	}
	return string(result)
}

//nolint
//func GetIssueID() string {
//
//	randLength := types.IDLength - len(types.IDPreStr)
//	randString := GetRandomString(randLength)
//	return types.IDPreStr + randString
//}

func IsIssueId(issueID string) bool {
	if strings.HasPrefix(issueID, types.IDPreStr) {
		return true
	}

	return true
}

func CheckIssueId(issueID string) sdk.Error {
	if !IsIssueId(issueID) {
		return errors.ErrIssueID(issueID)
	}
	return nil
}
func MulDecimals(totalSupply sdk.Int, decimals uint) sdk.Int {

	multiple := math.Pow10(int(decimals))
	multipleStr := strconv.FormatFloat(multiple, 'f', 0, 64)
	multipleDecimals, _ := sdk.NewIntFromString(multipleStr)

	return totalSupply.Mul(multipleDecimals)
}
func QuoDecimals(totalSupply sdk.Int, decimals uint) sdk.Int {

	quo := math.Pow10(int(decimals))
	quoStr := strconv.FormatFloat(quo, 'f', 0, 64)
	quoDecimals, _ := sdk.NewIntFromString(quoStr)

	return totalSupply.Quo(quoDecimals)
}
func BurnCheck(cdc *codec.Codec, cliCtx context.CLIContext, operator auth.Account, burnFrom sdk.AccAddress, issueID string, amount sdk.Int, burnType string) (sdk.Int, error) {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return amount, err
	}

	cdc.MustUnmarshalJSON(res, &issueInfo)

	amount = MulDecimals(amount, issueInfo.GetDecimals())

	coins := operator.GetCoins()

	switch burnType {
	case types.BurnOwner:
		{
			if !operator.GetAddress().Equals(issueInfo.GetOwner()) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueID))
			}
			if issueInfo.GetBurnOff() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
			}
		}
	case types.BurnFrom:
		{
			if issueInfo.GetBurnFromOff() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
			}
			if !operator.GetAddress().Equals(burnFrom) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueID))
			}
		}
	case types.BurnAny:
		{
			if !operator.GetAddress().Equals(issueInfo.GetOwner()) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueID))
			}
			if issueInfo.GetBurnAnyOff() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
			}
			if operator.GetAddress().Equals(burnFrom) {
				//burnFrom
				if issueInfo.GetBurnFromOff() {
					return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
				}
			}
		}
	default:
		{
			panic("not support")
		}

	}
	// ensure account has enough coins
	if !coins.IsAllGTE(sdk.Coins{sdk.Coin{Denom: issueID, Amount: amount}}) {
		return amount, fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", operator.GetAddress())
	}

	return amount, nil
}
func IssueOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, operator auth.Account, issueID string) (types.Issue, error) {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return nil, err
	}
	cdc.MustUnmarshalJSON(res, &issueInfo)

	if !operator.GetAddress().Equals(issueInfo.GetOwner()) {
		return nil, errors.Errorf(errors.ErrOwnerMismatch(issueID))
	}
	return issueInfo, nil
}
