package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-access-logic-svc/internal/accessAuthLogin"
	"sgp-access-logic-svc/kit/constants"
)

type GetDataAuthLoginSvc struct {
	repoDB accessAuthLogin.Repository
	logger kitlog.Logger
}

func NewGetDataAuthLoginSvc(repoDB accessAuthLogin.Repository, logger kitlog.Logger) *GetDataAuthLoginSvc {
	return &GetDataAuthLoginSvc{repoDB: repoDB, logger: logger}
}

func (g *GetDataAuthLoginSvc) GetDataAuthLoginSvc(ctx context.Context, user, password string) (accessAuthLogin.GetDataAuthLoginResponse, error) {
	g.logger.Log("Starting Get access login", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetDataAuthLoginRepo(ctx, user)
	if err != nil {
		g.logger.Log("Error trying to obtained data login", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return accessAuthLogin.GetDataAuthLoginResponse{
			Status: false,
		}, err
	}

	if resp.User == user && resp.Password == password {
		return accessAuthLogin.GetDataAuthLoginResponse{
			Status: true,
		}, nil
	} else {
		return accessAuthLogin.GetDataAuthLoginResponse{
			Status: false,
		}, nil
	}

}
