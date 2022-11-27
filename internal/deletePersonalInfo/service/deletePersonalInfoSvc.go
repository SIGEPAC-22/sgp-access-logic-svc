package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-access-logic-svc/internal/deletePersonalInfo"
	"sgp-access-logic-svc/kit/constants"
	"strconv"
)

type DeletePersonalInfoService struct {
	repoDB deletePersonalInfo.Repository
	logger kitlog.Logger
}

func NewDeletePersonalInfoService(repoDB deletePersonalInfo.Repository, logger kitlog.Logger) *DeletePersonalInfoService {
	return &DeletePersonalInfoService{repoDB: repoDB, logger: logger}
}

func (d DeletePersonalInfoService) DeletePersonalInfoSvc(ctx context.Context, Id string) (deletePersonalInfo.DeletePersonalInfoResponse, error) {
	d.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := d.repoDB.DeletePersonalInfoRepo(ctx, int64(IdConverter))
	if err != nil {
		d.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return deletePersonalInfo.DeletePersonalInfoResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			d.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return deletePersonalInfo.DeletePersonalInfoResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return deletePersonalInfo.DeletePersonalInfoResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
