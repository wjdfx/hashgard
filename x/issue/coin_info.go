package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CoinIssueInfo struct {
	Issuer          sdk.AccAddress `json:"issuer"`
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	TotalSupply     sdk.Int        `json:"total_supply"`
	Decimals        uint           `json:"decimals"`
	MintingFinished bool           `json:"minting_finished"`
}

// CoinInfo stores meta data about a coin
type CoinInfo struct {
	IssueId string `json:"issue_id"`
	CoinIssueInfo
}

//TODO
func (coinIssueInfo CoinIssueInfo) String() string {
	return fmt.Sprintf("IssueId:%s,Issuer:%s", coinIssueInfo.Issuer, coinIssueInfo.Issuer.String())
}
