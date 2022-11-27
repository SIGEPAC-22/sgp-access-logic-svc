package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-access-logic-svc/internal/accessAuthLogin"
	"sgp-access-logic-svc/kit/constants"
)

type GetDataAuthLoginRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewGetDataAuthLoginRepository(db *sql.DB, logger kitlog.Logger) *GetDataAuthLoginRepository {
	return &GetDataAuthLoginRepository{db: db, logger: logger}
}

func (g *GetDataAuthLoginRepository) GetDataAuthLoginRepo(ctx context.Context, user string) (accessAuthLogin.GetDataAuthResponseRepo, error) {

	sql := g.db.QueryRowContext(ctx, "SELECT psi_user, psi_password FROM psi_personal_information where psi_user = ?;", user)
	g.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))

	var respDB SqlGetDataLogin

	if err := sql.Scan(&respDB.User, &respDB.Password); err != nil {
		g.logger.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
	}

	resp := accessAuthLogin.GetDataAuthResponseRepo{
		User:     respDB.User,
		Password: respDB.Password,
	}
	return resp, nil
}
