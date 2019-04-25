package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/issue/utils"

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
	if len(msg.Owner) == 0 {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.CoinIssueInfo.TotalSupply.IsZero() || !msg.CoinIssueInfo.TotalSupply.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	}
	if utils.QuoDecimals(msg.CoinIssueInfo.TotalSupply, msg.CoinIssueInfo.Decimals).GT(types.CoinMaxTotalSupply) {
		return errors.ErrCoinTotalSupplyMaxValueNotValid()
	}
	if len(msg.Name) > types.CoinNameMaxLength {
		return errors.ErrCoinNamelNotValid()
	}
	if len(msg.Symbol) < types.CoinSymbolMinLength || len(msg.Symbol) > types.CoinSymbolMaxLength {
		return errors.ErrCoinSymbolNotValid()
	}
	if msg.Decimals > types.CoinDecimalsMaxValue {
		return errors.ErrCoinDecimalsMaxValueNotValid()
	}
	if len(msg.Description) > types.CoinDescriptionMaxLength {
		return errors.ErrCoinDescriptionMaxLengthNotValid()
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
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgIssue) String() string {
	return fmt.Sprintf("MsgIssue{%s - %s}", "", msg.Owner.String())
}
