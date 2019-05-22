package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Box tags
var (
	TxCategory = "box"

	Action    = sdk.TagAction
	Category  = sdk.TagCategory
	Sender    = sdk.TagSender
	Feature   = "feature"
	Owner     = "owner"
	Operation = "operation"
	Interest  = "interest"
	BoxID     = "box-id"
	BoxType   = "box-type"
	BoxStatus = "box-status"
	Seq       = "seq"
)
