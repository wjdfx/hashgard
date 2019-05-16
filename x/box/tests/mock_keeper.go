package tests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/types"
)

type IssueKeeper struct {
}

//New issue keeper Instance
func NewIssueKeeper() IssueKeeper {
	return IssueKeeper{}
}

//Returns issue by issueID
func (keeper IssueKeeper) GetIssue(ctx sdk.Context, issueID string) *types.CoinIssueInfo {

	coinIssueInfo := types.CoinIssueInfo{
		Decimals: TestTokenDecimals,
	}
	return &coinIssueInfo
}
