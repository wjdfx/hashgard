package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
	issueerr "github.com/hashgard/hashgard/x/issue/errors"
)

//Process Future box

func (keeper Keeper) ProcessFutureBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	coinIssueInfo := keeper.GetIssueKeeper().GetIssue(ctx, box.TotalAmount.Token.Denom)
	if coinIssueInfo == nil {
		return issueerr.ErrUnknownIssue(box.Deposit.Interest.Token.Denom)
	}
	box.BoxStatus = types.BoxDepositing
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
	keeper.addAddressDeposit(ctx, box.BoxId, sender, types.NewBoxDeposit(deposit.Amount))
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
	boxDeposit := keeper.GetDepositByAddress(ctx, box.BoxId, sender)
	if boxDeposit.Amount.LT(deposit.Amount) {
		return errors.ErrNotEnoughAmount()
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.NewCoins(deposit), box.BoxId); err != nil {
		return err
	}
	for i, v := range box.Future.Deposits {
		if v.Address.Equals(sender) {
			box.Future.Deposits[i].Amount = box.Future.Deposits[i].Amount.Sub(deposit.Amount)
			break
		}
	}
	boxDeposit.Amount = boxDeposit.Amount.Sub(deposit.Amount)
	if boxDeposit.Amount.IsZero() {
		keeper.removeAddressDeposit(ctx, box.BoxId, sender)
	} else {
		keeper.setAddressDeposit(ctx, box.BoxId, sender, boxDeposit)
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
	for i, item := range box.Future.TimeLine {
		seq := i + 1
		keeper.InsertActiveBoxQueue(ctx, item, keeper.getFutureBoxSeqString(box, seq))
	}
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
func (keeper Keeper) processFutureBoxActiveByEndBlocker(ctx sdk.Context, box *types.BoxInfo, seq int) sdk.Error {
	if types.BoxActived != box.BoxStatus {
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if seq == 0 {
		return nil
	}
	boxDenom := utils.GetCoinDenomByFutureBoxSeq(box.BoxId, seq)
	for _, items := range box.Future.Receivers {
		address, _ := sdk.AccAddressFromBech32(items[0])
		amount, _ := sdk.NewIntFromString(items[seq])
		//fmt.Println(address.String() + ":" + boxDenom + ":" + amount.String())
		_, err := keeper.GetBankKeeper().SubtractCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(boxDenom, amount)))
		if err != nil {
			return err
		}
		if err := keeper.FetchDepositedCoin(ctx, address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.BoxId); err != nil {
			return err
		}
	}
	timeLine := box.Future.TimeLine[seq-1]
	if box.Future.Distributed == nil {
		box.Future.Distributed = []int64{timeLine}
	} else {
		box.Future.Distributed = append(box.Future.Distributed, timeLine)
	}
	if seq == len(box.Future.TimeLine) {
		box.BoxStatus = types.BoxFinished
	}
	keeper.RemoveFromActiveBoxQueue(ctx, timeLine, keeper.getFutureBoxSeqString(box, seq))
	keeper.setBox(ctx, box)
	return nil
}
