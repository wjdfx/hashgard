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
	CodeIssueCoinDecimalsNotValid sdk.CodeType = 5
	CodeUnknownIssue              sdk.CodeType = 6
	CanNotMint                    sdk.CodeType = 7
	CanNotBurn                    sdk.CodeType = 8
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrIssuerMismatch(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with coin %s", issueID))
}
func ErrCoinDecimalsMaxValueNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals max value is %d", types.CoinDecimalsMaxValue))
}

func ErrCoinSymbolNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueSymbolNotValid, fmt.Sprintf("Symbol max length is %d", types.CoinSymbolMaxLength))
}

func ErrCoinNamelNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueNameNotValid, fmt.Sprintf("Name max length is %d", types.CoinNameMaxLength))
}

func ErrIssueID(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueIDNotValid, fmt.Sprintf("Issue-id %s is not a valid issueId", issueID))
}

func ErrCanNotMint(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CanNotMint, fmt.Sprintf("Can not mint with coin %s", issueID))
}

//nolint
func ErrCanNotBurn(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CanNotBurn, fmt.Sprintf("Can not burn with coin %s", issueID))
}

//nolint
func ErrUnknownIssue(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}
