package utils

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

func CheckFreeze(cdc *codec.Codec, cliCtx context.CLIContext, issueID string, from sdk.AccAddress, to sdk.AccAddress) error {

	res, err := issuequeriers.QueryIssueFreeze(issueID, from, cliCtx)
	if err != nil {
		return err
	}
	var nowTime = time.Now()

	var freeze types.Freeze
	cdc.MustUnmarshalJSON(res, &freeze)

	if freeze.OutEndTime > 0 && time.Unix(freeze.OutEndTime, 0).After(nowTime) {
		return errors.Errorf(errors.ErrCanNotTransferOut(issueID, from.String()))
	}

	res, err = issuequeriers.QueryIssueFreeze(issueID, to, cliCtx)
	if err != nil {
		return err
	}

	cdc.MustUnmarshalJSON(res, &freeze)

	if freeze.InEndTime > 0 && time.Unix(freeze.InEndTime, 0).After(nowTime) {
		return errors.Errorf(errors.ErrCanNotTransferIn(issueID, to.String()))
	}

	return nil
}
