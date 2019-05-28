package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Exchange tags
var (
	TxCategory = "exchange"

	Action      = sdk.TagAction
	Category    = sdk.TagCategory
	Sender      = sdk.TagSender
	OrderId     = "order_id"
	OrderStatus = "order_status"
	SupplyDenom = "supply_denom"
	TargetDenom = "target_denom"
	SupplyTurnover = "supply_turnover"
	TargetTurnover = "target_turnover"
)
