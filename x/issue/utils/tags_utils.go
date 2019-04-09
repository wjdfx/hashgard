package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/types"
)

//nolint
func AppendIssueInfoTag(issueID string, coinIssueInfo *types.CoinIssueInfo) sdk.Tags {
	return sdk.NewTags(
		tags.IssueID, issueID,
	)
	//return sdk.NewTags(
	//	issuetags.IssueID, issueID,
	//	issuetags.Name, coinIssueInfo.GetName(),
	//	issuetags.Symbol, coinIssueInfo.GetSymbol(),
	//	issuetags.TotalSupply, coinIssueInfo.GetTotalSupply().String(),
	//	issuetags.MintingFinished, strconv.FormatBool(coinIssueInfo.MintingFinished),
	//)
}
