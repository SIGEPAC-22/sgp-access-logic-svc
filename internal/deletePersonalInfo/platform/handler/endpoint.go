package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-logic-svc/internal/deletePersonalInfo"
)

func MakeDeletePersonalInfoEndpoint(d deletePersonalInfo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePersonalInfoInternalRequest)
		resp, err := d.DeletePersonalInfoSvc(req.ctx, req.Id)
		return DeletePersonalInfoInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type DeletePersonalInfoInternalResponse struct {
	Response interface{}
	Err      error
}

type DeletePersonalInfoInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
