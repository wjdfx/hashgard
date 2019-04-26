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
	return strings.HasPrefix(issueID, types.IDPreStr)
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
func CheckAllowance(cdc *codec.Codec, cliCtx context.CLIContext, issueID string, owner sdk.AccAddress, spender sdk.AccAddress, amount sdk.Int) error {

	res, err := issuequeriers.QueryIssueAllowance(issueID, owner, spender, cliCtx)
	if err != nil {
		return err
	}
	var approval types.Approval
	cdc.MustUnmarshalJSON(res, &approval)

	if approval.Amount.LT(amount) {
		return errors.Errorf(errors.ErrNotEnoughAmountToTransfer())
	}

	return nil

}
func GetIssueByID(cdc *codec.Codec, cliCtx context.CLIContext, issueID string) (types.Issue, error) {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return nil, err
	}

	cdc.MustUnmarshalJSON(res, &issueInfo)

	return issueInfo, nil
}
func BurnCheck(sender auth.Account, burnFrom sdk.AccAddress, issueInfo types.Issue, amount sdk.Int, burnType string) (sdk.Int, error) {

	amount = MulDecimals(amount, issueInfo.GetDecimals())

	coins := sender.GetCoins()

	switch burnType {
	case types.BurnOwner:
		{
			if !sender.GetAddress().Equals(issueInfo.GetOwner()) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if !sender.GetAddress().Equals(burnFrom) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if issueInfo.IsBurnOwnerDisabled() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
		}
	case types.BurnHolder:
		{
			if issueInfo.IsBurnHolderDisabled() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
			if !sender.GetAddress().Equals(burnFrom) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
		}
	case types.BurnFrom:
		{
			if !sender.GetAddress().Equals(issueInfo.GetOwner()) {
				return amount, errors.Errorf(errors.ErrOwnerMismatch(issueInfo.GetIssueId()))
			}
			if issueInfo.IsBurnFromDisabled() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), burnType))
			}
			if issueInfo.GetOwner().Equals(burnFrom) {
				//burnFrom
				if issueInfo.IsBurnOwnerDisabled() {
					return amount, errors.Errorf(errors.ErrCanNotBurn(issueInfo.GetIssueId(), types.BurnOwner))
				}
			}
		}
	default:
		{
			panic("not support")
		}

	}
	// ensure account has enough coins
	if !coins.IsAllGTE(sdk.NewCoins(sdk.NewCoin(issueInfo.GetIssueId(), amount))) {
		return amount, fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", sender.GetAddress())
	}
	return amount, nil
}
func IssueOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, sender auth.Account, issueID string) (types.Issue, error) {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return nil, err
	}
	cdc.MustUnmarshalJSON(res, &issueInfo)

	if !sender.GetAddress().Equals(issueInfo.GetOwner()) {
		return nil, errors.Errorf(errors.ErrOwnerMismatch(issueID))
	}
	return issueInfo, nil
}
