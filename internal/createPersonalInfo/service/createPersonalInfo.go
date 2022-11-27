package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-access-logic-svc/internal/createPersonalInfo"
	"sgp-access-logic-svc/kit/constants"
	"strconv"
	"time"
)

type PersonalInfoSvc struct {
	repoDB createPersonalInfo.Repository
	logger kitlog.Logger
}

func NewCreatePersonalInfoSvc(repoDB createPersonalInfo.Repository, logger kitlog.Logger) *PersonalInfoSvc {
	return &PersonalInfoSvc{repoDB: repoDB, logger: logger}
}

func (c PersonalInfoSvc) CreatePersonalInfoSvc(ctx context.Context, firstName, secondName, firstLastName, secondLastName, sex, dateBirth, documentType, documentNumber, user, password, typeUser string) (createPersonalInfo.PersonalInfoResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	sexConverter, _ := strconv.Atoi(sex)
	documentConverter, _ := strconv.Atoi(documentType)
	typeUserConverter, _ := strconv.Atoi(typeUser)

	timeNow := time.Now()

	resp, err := c.repoDB.CreatePersonalInfoRepo(ctx, firstName, secondName, firstLastName, secondLastName, sexConverter, dateBirth, documentConverter, documentNumber, user, password, typeUserConverter, timeNow)
	if err != nil {
		c.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createPersonalInfo.PersonalInfoResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			c.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return createPersonalInfo.PersonalInfoResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return createPersonalInfo.PersonalInfoResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
