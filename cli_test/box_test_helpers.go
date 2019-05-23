package clitest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/app"
	"github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/stretchr/testify/require"
)

//___________________________________________________________________________________
// hashgardcli box

// TxLockBoxCreate is hashgardcli box create-lock
//hashgardcli box create-lock foocoin 100coin174876e800 1557983880 --from joehe -y
func (f *Fixtures) TxLockBoxCreate(sender string, params *params.BoxLockParams, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box create-lock %s %s %d --from=%s %v", params.Name, params.TotalAmount.Token.String(),
		params.Lock.EndTime, sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositBoxCreate is hashgardcli box create-deposit
//hashgardcli box create-deposit foocoin 10000coin174876e800 --from joehe
// --bottom-line=1000 --price=100 --start-time=1557982140 --establish-time=1557982141  --maturity-time=1557982142 --interest=200coin174876e801 -y
func (f *Fixtures) TxDepositBoxCreate(sender string, params *params.BoxDepositParams, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box create-deposit %s %s "+
		"--bottom-line=%s "+
		"--price=%s --start-time=%d --establish-time=%d "+
		"--maturity-time=%d "+
		"--interest=%s "+
		"--from=%s %v", params.Name, params.TotalAmount.Token.String(),
		params.Deposit.BottomLine.String(), params.Deposit.Price.String(), params.Deposit.StartTime,
		params.Deposit.EstablishTime, params.Deposit.MaturityTime, params.Deposit.Interest.Token.String(), sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxFutureBoxCreate is hashgardcli box create-future
//hashgardcli box create-future joe 1800coin174876e800 /home/f.json -y --from joehe
func (f *Fixtures) TxFutureBoxCreate(sender string, params *params.BoxFutureParams, flags ...string) (bool, string, string) {
	json, _ := json.Marshal(params.Future)
	fileName := path.Join(f.GDHome, "future_data.json")
	if err := ioutil.WriteFile(fileName, json, os.ModeDir); err != nil {
		panic(err)
	}
	cmd := fmt.Sprintf("../build/hashgardcli box create-future %s %s %s "+
		"--from=%s %v", params.Name, params.TotalAmount.Token.String(), fileName, sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositBoxCreate is hashgardcli box interest-injection
//hashgardcli box interest-injection boxab3jlxpt2pt 200 --from joehe -y
func (f *Fixtures) TxDepositBoxInterestInjection(sender string, boxID string, amount sdk.Int, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box interest-injection %s %s "+
		"--from=%s %v", boxID, amount.String(), sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositBoxCreate is hashgardcli box interest-fetch
//hashgardcli box interest-fetch boxab3jlxpt2pt 200 --from joehe -y
func (f *Fixtures) TxDepositBoxInterestFetch(sender string, boxID string, amount sdk.Int, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box interest-fetch %s %s "+
		"--from=%s %v", boxID, amount.String(), sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositTo is hashgardcli box deposit-to
//hashgardcli box deposit-to boxab3jlxpt2pt 1000 --from test -y
func (f *Fixtures) TxDepositTo(sender string, boxID string, amount sdk.Int, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box deposit-to %s %s "+
		"--from=%s %v", boxID, amount.String(), sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxDepositTo is hashgardcli box deposit-fetch
//hashgardcli box deposit-fetch boxab3jlxpt2pt 1000 --from test -y
func (f *Fixtures) TxDepositFetch(sender string, boxID string, amount sdk.Int, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box deposit-fetch %s %s "+
		"--from=%s %v", boxID, amount.String(), sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxWithdraw is hashgardcli box withdraw
//hashgardcli box withdraw boxab3jlxpt2pt --from test -y
func (f *Fixtures) TxWithdraw(sender string, boxID string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli box withdraw %s "+
		"--from=%s %v", boxID, sender, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// QueryBox is hashgardcli box query-box
//hashgardcli box query-box boxab3jlxpt2ps
func (f *Fixtures) QueryFutureBox(boxID string, flags ...string) utils.FutureBoxInfo {
	cmd := fmt.Sprintf("../build/hashgardcli box query-box %s %v", boxID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var boxInfo utils.FutureBoxInfo
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &boxInfo)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return boxInfo
}
