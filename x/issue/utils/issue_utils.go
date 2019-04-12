package utils

import (
	"math"
	"math/rand"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
