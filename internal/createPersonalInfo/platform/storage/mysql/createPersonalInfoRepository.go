package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-access-logic-svc/kit/constants"
	"time"
)

type CreatePersonalInfoRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewCreatePersonalInfoRepository(db *sql.DB, logger kitlog.Logger) *CreatePersonalInfoRepository {
	return &CreatePersonalInfoRepository{db: db, logger: logger}
}

func (c CreatePersonalInfoRepository) CreatePersonalInfoRepo(ctx context.Context, firstName string, secondName string, firstLastName string, secondLastName string, sex int, dateBirth string, documentType int, documentNumber string, user string, password string, typeUser int, date time.Time) (bool, error) {

	sql, err := c.db.ExecContext(ctx, "INSERT INTO psi_personal_information(psi_first_name,psi_second_name,psi_first_last_name,psi_second_last_name,psi_sex_id,psi_date_of_birth,\npsi_document_type_id,psi_document_number,psi_user,psi_password,psi_type_user,psi_account_creation_date,psi_data_of_last_use)VALUES\n(?,?,?,?,?,?,?,?,?,?,?,?,?);", firstName, secondName, firstLastName, secondLastName, sex, dateBirth, documentType, documentNumber, user, password, typeUser, date, date)
	c.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		c.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			c.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
