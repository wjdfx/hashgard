package errors

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

const (
	CodeBoxOwnerMismatch          sdk.CodeType = 1
	CodeBoxIDNotValid             sdk.CodeType = 2
	CodeBoxNameNotValid           sdk.CodeType = 3
	CodeAmountNotValid            sdk.CodeType = 4
	CodeBoxDescriptionNotValid    sdk.CodeType = 5
	CodeUnknownBox                sdk.CodeType = 6
	CodeUnknownBoxType            sdk.CodeType = 7
	CodeUnknownOperation          sdk.CodeType = 8
	CodeInterestInjectionNotValid sdk.CodeType = 9
	CodeInterestFetchNotValid     sdk.CodeType = 10
	CodeNotEnoughAmount           sdk.CodeType = 11
	CodeTimeNotValid              sdk.CodeType = 12
	CodeNotAllowedOperation       sdk.CodeType = 13
	CodeNotSupportOperation       sdk.CodeType = 14
	CodeUnknownFeature            sdk.CodeType = 15
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrOwnerMismatch(boxID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxOwnerMismatch, fmt.Sprintf("Owner mismatch with box %s", boxID))
}
func ErrTimeNotValid(timeKey string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeTimeNotValid, "%s is not a valid timestamp", timeKey)
}
func ErrAmountNotValid(key string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeAmountNotValid, "%s is not a valid amount", key)
}
func ErrInterestInjectionNotValid(coin sdk.Coin) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeInterestInjectionNotValid, "%s is not a valid interest injection", coin.String())
}
func ErrInterestFetchNotValid(coin sdk.Coin) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeInterestFetchNotValid, "%s is not a valid interest fetch", coin.String())
}
func ErrBoxPriceNotValid(name string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxNameNotValid, fmt.Sprintf("Price mismatch with box %s", name))
}
func ErrBoxNameNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxNameNotValid, fmt.Sprintf("Name max length is %d", types.BoxNameMaxLength))
}
func ErrBoxDescriptionNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxDescriptionNotValid, "Description is not valid json")
}
func ErrBoxDescriptionMaxLengthNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxDescriptionNotValid, "Description max length is %d", types.BoxDescriptionMaxLength)
}
func ErrBoxID(boxID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxIDNotValid, fmt.Sprintf("Box-id %s is not a valid boxId", boxID))
}
func ErrUnknownBox(boxID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownBox, fmt.Sprintf("Unknown box with id %s", boxID))
}
func ErrUnknownBoxType() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownBoxType, fmt.Sprintf("Unknown type"))
}
func ErrUnknownOperation() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownOperation, fmt.Sprintf("Unknown operation"))
}
func ErrNotEnoughAmount() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotEnoughAmount, fmt.Sprintf("Not enough amount"))
}
func ErrNotSupportOperation() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotSupportOperation, fmt.Sprintf("Not support operation the box"))
}
func ErrNotAllowedOperation(status string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotAllowedOperation, fmt.Sprintf("Not allowed operation in current status: %s", status))
}
func ErrUnknownFeatures() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownFeature, fmt.Sprintf("Unknown feature"))
}
