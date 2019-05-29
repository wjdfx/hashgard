package box

import (
	"bytes"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all box state that must be provided at genesis
type GenesisState struct {
	StartingLockId       uint64    `json:"starting_lock_id"`
	StartingDepositId    uint64    `json:"starting_deposit_id"`
	StartingFutureId     uint64    `json:"starting_future_id"`
	LockBoxs             []BoxInfo `json:"lock_boxs"`
	DepositBoxs          []BoxInfo `json:"deposit_boxs"`
	FutureBoxs           []BoxInfo `json:"future_boxs"`
	LockBoxCreateFee     sdk.Coin  `json:"lock_box_create_fee"`
	DepositBoxCreateFee  sdk.Coin  `json:"deposit_box_create_fee"`
	FutureBoxCreateFee   sdk.Coin  `json:"future_box_create_fee"`
	BoxEnableTransferFee sdk.Coin  `json:"box_enable_transfer_fee"`
	BoxDescribeFee       sdk.Coin  `json:"box_describe_fee"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(startingLockId uint64, startingDepositId uint64, startingFutureId uint64) GenesisState {
	return GenesisState{
		StartingLockId:       startingLockId,
		StartingDepositId:    startingDepositId,
		StartingFutureId:     startingFutureId,
		LockBoxCreateFee:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntWithDecimal(100, 18)),
		DepositBoxCreateFee:  sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntWithDecimal(1000, 18)),
		FutureBoxCreateFee:   sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntWithDecimal(1000, 18)),
		BoxEnableTransferFee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntWithDecimal(1000, 18)),
		BoxDescribeFee:       sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntWithDecimal(100, 18))}
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
	if err := keeper.SetInitialBoxStartingId(ctx, types.Lock, data.StartingLockId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingId(ctx, types.Deposit, data.StartingDepositId); err != nil {
		panic(err)
	}
	if err := keeper.SetInitialBoxStartingId(ctx, types.Future, data.StartingFutureId); err != nil {
		panic(err)
	}

	keeper.SetLockBoxCreateFee(ctx, data.LockBoxCreateFee)
	keeper.SetDepositBoxCreateFee(ctx, data.DepositBoxCreateFee)
	keeper.SetFutureBoxCreateFee(ctx, data.FutureBoxCreateFee)
	keeper.SetEnableTransferFee(ctx, data.BoxEnableTransferFee)
	keeper.SetBoxDescribeFee(ctx, data.BoxDescribeFee)

	if data.LockBoxs != nil {
		for _, box := range data.LockBoxs {
			keeper.AddBox(ctx, &box)
			if box.Status == types.LockBoxLocked {
				keeper.InsertActiveBoxQueue(ctx, box.Lock.EndTime, box.Id)
			}
		}
	}

	if data.DepositBoxs != nil {
		for _, box := range data.DepositBoxs {
			keeper.AddBox(ctx, &box)
			switch box.Status {
			case types.BoxCreated:
				keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.Id)
			case types.BoxDepositing:
				keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.Id)
			case types.DepositBoxInterest:
				keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.Id)
			}
		}
	}

	if data.FutureBoxs != nil {
		for _, box := range data.FutureBoxs {
			keeper.AddBox(ctx, &box)
			switch box.Status {
			case types.BoxDepositing:
				keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[0], box.Id)
			case types.BoxActived:
				keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[len(box.Future.TimeLine)-1], box.Id)
			}
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	genesisState := GenesisState{}
	var err sdk.Error

	genesisState.StartingLockId, err = keeper.PeekCurrentBoxID(ctx, types.Lock)
	if err != nil {
		panic(err)
	}
	genesisState.StartingDepositId, err = keeper.PeekCurrentBoxID(ctx, types.Deposit)
	if err != nil {
		panic(err)
	}
	genesisState.StartingFutureId, err = keeper.PeekCurrentBoxID(ctx, types.Future)
	if err != nil {
		panic(err)
	}
	genesisState.LockBoxs = keeper.ListAll(ctx, types.Lock)
	genesisState.DepositBoxs = keeper.ListAll(ctx, types.Deposit)
	genesisState.FutureBoxs = keeper.ListAll(ctx, types.Future)

	genesisState.LockBoxCreateFee = keeper.GetLockBoxCreateFee(ctx)
	genesisState.DepositBoxCreateFee = keeper.GetDepositBoxCreateFee(ctx)
	genesisState.FutureBoxCreateFee = keeper.GetFutureBoxCreateFee(ctx)
	genesisState.BoxEnableTransferFee = keeper.GetEnableTransferFee(ctx)
	genesisState.BoxDescribeFee = keeper.GetBoxDescribeFee(ctx)

	return genesisState
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
