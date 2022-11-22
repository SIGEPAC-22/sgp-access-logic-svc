package createPersonalInfo

import (
	"context"
	"time"
)

type Repository interface {
	CreatePersonalInfoRepo(ctx context.Context, firstName string, secondName string, firstLastName string, secondLastName string, sex int, dateBirth string, documentType int, documentNumber string, user string, password string, typeUser int, date time.Time) (bool, error)
}

type Service interface {
	CreatePersonalInfoSvc(ctx context.Context, firstName, secondName, firstLastName, secondLastName, sex, dateBirth, documentType, documentNumber, user, password, typeUser string) (PersonalInfoResponse, error)
}

type PersonalInfoRequest struct {
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
}

type PersonalInfoResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
