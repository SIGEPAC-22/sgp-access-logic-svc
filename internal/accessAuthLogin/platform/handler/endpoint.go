package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-logic-svc/internal/accessAuthLogin"
)

func MakeGetDataAuthLoginEndpoint(a accessAuthLogin.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDataAuthLoginInternalRequest)
		resp, err := a.GetDataAuthLoginSvc(req.ctx, req.User, req.Password)
		return GetDataAuthLoginInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetDataAuthLoginInternalResponse struct {
	Response interface{}
	Err      error
}

type GetDataAuthLoginInternalRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
	ctx      context.Context
}
