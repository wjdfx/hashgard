package keeper

import (
	"github.com/hashgard/hashgard/x/box/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process deposit box

func (keeper Keeper) ProcessDepositBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	decimal, err := keeper.GetCoinDecimals(ctx, box.Deposit.Interest.Token)
	if err != nil {
		return err
	}
	if box.Deposit.Interest.Decimals != decimal {
		return errors.ErrDecimalsNotValid(box.Deposit.Interest.Decimals)
	}
	box.BoxStatus = types.BoxCreated
	box.Deposit.TotalDeposit = sdk.ZeroInt()
	box.Deposit.WithdrawalInterest = sdk.ZeroInt()
	box.Deposit.WithdrawalShare = sdk.ZeroInt()
	//box.Deposit.TotalInterestInjection = sdk.ZeroInt()
	box.Deposit.Share = sdk.ZeroInt()
	keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)
	return nil
}
func (keeper Keeper) ProcessDepositBoxInterest(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin, operation string) (*types.BoxInfo, sdk.Error) {
	switch operation {
	case types.Injection:
		return keeper.injectionDepositBoxInterest(ctx, boxID, sender, interest)
	case types.Fetch:
		return keeper.fetchInterestFromDepositBox(ctx, boxID, sender, interest)

	}
	return nil, errors.ErrUnknownOperation()
}
func (keeper Keeper) injectionDepositBoxInterest(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.BoxCreated != box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if len(box.Deposit.InterestInjections) >= types.BoxMaxInjectionInterest {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjections == nil {
		box.Deposit.InterestInjections = []types.AddressDeposit{{Address: sender, Amount: interest.Amount}}
	} else {
		exist := false
		for i, v := range box.Deposit.InterestInjections {
			totalInterest = totalInterest.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Add(interest.Amount)
				exist = true
			}
		}
		if !exist {
			box.Deposit.InterestInjections = append(box.Deposit.InterestInjections, types.NewAddressDeposit(sender, interest.Amount))
		}
	}
	totalInterest = totalInterest.Add(interest.Amount)
	if totalInterest.GT(box.Deposit.Interest.Token.Amount) {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{interest}, boxID); err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) fetchInterestFromDepositBox(ctx sdk.Context, boxID string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return nil, errors.ErrUnknownBox(boxID)
	}
	if types.BoxCreated != box.BoxStatus {
		return nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectionNotValid(interest)
	}
	if box.Deposit.InterestInjections == nil {
		return nil, errors.ErrInterestFetchNotValid(interest)
	} else {
		remove := -1
		for i, v := range box.Deposit.InterestInjections {
			if v.Address.Equals(sender) {
				if box.Deposit.InterestInjections[i].Amount.LT(interest.Amount) {
					return nil, errors.ErrNotEnoughAmount()
				}
				box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Sub(interest.Amount)
				if box.Deposit.InterestInjections[i].Amount.IsZero() {
					remove = i
				}
			}
		}
		if remove != -1 {
			box.Deposit.InterestInjections = append(box.Deposit.InterestInjections[:remove], box.Deposit.InterestInjections[remove+1:]...)
		}
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.Coins{interest}, boxID); err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) processDepositBoxDeposit(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin, operation string) sdk.Error {
	switch operation {
	case types.DepositTo:
		return keeper.depositToDepositBox(ctx, box, sender, deposit)
	case types.Fetch:
		return keeper.fetchDepositFromDepositBox(ctx, box, sender, deposit)
	}
	return errors.ErrUnknownOperation()
}
func (keeper Keeper) depositToDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {
	if !deposit.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	if box.TotalAmount.Token.Denom != deposit.Denom {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	box.Deposit.TotalDeposit = box.Deposit.TotalDeposit.Add(deposit.Amount)
	if box.Deposit.TotalDeposit.GT(box.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{deposit}, box.BoxId); err != nil {
		return err
	}
	share := deposit.Amount.Quo(box.Deposit.Price)
	_, err := keeper.ck.AddCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.BoxId, share)))
	if err != nil {
		return err
	}
	box.Deposit.Share = box.Deposit.Share.Add(share)
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) fetchDepositFromDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, deposit sdk.Coin) sdk.Error {
	if !deposit.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(deposit.Amount.String())
	}
	if box.TotalAmount.Token.Denom != deposit.Denom {
		return errors.ErrAmountNotValid(deposit.Denom)
	}
	share := deposit.Amount.Quo(box.Deposit.Price)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.BoxId, share)))
	if err != nil {
		return err
	}
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.NewCoins(deposit), box.BoxId); err != nil {
		return err
	}
	box.Deposit.Share = box.Deposit.Share.Sub(share)
	if box.Deposit.Share.IsZero() {
		keeper.RemoveBox(ctx, box)
	} else {
		keeper.setBox(ctx, box)
	}
	return nil
}
func (keeper Keeper) ProcessDepositBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	switch box.BoxStatus {
	case types.BoxCreated:
		return keeper.processBoxCreatedByEndBlocker(ctx, box)
	case types.BoxDepositing:
		return keeper.processDepositBoxDepositToByEndBlocker(ctx, box)
	case types.DepositBoxInterest:
		return keeper.processDepositBoxInterestByEndBlocker(ctx, box)
	default:
		return errors.ErrNotAllowedOperation(box.BoxStatus)
	}
}
func (keeper Keeper) backBoxInterestInjections(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjections {
		if err := keeper.FetchDepositedCoin(ctx, v.Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.BoxId); err != nil {
			return err
		}
	}
	return nil
}
func (keeper Keeper) backBoxUnUsedInterestInjections(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	totalCoupon := box.TotalAmount.Token.Amount.Quo(box.Deposit.Price)
	if totalCoupon.Equal(box.Deposit.Share) {
		return nil
	}
	unused := utils.CalcInterest(box.Deposit.PerCoupon, totalCoupon.Sub(box.Deposit.Share), box.Deposit.Interest)
	interestInjectionsLen := len(box.Deposit.InterestInjections)
	if interestInjectionsLen == 0 {
		if err := keeper.FetchDepositedCoin(ctx, box.Deposit.InterestInjections[0].Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, unused)), box.BoxId); err != nil {
			return err
		}
		box.Deposit.InterestInjections[0].Amount = box.Deposit.InterestInjections[0].Amount.Sub(unused)
	} else {
		total := sdk.ZeroInt()
		for i, v := range box.Deposit.InterestInjections {
			var amount sdk.Int
			if i == interestInjectionsLen-1 {
				amount = unused.Sub(total)
			} else {
				amount = sdk.NewDecFromInt(v.Amount).QuoInt(box.Deposit.Interest.Token.Amount).MulInt(unused).TruncateInt()
			}
			if err := keeper.FetchDepositedCoin(ctx, v.Address,
				sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, amount)), box.BoxId); err != nil {
				return err
			}
			box.Deposit.InterestInjections[i].Amount = box.Deposit.InterestInjections[i].Amount.Sub(amount)
			total = total.Add(amount)
		}
	}
	box.Deposit.Interest.Token.Amount = box.Deposit.Interest.Token.Amount.Sub(unused)
	return nil
}
func (keeper Keeper) backBoxAllDeposit(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjections {
		if err := keeper.FetchDepositedCoin(ctx, v.Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.BoxId); err != nil {
			return err
		}
	}
	return nil
}
func (keeper Keeper) processBoxCreatedByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.BoxCreated {
		return nil
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjections != nil {
		for _, v := range box.Deposit.InterestInjections {
			totalInterest = totalInterest.Add(v.Amount)
		}
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.StartTime, box.BoxId)
	if box.Deposit.Interest.Token.Amount.Equal(totalInterest) {
		box.BoxStatus = types.BoxDepositing
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)
		keeper.setBox(ctx, box)
	} else {
		box.BoxStatus = types.BoxClosed
		if err := keeper.backBoxInterestInjections(ctx, box); err != nil {
			return err
		}
		keeper.RemoveBox(ctx, box)
	}
	return nil
}
func (keeper Keeper) processDepositBoxDepositToByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.BoxDepositing {
		return nil
	}

	box.BoxStatus = types.BoxClosed
	if box.Deposit.TotalDeposit.GTE(box.Deposit.BottomLine) {
		box.BoxStatus = types.DepositBoxInterest
	}

	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.BoxId)

	if box.BoxStatus == types.BoxClosed {
		if err := keeper.backBoxInterestInjections(ctx, box); err != nil {
			return err
		}
		//keeper.RemoveBox(ctx, box)
	} else {
		if err := keeper.backBoxUnUsedInterestInjections(ctx, box); err != nil {
			return err
		}
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
		//keeper.setBox(ctx, box)
	}
	if box.Deposit.TotalDeposit.IsZero() {
		keeper.RemoveBox(ctx, box)
	} else {
		keeper.setBox(ctx, box)
	}

	return nil
}
func (keeper Keeper) processDepositBoxWithdraw(ctx sdk.Context, boxID string, sender sdk.AccAddress) (sdk.Int, *types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, boxID)
	if box == nil {
		return sdk.ZeroInt(), nil, errors.ErrUnknownBox(boxID)
	}
	if types.Deposit != box.BoxType {
		return sdk.ZeroInt(), nil, errors.ErrNotSupportOperation()
	}
	if types.BoxFinished != box.BoxStatus {
		return sdk.ZeroInt(), nil, errors.ErrNotAllowedOperation(box.BoxStatus)
	}
	share := keeper.GetBankKeeper().GetCoins(ctx, sender).AmountOf(boxID)
	box.Deposit.WithdrawalShare = share.Add(box.Deposit.WithdrawalShare)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(boxID, share)))
	if err != nil {
		return sdk.ZeroInt(), nil, err
	}

	amount := share.Mul(box.Deposit.Price)
	if err := keeper.FetchDepositedCoin(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.BoxId); err != nil {
		return sdk.ZeroInt(), nil, err
	}

	interest := utils.CalcInterest(box.Deposit.PerCoupon, share, box.Deposit.Interest)

	if box.Deposit.WithdrawalShare == box.Deposit.Share {
		totalInterest := sdk.ZeroInt()
		for _, v := range box.Deposit.InterestInjections {
			totalInterest = totalInterest.Add(v.Amount)
		}
		interest = totalInterest.Sub(box.Deposit.WithdrawalInterest)
	}

	if err = keeper.FetchDepositedCoin(ctx, sender,
		sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, interest)), box.BoxId); err != nil {
		return sdk.ZeroInt(), nil, err
	}

	box.Deposit.WithdrawalInterest = interest.Add(box.Deposit.WithdrawalInterest)
	keeper.setBox(ctx, box)
	return interest, box, nil

}
func (keeper Keeper) processDepositBoxInterestByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.BoxStatus != types.DepositBoxInterest {
		return nil
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.BoxId)
	box.BoxStatus = types.BoxFinished
	keeper.setBox(ctx, box)
	return nil
}
