package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param issue for issue
type IssueParams struct {
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	TotalSupply     sdk.Int `json:"total_supply"`
	Decimals        uint    `json:"decimals"`
	Description     string  `json:"description"`
	BurnOff         bool    `json:"burning_off"`
	BurnFromOff     bool    `json:"burning_from_off"`
	BurnAnyOff      bool    `json:"burning_any_off"`
	MintingFinished bool    `json:"minting_finished"`
}
