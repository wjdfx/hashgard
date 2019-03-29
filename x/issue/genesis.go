package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keepers"
	"github.com/hashgard/hashgard/x/issue/params"
)

// GenesisState - all issue state that must be provided at genesis
type GenesisState struct {
	IssueParams params.IssueParams `json:"issue_params"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(ip params.IssueParams) GenesisState {
	return GenesisState{IssueParams: ip}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	minDepositTokens := sdk.TokensFromTendermintPower(10)
	return NewGenesisState(
		params.IssueParams{MinDeposit: sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, minDepositTokens)}})
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keepers.Keeper, data GenesisState) {
	keeper.SetIssueParams(ctx, data.IssueParams)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keepers.Keeper) GenesisState {
	return NewGenesisState(keeper.GetIssueParams(ctx))
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
