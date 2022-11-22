package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/log"
	"gopkg.in/validator.v2"
	"sgp-access-logic-svc/kit/constants"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func UpdateInfoPatientTransportMiddleware(log kitlog.Logger) Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(UpdatePersonalInfoInternalRequest)
			if err := validator.Validate(&req); err != nil {
				log.Log("invalid request", "error", err.Error(), "request", req)
				return UpdatePersonalInfoInternalResponse{
					Response: constants.ErrorDataError.Error() + " - " + err.Error(),
					Err:      constants.ErrorDataError,
				}, nil
			}
			defer log.Log("process finished", "request", req)
			return e(ctx, req)
		}
	}
}
