package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Exchange tags
var (
	Action      = sdk.TagAction
	OrderId     = "order_id"
	Seller      = "seller"
	Buyer       = "buyer"
	SupplyToken = "supply_token"
	TargetToken = "target_token"
	OrderStatus = "order_status"
)
