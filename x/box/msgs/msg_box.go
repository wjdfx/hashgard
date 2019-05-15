package msgs

//
//import (
//	"fmt"
//	"time"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//
//	"github.com/hashgard/hashgard/x/box/errors"
//	"github.com/hashgard/hashgard/x/box/types"
//)
//
//// MsgBox to allow a registered boxr
//// to box new coins.
//type MsgBox struct {
//	*types.BoxInfo
//}
//
////New MsgBox Instance
//func NewMsgBox(boxInfo *types.BoxInfo) MsgBox {
//	return MsgBox{boxInfo}
//}
//
//// Route Implements Msg.
//func (msg MsgBox) Route() string { return types.RouterKey }
//
//// Type Implements Msg.789
//func (msg MsgBox) Type() string { return types.TypeMsgBox }
//
//// Implements Msg. Ensures addresses are valid and Coin is positive
//func (msg MsgBox) ValidateBasic() sdk.Error {
//	if len(msg.Owner) == 0 {
//		return sdk.ErrInvalidAddress("Owner address cannot be empty")
//	}
//	if err := types.CheckBoxType(msg.BoxType); err != nil {
//		return err
//	}
//	if err := msg.validateLockBox(); err != nil {
//		return err
//	}
//	if err := msg.validateDepositBox(); err != nil {
//		return err
//	}
//	if err := msg.validateFutureBox(); err != nil {
//		return err
//	}
//	if msg.TotalAmount.IsZero() || msg.TotalAmount.Amount.IsNegative() {
//		return errors.ErrAmountNotValid("Token amount")
//	}
//	if len(msg.Name) > types.BoxNameMaxLength {
//		return errors.ErrBoxNameNotValid()
//	}
//	if len(msg.Description) > types.BoxDescriptionMaxLength {
//		return errors.ErrBoxDescriptionMaxLengthNotValid()
//	}
//	return nil
//}
//func (msg MsgBox) validateLockBox() sdk.Error {
//	if types.Lock != msg.BoxType {
//		return nil
//	}
//	if msg.Lock.EndTime.Before(time.Now()) {
//		return errors.ErrTimeNotValid("EndTime")
//	}
//	return nil
//}
//
//func (msg MsgBox) validateDepositBox() sdk.Error {
//	if types.Deposit != msg.BoxType {
//		return nil
//	}
//	if msg.Deposit.StartTime.Before(time.Now()) {
//		return errors.ErrTimeNotValid("StartTime")
//	}
//	if msg.Deposit.EstablishTime.Before(time.Now()) || msg.Deposit.EstablishTime.Before(msg.Deposit.StartTime) {
//		return errors.ErrTimeNotValid("EstablishTime")
//	}
//	if msg.Deposit.MaturityTime.Before(time.Now()) || msg.Deposit.MaturityTime.Before(msg.Deposit.EstablishTime) {
//		return errors.ErrTimeNotValid("MaturityTime")
//	}
//	if msg.Deposit.BottomLine.IsNegative() || msg.Deposit.BottomLine.GT(msg.TotalAmount.Amount) {
//		return errors.ErrAmountNotValid("BottomLine")
//	}
//	if msg.Deposit.Interest.IsNegative() || msg.Deposit.Interest.IsZero() {
//		return errors.ErrAmountNotValid("Interest")
//	}
//	if msg.Deposit.Price.IsZero() || msg.Deposit.Price.IsNegative() || !msg.TotalAmount.Amount.Mod(msg.Deposit.Price).IsZero() {
//		return errors.ErrAmountNotValid("Price")
//	}
//
//	return nil
//}
//func (msg MsgBox) validateFutureBox() sdk.Error {
//	if types.Future != msg.BoxType {
//		return nil
//	}
//	if msg.Future.TimeLine == nil || msg.Future.Receivers == nil ||
//		len(msg.Future.TimeLine) == 0 || len(msg.Future.Receivers) == 0 {
//		return errors.ErrNotSupportOperation()
//	}
//	if len(msg.Future.TimeLine) != len(msg.Future.Receivers) {
//		return errors.ErrNotSupportOperation()
//	}
//	if len(msg.Future.TimeLine) > types.BoxMaxInstalment {
//		return errors.ErrNotEnoughAmount()
//	}
//	if len(msg.Future.Receivers) > types.BoxMaxInstalment {
//		return errors.ErrNotEnoughAmount()
//	}
//	return nil
//}
//
//// GetSignBytes Implements Msg.
//func (msg MsgBox) GetSignBytes() []byte {
//	bz := MsgCdc.MustMarshalJSON(msg)
//	return sdk.MustSortJSON(bz)
//}
//
//// GetSigners Implements Msg.
//func (msg MsgBox) GetSigners() []sdk.AccAddress {
//	return []sdk.AccAddress{msg.Owner}
//}
//
//func (msg MsgBox) String() string {
//	return fmt.Sprintf("MsgBox{%s - %s}", "", msg.Owner.String())
//}
