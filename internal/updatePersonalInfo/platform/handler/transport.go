package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"sgp-access-logic-svc/kit/constants"
)

func NewUpdatePersonalInfoHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle(path,
		httptransport.NewServer(endpoints,
			DecodeRequestUpdatePersonalInfo,
			EncodeRequestUpdatePersonalInfo,
		)).Methods(http.MethodPut)
	return r
}

func DecodeRequestUpdatePersonalInfo(ctx context.Context, r *http.Request) (interface{}, error) {
	processID, _ := uuid.NewUUID()
	ctx = context.WithValue(ctx, constants.UUID, processID.String())
	id := r.URL.Query().Get("id")
	var confRequest UpdatePersonalInfoInternalRequest
	confRequest.Id = id
	err := json.NewDecoder(r.Body).Decode(&confRequest)
	confRequest.ctx = ctx
	return confRequest, err
}

func EncodeRequestUpdatePersonalInfo(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, _ := response.(UpdatePersonalInfoInternalResponse)
	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch resp.Err {
		case constants.ErrorDataError:
			w.WriteHeader(http.StatusBadRequest)
			break
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return json.NewEncoder(w).Encode(resp.Err.Error())
	}
	return json.NewEncoder(w).Encode(resp.Response)
}
