package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-access-logic-svc/internal/updatePersonalInfo"
	"sgp-access-logic-svc/kit/constants"
	"strconv"
)

type UpdateInfoPatientService struct {
	repoDB updatePersonalInfo.Repository
	logger kitlog.Logger
}

func NewUpdateInfoPatientService(repoDB updatePersonalInfo.Repository, logger kitlog.Logger) *UpdateInfoPatientService {
	return &UpdateInfoPatientService{repoDB: repoDB, logger: logger}
}

func (u UpdateInfoPatientService) UpdatePersonalInfoSvc(ctx context.Context, Id string, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType string, documentNumber string, user string, password string, typeUser string) (updatePersonalInfo.UpdatePersonalInfoResponse, error) {
	u.logger.Log("Starting Update Info Patient", constants.UUID, ctx.Value(constants.UUID))

	idConverter, _ := strconv.Atoi(Id)

	selectRespDB, errSelect := u.repoDB.SelectInfoPersonalRepo(ctx, idConverter)
	if errSelect != nil {

	}

	if documentType == "" {
		documentType = strconv.FormatInt(int64(selectRespDB.DocumentType), 10)
	}

	if typeUser == "" {
		typeUser = strconv.FormatInt(int64(selectRespDB.TypeUser), 10)
	}

	idDocumentType, _ := strconv.Atoi(documentType)

	idTypeUser, _ := strconv.Atoi(typeUser)

	resp, err := u.repoDB.UpdatePersonalInfoRepo(ctx, idConverter, firstName, secondName, lastFirstName, lastSecondName, idDocumentType, documentNumber, user, password, idTypeUser)
	if err != nil {
		u.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updatePersonalInfo.UpdatePersonalInfoResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			u.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return updatePersonalInfo.UpdatePersonalInfoResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return updatePersonalInfo.UpdatePersonalInfoResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
