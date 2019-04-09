package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/params"
)

// GenesisState - all issue state that must be provided at genesis
type GenesisState struct {
	IssueConfigParams params.IssueConfigParams `json:"issue"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(ip params.IssueConfigParams) GenesisState {
	return GenesisState{IssueConfigParams: ip}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	minDepositTokens := sdk.TokensFromTendermintPower(10)
	return NewGenesisState(
		params.IssueConfigParams{MinDeposit: sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, minDepositTokens)}})
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	keeper.SetIssueConfigParams(ctx, data.IssueConfigParams)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	return NewGenesisState(keeper.GetIssueConfigParams(ctx))
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
