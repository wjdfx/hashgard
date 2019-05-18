package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxDisableFeature to allow a registered owner
type MsgBoxDisableFeature struct {
	BoxId   string         `json:"box_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Feature string         `json:"feature"`
}

//New MsgBoxDisableFeature Instance
func NewMsgBoxDisableFeature(boxId string, sender sdk.AccAddress, feature string) MsgBoxDisableFeature {
	return MsgBoxDisableFeature{boxId, sender, feature}
}

//nolint
func (ci MsgBoxDisableFeature) GetBoxId() string {
	return ci.BoxId
}
func (ci MsgBoxDisableFeature) SetBoxId(boxId string) {
	ci.BoxId = boxId
}
func (ci MsgBoxDisableFeature) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci MsgBoxDisableFeature) SetSender(sender sdk.AccAddress) {
	ci.Sender = sender
}
func (ci MsgBoxDisableFeature) GetFeature() string {
	return ci.Feature
}
func (ci MsgBoxDisableFeature) SetFeature(feature string) {
	ci.Feature = feature
}

// Route Implements Msg.
func (msg MsgBoxDisableFeature) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxDisableFeature) Type() string { return types.TypeMsgBoxDisableFeature }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxDisableFeature) ValidateBasic() sdk.Error {
	if len(msg.BoxId) == 0 {
		return sdk.ErrInvalidAddress("BoxId cannot be empty")
	}
	_, ok := types.Features[msg.Feature]
	if !ok {
		return errors.ErrUnknownFeatures()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxDisableFeature) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxDisableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxDisableFeature) String() string {
	return fmt.Sprintf("MsgBoxDisableFeature{%s}", msg.BoxId)
}
