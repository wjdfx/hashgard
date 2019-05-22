package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Process Future box

func (keeper Keeper) ProcessFutureBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	box.BoxStatus = types.BoxDepositing
	box.Future.TotalWithdrawal = sdk.ZeroInt()
	keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.getFutureBoxSeqString(box, 0))
	return nil
}
func (keeper Keeper) processFutureBoxDeposit(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin, operation string) sdk.Error {
	switch operation {
	case types.DepositTo:
		return keeper.depositToFutureBox(ctx, box, sender, deposit)
	case types.Fetch:
		return keeper.fetchDepositFromFutureBox(ctx, box, sender, deposit)
	default:
		return errors.ErrUnknownOperation()
	}
}
func (keeper Keeper) depositToFutureBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {
	if box.Future.TimeLine[0] < time.Now().Unix() {
		return errors.ErrNotSupportOperation()
	}
	totalDeposit := sdk.ZeroInt()
	if box.Future.Deposits == nil {
		box.Future.Deposits = []types.AddressDeposit{{Address: sender, Amount: deposit.Amount}}
	} else {
		exist := false
		for i, v := range box.Future.Deposits {
			totalDeposit = totalDeposit.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Future.Deposits[i].Amount = box.Future.Deposits[i].Amount.Add(deposit.Amount)
				exist = true
			}
		}
		if !exist {
			box.Future.Deposits = append(box.Future.Deposits, types.NewAddressDeposit(sender, deposit.Amount))
		}
	}
	totalDeposit = totalDeposit.Add(deposit.Amount)
	if totalDeposit.GT(box.TotalAmount.Token.Amount) {
		return errors.ErrNotEnoughAmount()
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{deposit}, box.BoxId); err != nil {
		return err
	}
	if totalDeposit.Equal(box.TotalAmount.Token.Amount) {
		if err := keeper.processFutureBoxDistribute(ctx, box); err != nil {
			return err
		}
		box.BoxStatus = types.BoxActived
		keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.getFutureBoxSeqString(box, 0))
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) fetchDepositFromFutureBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {
	if box.BoxStatus == types.BoxActived {
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Future.Deposits == nil {
		return errors.ErrNotEnoughAmount()
	}
	exist := false
	for i, v := range box.Future.Deposits {
		if v.Address.Equals(sender) {
			if box.Future.Deposits[i].Amount.LT(deposit.Amount) {
				return errors.ErrNotEnoughAmount()
			}
			box.Future.Deposits[i].Amount = box.Future.Deposits[i].Amount.Sub(deposit.Amount)
			exist = true
			break
		}
	}
	if !exist {
		return errors.ErrNotEnoughAmount()
	}

	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.NewCoins(deposit), box.BoxId); err != nil {
		return err
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) processFutureBoxDistribute(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	var address sdk.AccAddress
	var total = sdk.ZeroInt()
	for _, items := range box.Future.Receivers {
		for j, rec := range items {
			if j == 0 {
				addr, err := sdk.AccAddressFromBech32(rec)
				if err != nil {
					return sdk.ErrInvalidAddress(rec)
				}
				address = addr
				continue
			}
			amount, ok := sdk.NewIntFromString(rec)
			if !ok {
				return errors.ErrAmountNotValid(rec)
			}
			boxDenom := utils.GetCoinDenomByFutureBoxSeq(box.BoxId, j)
			_, err := keeper.GetBankKeeper().AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(boxDenom, amount)))
			if err != nil {
				return err
			}
			total = total.Add(amount)
		}
	}
	if !total.Equal(box.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid("Receivers")
	}
	times := len(box.Future.TimeLine)
	keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[times-1], keeper.getFutureBoxSeqString(box, times))
	return nil
}
func (keeper Keeper) getFutureBoxSeqString(box *types.BoxInfo, seq int) string {
	return fmt.Sprintf("%s:%d", box.BoxId, seq)
}

func (keeper Keeper) ProcessFutureBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo, seq int) sdk.Error {
	switch box.BoxStatus {
	case types.BoxDepositing:
		return keeper.processFutureBoxDepositToByEndBlocker(ctx, box, seq)
	case types.BoxActived:
		return keeper.processFutureBoxActiveByEndBlocker(ctx, box, seq)
	default:
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
}
func (keeper Keeper) processFutureBoxDepositToByEndBlocker(ctx sdk.Context, box *types.BoxInfo, seq int) sdk.Error {
	if types.BoxDepositing != box.BoxStatus {
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Future.Deposits != nil {
		for _, v := range box.Future.Deposits {
			if err := keeper.FetchDepositedCoin(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, v.Amount)), box.BoxId); err != nil {
				return err
			}
		}
	}
	//box.BoxStatus=types.BoxClosed
	keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.getFutureBoxSeqString(box, seq))
	keeper.RemoveBox(ctx, box)
	return nil
}

func (keeper Keeper) processFutureBoxWithdraw(ctx sdk.Context, boxIDSeq string, sender sdk.AccAddress) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxIDSeq)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxIDSeq)
	}
	if types.Future != box.BoxType {
		return nil, errors.ErrNotSupportOperation()
	}
	if types.BoxCreated == box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	seq := utils.GetSeqFromFutureBoxSeq(boxIDSeq)
	if box.Future.TimeLine[seq-1] > time.Now().Unix() {
		return nil, errors.ErrNotAllowedOperation(types.BoxUndue)
	}
	amount := keeper.GetBankKeeper().GetCoins(ctx, sender).AmountOf(boxIDSeq)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(boxIDSeq, amount)))
	if err != nil {
		return nil, err
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.BoxId); err != nil {
		return nil, err
	}
	box.Future.TotalWithdrawal = amount.Add(box.Future.TotalWithdrawal)
	keeper.setBox(ctx, box)
	return box, nil
}

func (keeper Keeper) processFutureBoxActiveByEndBlocker(ctx sdk.Context, box *types.BoxInfo, seq int) sdk.Error {
	if types.BoxActived != box.BoxStatus {
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if seq == 0 {
		return nil
	}
	box.BoxStatus = types.BoxFinished
	keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[seq-1], keeper.getFutureBoxSeqString(box, seq))
	keeper.setBox(ctx, box)
	return nil
}
