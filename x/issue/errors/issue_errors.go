package errors

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

const (
	CodeIssuerMismatch            sdk.CodeType = 1
	CodeIssueIDNotValid           sdk.CodeType = 2
	CodeIssueNameNotValid         sdk.CodeType = 3
	CodeIssueSymbolNotValid       sdk.CodeType = 4
	CodeIssueTotalSupplyNotValid  sdk.CodeType = 5
	CodeIssueCoinDecimalsNotValid sdk.CodeType = 6
	CodeIssueDescriptionNotValid  sdk.CodeType = 7
	CodeUnknownIssue              sdk.CodeType = 8
	CanNotMint                    sdk.CodeType = 9
	CanNotBurn                    sdk.CodeType = 10
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrOwnerMismatch(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with coin %s", issueID))
}
func ErrCoinCanNotBurnOverFromAmount(issueID string, amount sdk.Int) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Can not burn %d from %s", amount.String(), issueID))
}
func ErrCoinDecimalsMaxValueNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals max value is %d", types.CoinDecimalsMaxValue))
}
func ErrCoinTotalSupplyMaxValueNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueTotalSupplyNotValid, fmt.Sprintf("Total supply max value is %s", types.CoinMaxTotalSupply.String()))
}
func ErrCoinSymbolNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueSymbolNotValid, fmt.Sprintf("Symbol length is %d-%d character", types.CoinSymbolMinLength, types.CoinSymbolMaxLength))
}

func ErrCoinNamelNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueNameNotValid, fmt.Sprintf("Name max length is %d", types.CoinNameMaxLength))
}
func ErrCoinDescriptionNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueDescriptionNotValid, "Description is not valid json")
}
func ErrCoinDescriptionMaxLengthNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueDescriptionNotValid, "Description max length is %d", types.CoinDescriptionMaxLength)
}
func ErrIssueID(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueIDNotValid, fmt.Sprintf("Issue-id %s is not a valid issueId", issueID))
}

func ErrCanNotMint(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CanNotMint, fmt.Sprintf("Can not mint with coin %s", issueID))
}
func ErrCanNotBurn(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CanNotBurn, fmt.Sprintf("Can not burn with coin %s", issueID))
}
func ErrUnknownIssue(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}
