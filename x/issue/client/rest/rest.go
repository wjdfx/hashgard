package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
)

const (
	RestParamsType = "type"
	RestIssueID    = "issue-id"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, queryRoute string) {
	registerQueryRoutes(cliCtx, r, cdc, queryRoute)
	registerTxRoutes(cliCtx, r, cdc)
}
