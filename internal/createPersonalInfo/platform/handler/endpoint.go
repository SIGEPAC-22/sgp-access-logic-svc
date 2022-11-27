package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-access-logic-svc/internal/createPersonalInfo"
)

func MakeCreatePersonalInfoEndpoint(c createPersonalInfo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePersonalInfoInternalRequest)
		resp, err := c.CreatePersonalInfoSvc(req.ctx, req.FirstLastName, req.SecondName, req.FirstLastName, req.SecondLastName, req.Sex, req.DateBirth, req.DocumentType, req.DocumentNumber, req.User, req.Password, req.TypeUser)
		return CreatePersonalInfoInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type CreatePersonalInfoInternalResponse struct {
	Response interface{}
	Err      error
}

type CreatePersonalInfoInternalRequest struct {
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	FirstLastName  string `json:"firstLastName"`
	SecondLastName string `json:"secondLastName"`
	Sex            string `json:"sex"`
	DateBirth      string `json:"dateBirth"`
	DocumentType   string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	User           string `json:"user"`
	Password       string `json:"password"`
	TypeUser       string `json:"typeUser"`
	ctx            context.Context
}
