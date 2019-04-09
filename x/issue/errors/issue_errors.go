package errors

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/types"
)

const (
	CodeIssuerMismatch      sdk.CodeType = 1
	CodeIssueIDNotValid     sdk.CodeType = 2
	CodeIssueNameNotValid   sdk.CodeType = 3
	CodeIssueSymbolNotValid sdk.CodeType = 4
	CodeUnknownIssue        sdk.CodeType = 5
	CanNotMint              sdk.CodeType = 6
	CanNotBurn              sdk.CodeType = 7
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrIssuerMismatch(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CodeIssuerMismatch, fmt.Sprintf("Issuer mismatch with coin %s", issueID))
}

//nolint
func ErrCoinSymbolNotValid(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeIssueSymbolNotValid, fmt.Sprintf("Symbol max length is %d", types.CoinSymbolMaxLength))
}

//nolint
func ErrCoinNamelNotValid(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeIssueNameNotValid, fmt.Sprintf("Name max length is %d", types.CoinNameMaxLength))
}

//nolint
func ErrIssueID(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CodeIssueIDNotValid, fmt.Sprintf("Issue-id %s is not a valid issueId", issueID))
}

//nolint
func ErrCanNotMint(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CanNotMint, fmt.Sprintf("Can not mint with coin %s", issueID))
}

//nolint
func ErrCanNotBurn(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CanNotBurn, fmt.Sprintf("Can not burn with coin %s", issueID))
}

//nolint
func ErrUnknownIssue(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}
