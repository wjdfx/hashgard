package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashgard/hashgard/x/issue/errors"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/params"
	"github.com/hashgard/hashgard/x/issue/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

type PostIssueReq struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	params.IssueParams `json:"issue"`
}
type PostDescriptionReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Description string       `json:"description"`
}
type PostIssueBaseReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/issue/create", postIssueHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/describe/{%s}", IssueID), postDescribeHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/mint/{%s}/{%s}/{%s}", IssueID, Amount, To), postMintHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn/{%s}/{%s}", IssueID, Amount), postBurnHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn-from/{%s}/{%s}/{%s}", IssueID, AccAddress, Amount), postBurnFromHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn-off/{%s}", IssueID), postBurnOffHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn-from-off/{%s}", IssueID), postBurnFromOffHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/burn-any-off/{%s}", IssueID), postBurnAnyOffHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/issue/finish-minting/{%s}", IssueID), postFinishMintingHandlerFn(cdc, cliCtx)).Methods("POST")
}
func postIssueHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
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
		if len(req.Description) > 0 && !json.Valid([]byte(req.Description)) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrCoinDescriptionNotValid().Error())
			return
		}
		coinIssueInfo := types.CoinIssueInfo{
			Owner:           fromAddress,
			Name:            req.Name,
			Symbol:          strings.ToUpper(req.Symbol),
			TotalSupply:     req.TotalSupply,
			Decimals:        req.Decimals,
			IssueTime:       time.Now(),
			Description:     req.Description,
			BurnOff:         req.BurnOff,
			BurnFromOff:     req.BurnFromOff,
			BurnAnyOff:      req.BurnAnyOff,
			MintingFinished: req.MintingFinished,
		}
		// create the message
		msg := msgs.CreateMsgIssue(&coinIssueInfo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postMintHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if err := issueutils.CheckIssueId(issueID); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		num, err := strconv.ParseInt(vars[Amount], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		amount := sdk.NewInt(num)
		to, err := sdk.AccAddressFromBech32(vars[To])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Query the issue
		res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var issueInfo types.Issue
		cdc.MustUnmarshalJSON(res, &issueInfo)
		msg := msgs.NewMsgIssueMint(issueID, fromAddress, amount, issueInfo.GetDecimals(), to)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
func postBurnHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if err := issueutils.CheckIssueId(issueID); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		amount, ok := sdk.NewIntFromString(vars[Amount])
		if !ok {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Amount not a valid int")
			return
		}

		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		account, err := cliCtx.GetAccount(fromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		amount, err = issueutils.BurnCheck(cdc, cliCtx, account, nil, issueID, amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := msgs.NewMsgIssueBurn(issueID, fromAddress, amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
func postBurnFromHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if err := issueutils.CheckIssueId(issueID); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		accAddress, err := sdk.AccAddressFromBech32(vars[AccAddress])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		amount, ok := sdk.NewIntFromString(vars[Amount])
		if !ok {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Amount not a valid int")
			return
		}

		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		account, err := cliCtx.GetAccount(fromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		amount, err = issueutils.BurnCheck(cdc, cliCtx, account, accAddress, issueID, amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := msgs.NewMsgIssueBurnFrom(issueID, fromAddress, accAddress, amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postFinishMintingHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIssueFlag(cdc, cliCtx, w, r, msgs.MsgIssueFinishMinting{})
	}
}
func postBurnOffHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIssueFlag(cdc, cliCtx, w, r, msgs.MsgIssueBurnOff{})
	}
}
func postBurnFromOffHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIssueFlag(cdc, cliCtx, w, r, msgs.MsgIssueBurnFromOff{})
	}
}
func postBurnAnyOffHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIssueFlag(cdc, cliCtx, w, r, msgs.MsgIssueBurnAnyOff{})
	}
}

func postIssueFlag(cdc *codec.Codec, cliCtx context.CLIContext, w http.ResponseWriter, r *http.Request, msg msgs.MsgFlag) {

	vars := mux.Vars(r)
	issueID := vars[IssueID]
	if err := issueutils.CheckIssueId(issueID); err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var req PostIssueBaseReq
	if !rest.ReadRESTReq(w, r, cdc, &req) {
		return
	}

	req.BaseReq = req.BaseReq.Sanitize()
	if !req.BaseReq.ValidateBasic(w) {
		return
	}
	fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//msg := msgs.NewMsgIssueBurnAnyOff(issueID, fromAddress)

	if err := msg.ValidateBasic(); err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	account, err := cliCtx.GetAccount(fromAddress)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = issueutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	switch msg.(type) {
	case msgs.MsgIssueBurnOff:
		msg = msgs.NewMsgIssueBurnOff(issueID, account.GetAddress())
	case msgs.MsgIssueBurnFromOff:
		msg = msgs.NewMsgIssueBurnFromOff(issueID, account.GetAddress())
	case msgs.MsgIssueBurnAnyOff:
		msg = msgs.NewMsgIssueBurnAnyOff(issueID, account.GetAddress())
	case msgs.MsgIssueFinishMinting:
		msg = msgs.NewMsgIssueFinishMinting(issueID, account.GetAddress())
	default:

	}

	clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})

}
func postDescribeHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		issueID := vars[IssueID]
		if err := issueutils.CheckIssueId(issueID); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		var req PostDescriptionReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if len(req.Description) <= 0 || !json.Valid([]byte(req.Description)) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrCoinDescriptionNotValid().Error())
			return
		}

		msg := msgs.NewMsgIssueDescription(issueID, fromAddress, []byte(req.Description))
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Query the issue
		_, err = issuequeriers.QueryIssueByID(issueID, cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
