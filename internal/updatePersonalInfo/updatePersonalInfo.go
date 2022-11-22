package updatePersonalInfo

import (
	"context"
)

type Repository interface {
	UpdatePersonalInfoRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType int, documentNumber string, user string, password string, typeUser int) (bool, error)
}

type Service interface {
	UpdatePersonalInfoSvc(ctx context.Context, Id string, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType string, documentNumber string, user string, password string, typeUser string) (UpdatePersonalInfoResponse, error)
}

type SelectPersonalInfoResponse struct {
	DocumentType int `json:"documentType"`
	Department   int `json:"department"`
	Foreign      int `json:"foreign"`
	Pregnant     int `json:"pregnant"`
}

type UpdatePersonalInfoRequest struct {
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
}

type UpdatePersonalInfoResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
