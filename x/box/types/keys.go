package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "box"
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
	BoxMaxId                uint64 = 99999999999999
	BoxMinId                uint64 = 10000000000000
	BoxMaxInstalment               = 99
	BoxMaxInjectionInterest        = 100
)

const (
	IDPreStr = "box"
	Custom   = "custom"
)
const (
	QueryParams        = "params"
	QueryList          = "list"
	QueryBox           = "query"
	QueryDepositList   = "deposit"
	QueryDepositAmount = "deposit-amount"
	QuerySearch        = "search"
)

//box status
const (
	BoxCreated    = "created"
	BoxDepositing = "depositing"
	BoxActived    = "actived"
	BoxClosed     = "closed"
	BoxFinished   = "finished"
)

//lock box status
const (
	LockBoxLocked   = "locked"
	LockBoxUnlocked = "unlocked"
)

//deposit box status
const (
	DepositBoxInterest = "interest"
)
const (
	Injection = "injection"
	DepositTo = "deposit-to"
	Fetch     = "fetch"
)

const (
	TypeMsgBox               = "box_create"
	TypeMsgBoxInterest       = "box_interest"
	TypeMsgBoxDeposit        = "box_deposit"
	TypeMsgBoxFuture         = "box_future"
	TypeMsgBoxDescription    = "box_description"
	TypeMsgBoxDisableFeature = "box_disable_feature"
)
const (
	KeyDelimiterString                   = ":"
	MaxPrecision                         = uint(6)
	CodeInvalidGenesis      sdk.CodeType = 102
	BoxNameMaxLength                     = 32
	BoxDescriptionMaxLength              = 1024
)
