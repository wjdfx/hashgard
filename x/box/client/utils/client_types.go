package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LockBoxInfo struct {
	BoxId         string         `json:"box_id"`
	BoxStatus     string         `json:"box_status"`
	Owner         sdk.AccAddress `json:"owner"`
	Name          string         `json:"name"`
	BoxType       string         `json:"type"`
	CreatedTime   int64          `json:"created_time"`
	TotalAmount   types.BoxToken `json:"total_amount"`
	Description   string         `json:"description"`
	TradeDisabled bool           `json:"trade_disabled"`
	Lock          types.LockBox  `json:"lock"`
}
type DepositBoxInfo struct {
	BoxId         string           `json:"box_id"`
	BoxStatus     string           `json:"box_status"`
	Owner         sdk.AccAddress   `json:"owner"`
	Name          string           `json:"name"`
	BoxType       string           `json:"type"`
	CreatedTime   int64            `json:"created_time"`
	TotalAmount   types.BoxToken   `json:"total_amount"`
	Description   string           `json:"description"`
	TradeDisabled bool             `json:"trade_disabled"`
	Deposit       types.DepositBox `json:"deposit"`
}
type FutureBoxInfo struct {
	BoxId         string          `json:"box_id"`
	BoxStatus     string          `json:"box_status"`
	Owner         sdk.AccAddress  `json:"owner"`
	Name          string          `json:"name"`
	BoxType       string          `json:"type"`
	CreatedTime   int64           `json:"created_time"`
	TotalAmount   types.BoxToken  `json:"total_amount"`
	Description   string          `json:"description"`
	TradeDisabled bool            `json:"trade_disabled"`
	Future        types.FutureBox `json:"future"`
}
type LockBoxInfos []LockBoxInfo
type DepositBoxInfos []DepositBoxInfo
type FutureBoxInfos []FutureBoxInfo

//nolint
func getString(BoxId string, BoxStatus string, Owner sdk.AccAddress, Name string, BoxType string, CreatedTime int64,
	TotalAmount types.BoxToken, Description string, TradeDisabled bool) string {
	return fmt.Sprintf(`BoxInfo:
  BoxId:			%s
  BoxStatus:			%s
  Owner:			%s
  Name:				%s
  BoxType:			%s
  TotalAmount:			%s
  CreatedTime:			%d
  Description:			%s
  TradeDisabled:		%t`,
		BoxId, BoxStatus, Owner.String(), Name, BoxType, TotalAmount.String(),
		CreatedTime, Description, TradeDisabled)
}

//nolint
func (bi LockBoxInfo) String() string {
	str := getString(bi.BoxId, bi.BoxStatus, bi.Owner, bi.Name, bi.BoxType,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TradeDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Lock.String())
}

//nolint
func (bi DepositBoxInfo) String() string {
	str := getString(bi.BoxId, bi.BoxStatus, bi.Owner, bi.Name, bi.BoxType,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TradeDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Deposit.String())
}

//nolint
func (bi FutureBoxInfo) String() string {
	str := getString(bi.BoxId, bi.BoxStatus, bi.Owner, bi.Name, bi.BoxType,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TradeDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Future.String())
}

//nolint
func (bi LockBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "EndTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%s\n",
			box.BoxId, box.Owner.String(), box.Name, box.TotalAmount.Token.String(), time.Unix(box.Lock.EndTime, 0).String())
	}
	return strings.TrimSpace(out)
}

//nolint
func (bi DepositBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "CreatedTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%d\n",
			box.BoxId, box.Owner.String(), box.Name, box.TotalAmount.String(), box.CreatedTime)
	}
	return strings.TrimSpace(out)
}

//nolint
func (bi FutureBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "CreatedTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-36s|%d\n",
			box.BoxId, box.Owner.String(), box.Name, box.TotalAmount.String(), box.CreatedTime)
	}
	return strings.TrimSpace(out)
}
