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
	randomBytes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
)

func getRandomString(l int) string {
	result := make([]rune, l)
	length := len(randomBytes)
	for i := range result {
		result[i] = randomBytes[rand.Intn(length)]
	}
	return string(result)
}

func GetIssueID() string {
	randLength := types.IDLength - len(types.IDPreStr)
	randString := getRandomString(randLength)
	return types.IDPreStr + randString
}

func IsIssueId(issueID string) bool {
	if len(issueID) == types.IDLength && strings.HasPrefix(issueID, types.IDPreStr) {
		return true
	}
	return false
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
func BurnCheck(cdc *codec.Codec, cliCtx context.CLIContext, operator auth.Account, burnFrom sdk.AccAddress, issueID string, amount sdk.Int) (sdk.Int, error) {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return amount, err
	}

	cdc.MustUnmarshalJSON(res, &issueInfo)

	if burnFrom == nil {
		if !operator.GetAddress().Equals(issueInfo.GetOwner()) {
			return amount, errors.Errorf(errors.ErrOwnerMismatch(issueID))
		}
		if issueInfo.GetBurnOff() {
			return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
		}
	}

	amount = MulDecimals(amount, issueInfo.GetDecimals())

	coins := operator.GetCoins()

	if operator.GetAddress().Equals(issueInfo.GetOwner()) {
		if operator.GetAddress().Equals(burnFrom) {
			if issueInfo.GetBurnOff() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
			}
		} else {
			if issueInfo.GetBurnAnyOff() {
				return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
			}
		}
	} else {
		if !operator.GetAddress().Equals(burnFrom) {
			return amount, errors.Errorf(errors.ErrOwnerMismatch(issueID))
		}
		if issueInfo.GetBurnFromOff() {
			return amount, errors.Errorf(errors.ErrCanNotBurn(issueID))
		}
		burnAccount, err := cliCtx.GetAccount(burnFrom)
		if err != nil {
			return amount, err
		}
		coins = burnAccount.GetCoins()
	}

	// ensure account has enough coins
	if !coins.IsAllGTE(sdk.Coins{sdk.Coin{Denom: issueID, Amount: amount}}) {
		return amount, fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", operator.GetAddress())
	}

	return amount, nil
}
func IssueOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, operator auth.Account, issueID string) error {
	var issueInfo types.Issue
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return err
	}
	cdc.MustUnmarshalJSON(res, &issueInfo)

	if !operator.GetAddress().Equals(issueInfo.GetOwner()) {
		return errors.Errorf(errors.ErrOwnerMismatch(issueID))
	}
	return nil
}
