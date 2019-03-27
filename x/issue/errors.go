package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName
	CanNotMint       sdk.CodeType      = 1
)

// Error constructors
func ErrCanNotMint(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, CanNotMint, fmt.Sprintf("Can not mint with coin %s", issueID))
}
