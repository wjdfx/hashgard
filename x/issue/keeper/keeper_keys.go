package keeper

import (
	"fmt"
	"strings"

	"github.com/hashgard/hashgard/x/issue/types"
)

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter   = []byte(":")
	KeyNextIssueID = []byte("newIssueID")
)

//func BytesString(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}
// Key for getting a specific issuer from the store
func KeyIssuer(issueIdStr string) []byte {
	return []byte(fmt.Sprintf("issues:%s", issueIdStr))
}

// Key for getting a specific address from the store
func KeyAddressIssues(addr string) []byte {
	return []byte(fmt.Sprintf("address:%s", addr))
}

// Key for getting a specific symbol from the store
func KeySymbolIssues(symbol string) []byte {
	return []byte(fmt.Sprintf("symbol:%s", strings.ToUpper(symbol)))
}

func KeyIssueIdStr(timestamp int64, seq int) string {
	return fmt.Sprintf("%s%d%02d", types.IDPreStr, timestamp, seq)
}
