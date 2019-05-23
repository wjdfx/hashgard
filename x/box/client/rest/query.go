package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/hashgard/hashgard/x/box/client/queriers"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/%s/{%s}", types.QuerierRoute, BoxID), queryBoxHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s/{%s}/{%s}", types.QuerierRoute, types.QuerySearch, BoxType, Name), queryBoxSearchHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/{%s}/%s", types.QuerierRoute, BoxType, types.QueryList), queryBoxsHandlerFn(cdc, cliCtx)).Methods("GET")
}
func queryBoxHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		boxID := vars[BoxID]
		if err := boxutils.CheckBoxId(boxID); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		res, err := queriers.QueryBoxByID(boxID, cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
func queryBoxSearchHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		res, err := queriers.QueryBoxByName(vars[BoxType], vars[Name], cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
func queryBoxsHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address, err := sdk.AccAddressFromBech32(r.URL.Query().Get(restAddress))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		boxQueryParams := params.BoxQueryParams{
			StartBoxId: r.URL.Query().Get(restStartBoxId),
			BoxType:    vars[BoxType],
			Owner:      address,
			Limit:      30,
		}
		strNumLimit := r.URL.Query().Get(restLimit)
		if len(strNumLimit) > 0 {
			limit, err := strconv.Atoi(r.URL.Query().Get(restLimit))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			boxQueryParams.Limit = limit
		}

		res, err := queriers.QueryBoxsList(boxQueryParams, cdc, cliCtx)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
