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
	CoinMaxTotalSupply, _        = sdk.NewIntFromString("1000000000000000000000000000000000000")
	CoinIssueMaxId        uint64 = 999999999999 //9999-12-31 23:59:59
	CoinIssueMinId        uint64 = 100000000000 //2019-01-01 00:00:00
)

const (
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
	TypeMsgIssueMint              = "issue_mint"
	TypeMsgIssueBurn              = "issue_burn"
	TypeMsgIssueBurnFrom          = "issue_burn_from"
	TypeMsgIssueBurnAny           = "issue_burn_any"
	TypeMsgIssueFinishMinting     = "issue_finish_minting"
	TypeMsgIssueDescription       = "issue_description"
	TypeMsgIssueTransferOwnership = "issue_transfer_ownership"
	TypeMsgIssueBurnOff           = "issue_burn_off"
	TypeMsgIssueBurnFromOff       = "issue_burn_from_off"
	TypeMsgIssueBurnAnyOff        = "issue_burn_any_off"
)
const (
	CoinDecimalsMaxValue                  = uint(18)
	CodeInvalidGenesis       sdk.CodeType = 102
	CoinNameMaxLength                     = 32
	CoinSymbolMinLength                   = 2
	CoinSymbolMaxLength                   = 8
	CoinDescriptionMaxLength              = 1024
)
