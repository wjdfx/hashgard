package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

type BoxLockParams struct {
	Sender      sdk.AccAddress `json:"sender"`
	Name        string         `json:"name"`
	BoxType     string         `json:"type"`
	TotalAmount types.BoxToken `json:"total_amount"`
	Description string         `json:"description"`
	Lock        types.LockBox  `json:"lock"`
}
