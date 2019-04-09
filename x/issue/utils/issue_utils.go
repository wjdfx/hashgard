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

//nolint
func getRandomString(l int) string {
	result := make([]rune, l)
	len := len(randomBytes)
	for i := range result {
		result[i] = randomBytes[rand.Intn(len)]
	}
	return string(result)
}

//nolint
func GetIssueID() string {
	randString := getRandomString(11)
	return types.IDPreStr + randString
}

//nolint
func IsIssueId(issueID string) bool {
	if len(issueID) == 15 && strings.HasPrefix(issueID, types.IDPreStr) {
		return true
	}
	return false
}

//nolint
func CheckIssueId(issueID string) sdk.Error {
	if !IsIssueId(issueID) {
		return errors.ErrIssueID(types.DefaultCodespace, issueID)
	}
	return nil
}
