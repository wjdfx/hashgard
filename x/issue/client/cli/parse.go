package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"

	"github.com/hashgard/hashgard/x/issue/params"
)

const (
	flagName            = "name"
	flagSymbol          = "symbol"
	flagTotalSupply     = "total-supply"
	flagMintingFinished = "minting_finished"
)

var issueFlags = []string{
	flagName,
	flagSymbol,
	flagTotalSupply,
	flagMintingFinished,
}

func parseIssueFlags() (params.IssueParams, error) {
	issues := params.IssueParams{}

	issues.Name = viper.GetString(flagName)
	issues.Symbol = viper.GetString(flagSymbol)
	issues.MintingFinished = false
	if viper.IsSet(flagMintingFinished) {
		issues.MintingFinished = viper.GetBool(flagMintingFinished)
	}
	res, ok := sdk.NewIntFromString(viper.GetString(flagTotalSupply))
	if !ok {
		sdk.ErrInvalidCoins(flagTotalSupply)
	}
	issues.TotalSupply = res
	return issues, nil
}
