package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"strings"

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
