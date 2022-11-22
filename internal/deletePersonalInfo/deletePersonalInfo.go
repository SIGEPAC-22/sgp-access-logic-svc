package deletePersonalInfo

import "context"

type Repository interface {
	DeletePersonalInfoRepo(ctx context.Context, Id int64) (bool, error)
}

type Service interface {
	DeletePersonalInfoSvc(ctx context.Context, Id string) (DeletePersonalInfoResponse, error)
}

type DeletePersonalInfoRequest struct {
	Id string `json:"id"`
}

type DeletePersonalInfoResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
