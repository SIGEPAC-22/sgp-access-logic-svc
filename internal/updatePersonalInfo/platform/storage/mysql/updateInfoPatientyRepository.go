package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-access-logic-svc/kit/constants"
)

type UpdatePersonalInfoRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdatePersonalInfoRepository(db *sql.DB, logger kitlog.Logger) *UpdatePersonalInfoRepository {
	return &UpdatePersonalInfoRepository{db: db, logger: logger}
}

func (u UpdatePersonalInfoRepository) UpdatePersonalInfoRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType int, documentNumber string, user string, password string, typeUser int) (bool, error) {

	sql, err := u.db.ExecContext(ctx, "UPDATE psi_personal_information SET psi_first_name = ?, psi_second_name = ?, psi_first_last_name = ?, psi_second_last_name = ?, psi_document_type_id = ?,\npsi_document_number = ?, psi_user = ?, psi_password = ?, psi_type_user = ? WHERE psi_id = ?;", firstName, secondName, lastFirstName, lastSecondName, documentType, documentNumber, user, password, typeUser, Id)
	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to update information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
