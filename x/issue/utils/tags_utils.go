package utils

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	issuetags "github.com/hashgard/hashgard/x/issue/tags"
)

func AppendIssueInfoTag(issueID string, coinIssueInfo domain.CoinIssueInfo) types.Tags {
	return sdk.NewTags(
		issuetags.IssueID, issueID,
		issuetags.Name, coinIssueInfo.Name,
		issuetags.TotalSupply, coinIssueInfo.TotalSupply.String(),
	)
}
