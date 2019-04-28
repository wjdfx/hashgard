package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

func postApproveHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return approveHandlerFn(cdc, cliCtx, types.Approve)
}
func postIncreaseApproval(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return approveHandlerFn(cdc, cliCtx, types.IncreaseApproval)
}
func postDecreaseApproval(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return approveHandlerFn(cdc, cliCtx, types.DecreaseApproval)
}
func approveHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext, approveType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req PostIssueBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

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

		account, err := cliCtx.GetAccount(fromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var msg sdk.Msg

		if types.Approve == approveType {
			msg = msgs.NewMsgIssueApprove(issueID, account.GetAddress(), accAddress, amount)
		} else if types.IncreaseApproval == approveType {
			msg = msgs.NewMsgIssueIncreaseApproval(issueID, account.GetAddress(), accAddress, amount)
		} else {
			msg = msgs.NewMsgIssueDecreaseApproval(issueID, account.GetAddress(), accAddress, amount)
		}

		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
