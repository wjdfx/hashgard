package box

import (
	"bytes"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/keeper"
)

// GenesisState - all box state that must be provided at genesis
type GenesisState struct {
	StartingLockBoxId    uint64 `json:"starting_lock_box_id"`
	StartingDepositBoxId uint64 `json:"starting_deposit_box_id"`
	StartingFutureBoxId  uint64 `json:"starting_future_box_id"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingLockBoxId uint64, startingDepositBoxId uint64, startingFutureBoxId uint64) GenesisState {
	return GenesisState{startingLockBoxId, startingDepositBoxId, startingFutureBoxId}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.BoxMinId, types.BoxMinId, types.BoxMinId)
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

	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Lock, data.StartingLockBoxId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Deposit, data.StartingDepositBoxId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Future, data.StartingFutureBoxId); err != nil {
		panic(err)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	genesisState := GenesisState{}
	var err sdk.Error

	genesisState.StartingLockBoxId, err = keeper.PeekCurrentBoxID(ctx, types.Lock)
	if err != nil {
		panic(err)
	}
	genesisState.StartingDepositBoxId, err = keeper.PeekCurrentBoxID(ctx, types.Deposit)
	if err != nil {
		panic(err)
	}
	genesisState.StartingFutureBoxId, err = keeper.PeekCurrentBoxID(ctx, types.Future)
	if err != nil {
		panic(err)
	}

	return genesisState
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
