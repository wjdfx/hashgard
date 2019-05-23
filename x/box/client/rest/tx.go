package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
	"github.com/hashgard/hashgard/x/box/types"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Lock), postLockBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Deposit), postDepositBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Future), postFutureBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")

	//r.HandleFunc(fmt.Sprintf("/box/approve/{%s}/{%s}/{%s}", IssueID, AccAddress, Amount), postIssueApproveHandlerFn(cdc, cliCtx)).Methods("POST")
}
