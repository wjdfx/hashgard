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
	CoinMaxTotalSupply, _       = sdk.NewIntFromString("1000000000000000000000000000000000000")
	CoinIssueMaxTimestamp int64 = 253402271999 //9999-12-31 23:59:59
	CoinIssueMinTimestamp int64 = 1546272000   //2019-01-01 00:00:00
)

const (
	IDLength = 16
	IDPreStr = "coin"
	Custom   = "custom"
)
const (
	QueryParams = "params"
	QueryIssues = "list"
	QueryIssue  = "query"
	QuerySearch = "search"
)
const (
	BurnOwner = "burnOwner"
	BurnFrom  = "burnFrom"
	BurnAny   = "burnAny"
)

const (
	TypeMsgIssue                  = "issue"
	TypeMsgIssueMint              = "issueMint"
	TypeMsgIssueBurn              = "issueBurn"
	TypeMsgIssueBurnFrom          = "issueBurnFrom"
	TypeMsgIssueBurnAny           = "issueBurnAny"
	TypeMsgIssueFinishMinting     = "issueFinishMinting"
	TypeMsgIssueDescription       = "issueDescription"
	TypeMsgIssueTransferOwnership = "issueTransferOwnership"
	TypeMsgIssueBurnOff           = "issueBurnOff"
	TypeMsgIssueBurnFromOff       = "issueBurnFromOff"
	TypeMsgIssueBurnAnyOff        = "issueBurnAnyOff"
)
const (
	CoinDecimalsMaxValue                  = uint(18)
	CodeInvalidGenesis       sdk.CodeType = 102
	CoinNameMaxLength                     = 32
	CoinSymbolMinLength                   = 2
	CoinSymbolMaxLength                   = 8
	CoinDescriptionMaxLength              = 1024
)
