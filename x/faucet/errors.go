package faucet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidInput	sdk.CodeType = 101
	CodeInvalidGenesis	sdk.CodeType = 102
	CodeHaveTooMany		sdk.CodeType = 103
)
