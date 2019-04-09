package msgs

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/errors"
	"github.com/hashgard/hashgard/x/issue/types"
)

// MsgIssue to allow a registered issuer
// to issue new coins.
type MsgIssue struct {
	*types.CoinIssueInfo
}

//New MsgIssue Instance
func NewMsgIssue(coinIssueInfo *types.CoinIssueInfo) MsgIssue {
	return MsgIssue{coinIssueInfo}
}

// Route Implements Msg.
func (msg MsgIssue) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssue) Type() string { return types.TypeMsgIssue }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssue) ValidateBasic() sdk.Error {
	if len(msg.Issuer) == 0 {
		return sdk.ErrInvalidAddress("Issuer address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.CoinIssueInfo.TotalSupply.IsZero() || !msg.CoinIssueInfo.TotalSupply.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	}
	if len(msg.Name) > types.CoinNameMaxLength {
		return errors.ErrCoinNamelNotValid(types.DefaultParamspace)
	}
	if len(msg.Symbol) > types.CoinSymbolMaxLength {
		return errors.ErrCoinSymbolNotValid(types.DefaultParamspace)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssue) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

func (msg MsgIssue) String() string {
	return fmt.Sprintf("MsgIssue{%s - %s}", "", msg.Issuer.String())
}
