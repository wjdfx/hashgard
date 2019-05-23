package params

import (
	"github.com/hashgard/hashgard/x/box/types"
)

type BoxLockParams struct {
	Name        string         `json:"name"`
	BoxType     string         `json:"type"`
	TotalAmount types.BoxToken `json:"total_amount"`
	Description string         `json:"description"`
	Lock        types.LockBox  `json:"lock"`
}
