package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/hashgard/hashgard/x/issue/domain"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

type PostIssueReq struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	Name            string       `json:"name"`
	TotalSupply     sdk.Int      `json:"total_supply"`
	MintingFinished bool         `json:"minting_finished"`
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/issue/issues", postProposalHandlerFn(cdc, cliCtx)).Methods("POST")
}
func postProposalHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostIssueReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			return
		}
		coinIssueInfo := domain.CoinIssueInfo{
			Issuer:          fromAddress,
			Name:            req.Name,
			TotalSupply:     req.TotalSupply,
			Decimals:        domain.DefaultDecimals,
			MintingFinished: req.MintingFinished,
		}
		// create the message
		msg := msgs.NewMsgIssue(&coinIssueInfo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
			return
		}
		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}
