package errors

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
)

// Error constructors
func ErrCanNotMint(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, domain.CanNotMint, fmt.Sprintf("Can not mint with coin %s", issueID))
}

func ErrCanNotBurn(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, domain.CanNotBurn, fmt.Sprintf("Can not burn with coin %s", issueID))
}

func ErrUnknownIssue(codespace sdk.CodespaceType, issueID string) sdk.Error {
	return sdk.NewError(codespace, domain.CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}
