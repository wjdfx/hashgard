package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all issue state that must be provided at genesis
type GenesisState struct {
	IssueParams IssueParams `json:"issue_params"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(ip IssueParams) GenesisState {
	return GenesisState{IssueParams: ip}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	minDepositTokens := sdk.TokensFromTendermintPower(10)
	return NewGenesisState(
		IssueParams{MinDeposit: sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, minDepositTokens)}})
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.setIssueParams(ctx, data.IssueParams)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return NewGenesisState(keeper.GetIssueParams(ctx))
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
