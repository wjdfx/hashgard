package utils

import (
	"math/rand"
	"strings"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/errors"
)

var (
	randomBytes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
)

func GetRandomString(l int) string {
	result := make([]rune, l)
	len := len(randomBytes)
	for i := range result {
		result[i] = randomBytes[rand.Intn(len)]
	}
	return string(result)
}

func GetIssueID() string {
	randString := GetRandomString(11)
	return domain.IDPreStr + randString
}
func IsIssueId(issueID string) bool {
	if len(issueID) == 15 && strings.HasPrefix(issueID, domain.IDPreStr) {
		return true
	}
	return false
}

func CheckIssueId(issueID string) error {
	if !IsIssueId(issueID) {
		return errors.ErrIssueID(domain.DefaultCodespace, issueID)
	}
	return nil
}
