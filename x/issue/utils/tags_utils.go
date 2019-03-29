package utils

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/domain"
	issuetags "github.com/hashgard/hashgard/x/issue/tags"
)

func AppendIssueInfoTag(issueID string, coinIssueInfo domain.CoinIssueInfo) types.Tags {
	tags := sdk.EmptyTags()
	tags = tags.AppendTag(issuetags.IssueID, issueID)
	tags = tags.AppendTag(issuetags.Name, coinIssueInfo.Name)
	tags = tags.AppendTag(issuetags.TotalSupply, coinIssueInfo.TotalSupply.String())
	return tags
}
