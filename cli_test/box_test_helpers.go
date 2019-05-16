package clitest

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/app"
	"github.com/hashgard/hashgard/x/box/params"
)

//___________________________________________________________________________________
// hashgardcli box

// TxLockBoxCreate is hashgardcli box create-lock
//hashgardcli box create-lock foocoin 100coin174876e800 1557983880 --from joehe -y
func (f *Fixtures) TxLockBoxCreate(params *params.BoxLockParams, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box create-lock %s %s %d --from=%s %v", params.Name, params.TotalAmount.Token.String(),
		params.Lock.EndTime, params.Sender.String(), f.Flags())
	fmt.Println(cmd)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositBoxCreate is hashgardcli box create-deposit
//hashgardcli box create-deposit foocoin 10000coin174876e800 --from joehe
// --bottom-line=1000 --price=100 --start-time=1557982140 --establish-time=1557982141  --maturity-time=1557982142 --interest=200coin174876e801 -y
func (f *Fixtures) TxDepositBoxCreate(params *params.BoxDepositParams, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box create-deposit %s %s "+
		"--bottom-line=%s "+
		"--price=%s --start-time=%d --establish-time=%d "+
		"--maturity-time=%d "+
		"--interest=%s "+
		"--from=%s %v", params.Name, params.TotalAmount.Token.String(),
		params.Deposit.BottomLine.String(), params.Deposit.Price.String(), params.Deposit.StartTime,
		params.Deposit.EstablishTime, params.Deposit.MaturityTime, params.Deposit.Interest.Token.String(), params.Sender.String(), f.Flags())
	fmt.Println(cmd)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositBoxCreate is hashgardcli box interest-injection
//hashgardcli box interest-injection boxab3jlxpt2pt 200 --from joehe -y
func (f *Fixtures) TxDepositBoxInterestInjection(sender string, boxID string, amount sdk.Int, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box interest-injection %s %s "+
		"--from=%s %v", boxID, amount.String(), sender, f.Flags())
	fmt.Println(cmd)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}
