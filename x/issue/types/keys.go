package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "issue"
	// StoreKey is the store key string for issue
	StoreKey = ModuleName
	// RouterKey is the message route for issue
	RouterKey = ModuleName
	// QuerierRoute is the querier route for issue
	QuerierRoute = ModuleName
	// Parameter store default namestore
	DefaultParamspace = ModuleName
)
const (
	DefaultCodespace sdk.CodespaceType = ModuleName
)

var (
	CoinMaxTotalSupply, _ = sdk.NewIntFromString("10000000000000000000000000000000000")
)

const (
	IDLength = 15
	IDPreStr = "coin"
	Custom   = "custom"
)
const (
	QueryParams = "params"
	QueryIssues = "list"
	QueryIssue  = "query"
)
const (
	TypeMsgIssue              = "issue"
	TypeMsgIssueMint          = "issueMint"
	TypeMsgIssueBurn          = "issueBurn"
	TypeMsgIssueFinishMinting = "issueFinishMinting"
	TypeMsgIssueDescription   = "issueDescription"
)
const (
	CoinDecimalsMaxValue     = uint(18)
	CoinNameMaxLength        = 32
	CoinSymbolMinLength      = 2
	CoinSymbolMaxLength      = 8
	CoinDescriptionMaxLength = 1024
)
