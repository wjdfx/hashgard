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
	CodeUnknownFeature            sdk.CodeType = 11
	CodeUnknownFreezeType         sdk.CodeType = 12
	CodeNotEnoughAmountToTransfer sdk.CodeType = 13
	CodeCanNotFreeze              sdk.CodeType = 14
	CodeFreezeEndTimeNotValid     sdk.CodeType = 15
	CodeNotTransferIn             sdk.CodeType = 16
	CodeNotTransferOut            sdk.CodeType = 17
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrOwnerMismatch(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with token %s", issueID))
}
func ErrCoinDecimalsMaxValueNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals max value is %d", types.CoinDecimalsMaxValue))
}
func ErrCoinDecimalsMultipleNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals must be a multiple of %d", types.CoinDecimalsMultiple))
}
func ErrCoinTotalSupplyMaxValueNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueTotalSupplyNotValid, fmt.Sprintf("Total supply max value is %s", types.CoinMaxTotalSupply.String()))
}
func ErrCoinSymbolNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeIssueSymbolNotValid, fmt.Sprintf("Symbol length is %d-%d character", types.CoinSymbolMinLength, types.CoinSymbolMaxLength))
}
func ErrFreezeEndTimestampNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeFreezeEndTimeNotValid, "end-time is not a valid timestamp")
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
	return sdk.NewError(types.DefaultCodespace, CanNotMint, fmt.Sprintf("Can not mint the token %s", issueID))
}
func ErrCanNotBurn(issueID string, burnType string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CanNotBurn, fmt.Sprintf("Can not burn the token %s by %s", issueID, burnType))
}
func ErrUnknownIssue(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}
func ErrUnknownFeatures() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownFeature, fmt.Sprintf("Unknown feature"))
}
func ErrCanNotFreeze(issueID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeCanNotFreeze, fmt.Sprintf("Can not freeze the token %s", issueID))
}
func ErrUnknownFreezeType() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownFreezeType, fmt.Sprintf("Unknown type"))
}
func ErrNotEnoughAmountToTransfer() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotEnoughAmountToTransfer, fmt.Sprintf("Not enough amount allowed to transfer"))
}
func ErrCanNotTransferIn(issueID string, accAddress string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotTransferIn, fmt.Sprintf("Can not transfer %s to %s", issueID, accAddress))
}
func ErrCanNotTransferOut(issueID string, accAddress string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotTransferOut, fmt.Sprintf("Can not transfer %s from %s", issueID, accAddress))
}
