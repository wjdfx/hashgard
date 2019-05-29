package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
)

type PostBoxBaseReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}
type PostDescriptionReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Description string       `json:"description"`
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Lock), postLockBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Deposit), postDepositBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/%s/create", types.Future), postFutureBoxCreateHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/deposit-to/{%s}/{%s}", BoxID, Amount), postDepositHandlerFn(cdc, cliCtx, types.DepositTo)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/deposit-fetch/{%s}/{%s}", BoxID, Amount), postDepositHandlerFn(cdc, cliCtx, types.Fetch)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/interest/injection/{%s}/{%s}", BoxID, Amount), postInterestHandlerFn(cdc, cliCtx, types.Injection)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/interest/fetch/{%s}/{%s}", BoxID, Amount), postInterestHandlerFn(cdc, cliCtx, types.Fetch)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/withdraw/{%s}", BoxID), postWithdrawHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/disable-feature/{%s}/{%s}", BoxID, Feature), postDisableFeatureHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/box/describe/{%s}", BoxID), postDescribeHandlerFn(cdc, cliCtx)).Methods("POST")

	//r.HandleFunc(fmt.Sprintf("/box/approve/{%s}/{%s}/{%s}", BoxID, AccAddress, Amount), postBoxApproveHandlerFn(cdc, cliCtx)).Methods("POST")
}
func postDisableFeatureHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[BoxID]
		if err := utils.CheckId(id); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		feature := vars[Feature]
		_, ok := types.Features[feature]
		if !ok {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrUnknownFeatures().Error())
			return
		}
		var req PostBoxBaseReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			return
		}
		account, err := cliCtx.GetAccount(fromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		boxInfo, err := clientutils.BoxOwnerCheck(cdc, cliCtx, account, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if feature == types.Transfer && boxInfo.GetBoxType() == types.Lock {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrNotSupportOperation().Error())
			return
		}
		msg := msgs.NewMsgBoxDisableFeature(id, fromAddress, feature)

		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}

}
func postDescribeHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars[BoxID]
		if err := boxutils.CheckId(id); err != nil {
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
		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			return
		}
		if len(req.Description) <= 0 || !json.Valid([]byte(req.Description)) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrBoxDescriptionNotValid().Error())
			return
		}

		msg := msgs.NewMsgBoxDescription(id, fromAddress, []byte(req.Description))
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		account, err := cliCtx.GetAccount(fromAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		_, err = clientutils.BoxOwnerCheck(cdc, cliCtx, account, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
