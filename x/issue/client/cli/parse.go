package cli

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue"
	"github.com/spf13/viper"
	"io/ioutil"
)

//TODO
func parseIssueFlags() (*issue.CoinIssueInfo, error) {
	issues := &issue.CoinIssueInfo{}
	issueFile := viper.GetString(flagIssue)

	if issueFile == "" {
		issues.Name = viper.GetString(flagName)
		res, ok := sdk.NewIntFromString(viper.GetString(flagTotalSupply))
		if !ok {
			sdk.ErrInvalidCoins(flagTotalSupply)
		}
		issues.TotalSupply = res
		return issues, nil
	}

	for _, flag := range issueFlags {
		if viper.GetString(flag) != "" {
			return nil, fmt.Errorf("--%s flag provided alongside --issue, which is a noop", flag)
		}
	}

	contents, err := ioutil.ReadFile(issueFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
