package domain

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
	IDPreStr = "gard"
)
const (
	TypeMsgIssue              = "issue"
	TypeMsgIssueMint          = "issueMint"
	TypeMsgIssueBurn          = "issueBurn"
	TypeMsgIssueFinishMinting = "issueFinishMinting"
	DefaultDecimals           = 18
)
const (
	DefaultCodespace sdk.CodespaceType = ModuleName
	CodeUnknownIssue sdk.CodeType      = 1
	CanNotMint       sdk.CodeType      = 2
	CanNotBurn       sdk.CodeType      = 3
)
