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
	BoxIdLength                    = 14
	BoxMaxId                uint64 = 99999999999999
	BoxMinId                uint64 = 10000000000000
	BoxMaxInstalment               = 99
	BoxMaxInjectionInterest        = 100
)

const (
	IDPreStr = "box"
	Custom   = "custom"
	Gard     = "gard"
	Agard    = "agard"
)
const (
	QueryParams = "params"
	QueryList   = "list"
	QueryBox    = "query"
	QuerySearch = "search"
)

//box status
const (
	BoxCreated    = "created"
	BoxDepositing = "depositing"
	BoxActived    = "actived"
	BoxUndue      = "undue"
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
	TypeMsgBoxCreateLock    = "box_create_lock"
	TypeMsgBoxCreateDeposit = "box_create_deposit"
	TypeMsgBoxCreateFuture  = "box_create_future"

	TypeMsgBoxWithdraw = "box_withdraw"

	TypeMsgBoxInterest = "box_interest"
	TypeMsgBoxDeposit  = "box_deposit"

	TypeMsgBoxDescription    = "box_description"
	TypeMsgBoxDisableFeature = "box_disable_feature"
)
const (
	KeyDelimiterString                   = ":"
	AgardDecimals                        = uint(1)
	GardDecimals                         = uint(18)
	MaxPrecision                         = uint(6)
	CodeInvalidGenesis      sdk.CodeType = 102
	BoxNameMaxLength                     = 32
	BoxDescriptionMaxLength              = 1024
)
