package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-logic-svc/internal/updatePersonalInfo"
)

func MakeUpdatePersonalInfoEndpoint(u updatePersonalInfo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePersonalInfoInternalRequest)
		resp, err := u.UpdatePersonalInfoSvc(req.ctx, req.Id, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DocumentType, req.DocumentNumber, req.User, req.Password, req.TypeUser)
		return UpdatePersonalInfoInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdatePersonalInfoInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdatePersonalInfoInternalRequest struct {
	Id             string `json:"id"`
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	LastFirstName  string `json:"lastFirstName"`
	LastSecondName string `json:"lastSecondName"`
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	User           string `json:"user"`
	Password       string `json:"password"`
	TypeUser       string `json:"typeUser"`
	ctx            context.Context
}
