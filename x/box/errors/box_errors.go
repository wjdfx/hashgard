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
	CodeDecimalsNotValid          sdk.CodeType = 5
	CodeTimelineNotValid          sdk.CodeType = 6
	CodeBoxDescriptionNotValid    sdk.CodeType = 7
	CodeUnknownBox                sdk.CodeType = 8
	CodeUnknownBoxType            sdk.CodeType = 9
	CodeUnknownOperation          sdk.CodeType = 10
	CodeInterestInjectionNotValid sdk.CodeType = 11
	CodeInterestFetchNotValid     sdk.CodeType = 12
	CodeNotEnoughAmount           sdk.CodeType = 13
	CodeTimeNotValid              sdk.CodeType = 14
	CodeNotAllowedOperation       sdk.CodeType = 15
	CodeNotSupportOperation       sdk.CodeType = 16
	CodeUnknownFeature            sdk.CodeType = 17
	CodeNotTransfer               sdk.CodeType = 18
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrOwnerMismatch(id string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxOwnerMismatch, fmt.Sprintf("Owner mismatch with box %s", id))
}
func ErrDecimalsNotValid(decimals uint) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeDecimalsNotValid, "%d is not a valid decimals", decimals)
}
func ErrTimelineNotValid(time []int64) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeTimelineNotValid, "%d is not a valid time line", time)
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
	return sdk.NewError(types.DefaultCodespace, CodeBoxNameNotValid, fmt.Sprintf("Price mismatch with %s", name))
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
func ErrBoxID(id string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeBoxIDNotValid, fmt.Sprintf("id %s is not a valid id", id))
}
func ErrUnknownBox(id string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownBox, fmt.Sprintf("Unknown box with id %s", id))
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
	return sdk.NewError(types.DefaultCodespace, CodeNotSupportOperation, fmt.Sprintf("Not support operation"))
}
func ErrNotAllowedOperation(status string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotAllowedOperation, fmt.Sprintf("Not allowed operation in current status: %s", status))
}
func ErrUnknownFeatures() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownFeature, fmt.Sprintf("Unknown feature"))
}
func ErrCanNotTransfer(id string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeNotTransfer, fmt.Sprintf("The box %s Can not be transfer", id))
}
