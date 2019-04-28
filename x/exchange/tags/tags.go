package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Exchange tags
var (
	TxCategory = "exchange"

	Action      = sdk.TagAction
	Category     = sdk.TagCategory
	Sender       = sdk.TagSender
	OrderId     = "order_id"
	SupplyToken = "supply_token"
	TargetToken = "target_token"
	OrderStatus = "order_status"
)
