package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidInput   sdk.CodeType = 101
	CodeInvalidGenesis sdk.CodeType = 102
	CodeOrderNotExist  sdk.CodeType = 103
	CodeNoPermission   sdk.CodeType = 104
	CodeNotMatchTarget sdk.CodeType = 105
	CodeTooLess        sdk.CodeType = 106
)
