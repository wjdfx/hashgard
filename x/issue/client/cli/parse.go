package cli

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"io/ioutil"

	"github.com/hashgard/hashgard/x/issue/domain"
)

const (
	flagName        = "name"
	flagTotalSupply = "total-supply"
	flagIssue       = "issue"
)

var issueFlags = []string{
	flagName,
	flagTotalSupply,
}

//TODO
func parseIssueFlags() (*domain.CoinIssueInfo, error) {
	issues := &domain.CoinIssueInfo{}
	issueFile := viper.GetString(flagIssue)

	if issueFile == "" {
		name := viper.GetString(flagName)
		issues.Name = name
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
