package accessAuthLogin

import "context"

type Repository interface {
	GetDataAuthLoginRepo(ctx context.Context, user string) (GetDataAuthResponseRepo, error)
}

type Service interface {
	GetDataAuthLoginSvc(ctx context.Context, user, password string) (GetDataAuthLoginResponse, error)
}

type GetDataAuthLoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type GetDataAuthResponseRepo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type GetDataAuthLoginResponse struct {
	Status bool `json:"status"`
}
