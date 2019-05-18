package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BoxToken struct {
	Token    sdk.Coin `json:"token"`
	Decimals uint     `json:"decimals"`
}

//nolint
func (bi BoxToken) String() string {
	return fmt.Sprintf(`
  Token:			%s
  Decimals:			%d`,
		bi.Token.String(), bi.Decimals)
}
