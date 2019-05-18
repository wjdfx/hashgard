package keeper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter      = []byte(types.KeyDelimiterString)
	PrefixActiveQueue = []byte("active")
)

func KeyBoxIdStr(boxType string, seq uint64) string {
	return fmt.Sprintf("%s%s%s", types.IDPreStr, types.GetMustBoxTypeValue(boxType), strconv.FormatUint(seq, 36))
}

// Key for getting a specific issuer from the store
func KeyNextBoxID(boxType string) []byte {
	return []byte(fmt.Sprintf("newBoxID:%s", boxType))
}
func KeyBox(boxIdStr string) []byte {
	return []byte(fmt.Sprintf("ids:%s:%s", utils.GetBoxTypeByValue(boxIdStr), boxIdStr))
}

// Key for getting a specific address from the store
func KeyAddress(boxType string, accAddress sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("address:%s:%s", boxType, accAddress.String()))
}
func KeyName(boxType string, name string) []byte {
	return []byte(fmt.Sprintf("name:%s:%s", boxType, strings.ToLower(name)))
}
func KeyAddressDeposit(boxID string, accAddress sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("deposit:%s:%s", boxID, accAddress.String()))
}

func GetAddressFromKeyAddressDeposit(keyAddressDeposit []byte) sdk.AccAddress {
	str := fmt.Sprintf("%s", keyAddressDeposit)
	keys := strings.Split(str, ":")
	address, _ := sdk.AccAddressFromBech32(keys[2])
	return address
}
func PrefixKeyDeposit(boxID string) []byte {
	return []byte(fmt.Sprintf("deposit:%s", boxID))
}

// Returns the key for a boxID in the activeQueue
func PrefixActiveBoxQueueTime(endTime int64) []byte {
	return []byte(fmt.Sprintf("active:%d", endTime))
}

// Returns the key for a proposalID in the activeQueue
func KeyActiveBoxQueue(endTime int64, boxIdStr string) []byte {
	return []byte(fmt.Sprintf("active:%d:%s", endTime, boxIdStr))
}
