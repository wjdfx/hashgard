package issue

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Governance message types and routes
const (
	TypeMsgIssue    = "issue"
	DefaultDecimals = 18
)

// MsgIssue to allow a registered issuer
// to issue new coins.
type MsgIssue struct {
	Issuer        sdk.AccAddress `json:"issuer_address"`
	CoinIssueInfo *CoinIssueInfo `json:"coin_info"`
}

func NewMsgIssue(issuer sdk.AccAddress, coinIssueInfo *CoinIssueInfo) MsgIssue {
	return MsgIssue{
		Issuer:        issuer,
		CoinIssueInfo: coinIssueInfo,
	}
}

// Route Implements Msg.
func (msg MsgIssue) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgIssue) Type() string { return TypeMsgIssue }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssue) ValidateBasic() sdk.Error {
	if len(msg.Issuer) == 0 {
		return sdk.ErrInvalidAddress("Issuer address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.CoinIssueInfo.TotalSupply.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssue) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

func (msg MsgIssue) String() string {
	return fmt.Sprintf("MsgIssue{%s - %s}", msg.Issuer, msg.CoinIssueInfo.String())
}
