package issue

import (
	"bytes"

	"github.com/hashgard/hashgard/x/issue/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
)

// GenesisState - all issue state that must be provided at genesis
type GenesisState struct {
	StartingIssueId uint64 `json:"starting_issue_id"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingIssueId uint64) GenesisState {
	return GenesisState{startingIssueId}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.CoinIssueMinId)
}

// Returns if a GenesisState is empty or has data in it
func (data GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return data.Equal(emptyGenState)
}

// Checks whether 2 GenesisState structs are equivalent.
func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := MsgCdc.MustMarshalBinaryBare(data)
	b2 := MsgCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	err := keeper.SetInitialIssueStartingIssueId(ctx, data.StartingIssueId)
	if err != nil {
		// TODO: Handle this with #870
		panic(err)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	startingIssueId, _ := keeper.PeekCurrentIssueID(ctx)
	return GenesisState{
		StartingIssueId: startingIssueId,
	}
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
