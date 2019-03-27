package issue

import (
	"math/rand"
)

const (
	IDPreStr = "gard"
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
	return IDPreStr + randString
}
