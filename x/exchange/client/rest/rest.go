package rest

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/hashgard/hashgard/x/exchange"
)

// nolint
const (
	RestOrderId		= "order-id"
	RestAddress		= "address"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/exchange/order/{%s}", RestOrderId), queryOrderHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/exchange/orders/{%s}", RestAddress), queryOrdersByAddrHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/exchange/frozen/{%s}", RestAddress), queryFrozenFundByAddrHandlerFn(cdc, cliCtx)).Methods("GET")

	r.HandleFunc("/exchange/order", postOrderHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/exchange/take/{%s}", RestOrderId), postTakeOrderHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/exchange/withdrawal/{%s}", RestOrderId), postWithdrawalOrderHandlerFn(cdc, cliCtx)).Methods("POST")
}

func queryOrderHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("orderId required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		orderId, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		params := exchange.NewQueryOrderParams(orderId)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryOrder), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func queryOrdersByAddrHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strAddress := vars[RestAddress]

		address, err := sdk.AccAddressFromBech32(strAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := exchange.NewQueryOrdersParams(address)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryAllOrdersByAddress), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func queryFrozenFundByAddrHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strAddress := vars[RestAddress]

		address, err := sdk.AccAddressFromBech32(strAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := exchange.NewQueryFrozenFundParams(address)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryFrozenFund), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

type PostOrderReq struct {
	BaseReq		rest.BaseReq	`json:"base_req"`
	Seller		sdk.AccAddress	`json:"seller"`
	Supply		sdk.Coin		`json:"supply"`
	Target		sdk.Coin		`json:"target"`
}

type TakeOrderReq struct {
	BaseReq		rest.BaseReq	`json:"base_req"`
	Amount		sdk.Coin		`json:"seller"`
	Buyer		sdk.AccAddress	`json:"buyer"`
}

type WithdrawalOrderReq struct {
	BaseReq		rest.BaseReq	`json:"base_req"`
	Seller		sdk.AccAddress	`json:"seller"`
}

func postOrderHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostOrderReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := exchange.NewMsgCreateOrder(req.Seller, req.Supply, req.Target)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
		}

		// derive the from account address and name from the Keybase
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// check
		if !bytes.Equal(fromAddress, req.Seller) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "seller address must be equal to from")
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}

func postTakeOrderHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("orderId required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		orderId, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		var req TakeOrderReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := exchange.NewMsgTakeOrder(orderId, req.Buyer, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
		}

		// derive the from account address and name from the Keybase
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// check
		if !bytes.Equal(fromAddress, req.Buyer) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "buyer address must be equal to from")
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}

func postWithdrawalOrderHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("orderId required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		orderId, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		var req WithdrawalOrderReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := exchange.NewMsgWithdrawalOrder(orderId, req.Seller)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		if req.BaseReq.GenerateOnly {
			clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
		}

		// derive the from account address and name from the Keybase
		fromAddress, _, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// check
		if !bytes.Equal(fromAddress, req.Seller) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "seller address must be equal to from")
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []sdk.Msg{msg}, cdc)
	}
}