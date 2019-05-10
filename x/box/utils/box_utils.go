package utils

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	boxqueriers "github.com/hashgard/hashgard/x/box/client/queriers"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"

	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

func IsBoxId(boxID string) bool {
	return strings.HasPrefix(boxID, types.IDPreStr)
}

func CheckBoxId(boxID string) sdk.Error {
	if !IsBoxId(boxID) {
		return errors.ErrBoxID(boxID)
	}
	return nil
}

func GetBoxByID(cdc *codec.Codec, cliCtx context.CLIContext, boxID string) (types.Box, error) {
	var boxInfo types.Box
	// Query the box
	res, err := boxqueriers.QueryBoxByID(boxID, cliCtx)
	if err != nil {
		return nil, err
	}
	cdc.MustUnmarshalJSON(res, &boxInfo)
	return boxInfo, nil
}

func BoxOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, sender auth.Account, boxID string) (types.Box, error) {
	boxInfo, err := GetBoxByID(cdc, cliCtx, boxID)
	if err != nil {
		return nil, err
	}
	if !sender.GetAddress().Equals(boxInfo.GetOwner()) {
		return nil, errors.Errorf(errors.ErrOwnerMismatch(boxID))
	}
	return boxInfo, nil
}
func GetBoxTokenDecimal(cdc *codec.Codec, cliCtx context.CLIContext, boxID string) (uint, error) {

	boxInfo, err := GetBoxByID(cdc, cliCtx, boxID)
	if err != nil {
		return 0, err
	}

	issueInfo, err := issueutils.GetIssueByID(cdc, cliCtx, boxInfo.GetTotalAmount().Denom)
	if err != nil {
		return 0, err
	}

	return issueInfo.GetDecimals(), nil
}

func GetBoxCoinByDecimal(cdc *codec.Codec, cliCtx context.CLIContext, coin sdk.Coin) sdk.Coin {

	issueInfo, _ := issueutils.GetIssueByID(cdc, cliCtx, coin.Denom)

	return sdk.Coin{fmt.Sprintf("%s(%s)", issueInfo.GetName(), coin.Denom), issueutils.QuoDecimals(coin.Amount, issueInfo.GetDecimals())}
}
func GetBoxTypeByValue(value string) string {
	value = strings.ReplaceAll(value, types.IDPreStr, "")
	for k, v := range types.BoxType {
		if strings.HasPrefix(value, v) {
			return k
		}

	}
	return ""
}
