package box

import (
	"bytes"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all box state that must be provided at genesis
type GenesisState struct {
	StartingLockBoxId    uint64    `json:"starting_lock_box_id"`
	StartingDepositBoxId uint64    `json:"starting_deposit_box_id"`
	StartingFutureBoxId  uint64    `json:"starting_future_box_id"`
	LockBoxs             []BoxInfo `json:"lock_boxs"`
	DepositBoxs          []BoxInfo `json:"deposit_boxs"`
	FutureBoxs           []BoxInfo `json:"future_boxs"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingLockBoxId uint64, startingDepositBoxId uint64, startingFutureBoxId uint64) GenesisState {
	return GenesisState{
		StartingLockBoxId:    startingLockBoxId,
		StartingDepositBoxId: startingDepositBoxId,
		StartingFutureBoxId:  startingFutureBoxId}
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
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Lock, data.StartingLockBoxId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Deposit, data.StartingDepositBoxId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingBoxId(ctx, types.Future, data.StartingFutureBoxId); err != nil {
		panic(err)
	}

	for _, box := range data.LockBoxs {
		keeper.AddBox(ctx, &box)
		if box.BoxStatus == types.LockBoxLocked {
			keeper.InsertActiveBoxQueue(ctx, box.Lock.EndTime, box.BoxId)
		}
	}
	for _, box := range data.DepositBoxs {
		keeper.AddBox(ctx, &box)
		switch box.BoxStatus {
		case types.BoxCreated:
			keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)
		case types.BoxDepositing:
			keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)
		case types.DepositBoxInterest:
			keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
		}
	}
	for _, box := range data.FutureBoxs {
		keeper.AddBox(ctx, &box)
		switch box.BoxStatus {
		case types.BoxDepositing:
			keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.GetFutureBoxSeqString(&box, 0))
		case types.BoxActived:
			times := len(box.Future.TimeLine)
			keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[times-1], keeper.GetFutureBoxSeqString(&box, times))
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
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
	genesisState.LockBoxs = keeper.ListAll(ctx, types.Lock)
	genesisState.DepositBoxs = keeper.ListAll(ctx, types.Deposit)
	genesisState.FutureBoxs = keeper.ListAll(ctx, types.Future)

	return genesisState
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
