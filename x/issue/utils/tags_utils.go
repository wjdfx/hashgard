package utils

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/hashgard/hashgard/x/issue/domain"
	issuetags "github.com/hashgard/hashgard/x/issue/tags"
)

func AppendIssueInfoTag(issueID string, coinIssueInfo *domain.CoinIssueInfo) types.Tags {
	return sdk.NewTags(
		issuetags.IssueID, issueID,
		issuetags.Name, coinIssueInfo.GetName(),
		issuetags.Symbol, coinIssueInfo.GetSymbol(),
		issuetags.TotalSupply, coinIssueInfo.GetTotalSupply().String(),
		issuetags.MintingFinished, strconv.FormatBool(coinIssueInfo.MintingFinished),
	)
}
