package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
)

func GetQueryBoxPath(boxID string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryBox, boxID)
}
func GetQueryBoxSearchPath(boxType string, name string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QuerySearch, boxType, name)
}
func GetQueryBoxsPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryList)
}
func GetQueryDepositListPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryDepositList)
}
func QueryBoxByName(boxType string, name string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryBoxSearchPath(boxType, name), nil)
}
func GetQueryDepositAmountPath(boxID string, accAddress sdk.AccAddress) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryDepositAmount, boxID, accAddress.String())
}

func QueryBoxByID(boxID string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryBoxPath(boxID), nil)
}
func QueryDepositAmountFromDepositBox(boxID string, accAddress sdk.AccAddress, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryDepositAmountPath(boxID, accAddress), nil)
}
func QueryBoxsList(params params.BoxQueryParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(GetQueryBoxsPath(), bz)
}
func QueryDepositList(params params.BoxQueryDepositListParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(GetQueryDepositListPath(), bz)
}
