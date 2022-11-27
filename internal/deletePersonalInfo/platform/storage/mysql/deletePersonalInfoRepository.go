package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-access-logic-svc/kit/constants"
)

type DeletePersonalInfoRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewDeletePersonalInfoRepository(db *sql.DB, logger kitlog.Logger) *DeletePersonalInfoRepository {
	return &DeletePersonalInfoRepository{db: db, logger: logger}
}

func (d DeletePersonalInfoRepository) DeletePersonalInfoRepo(ctx context.Context, Id int64) (bool, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	StatusInactive := config.GetInt("app-properties.getComorbidity.idStatusInactive")

	sql, err := d.db.ExecContext(ctx, "UPDATE psi_personal_information SET psi_state_data_id = ? WHERE psi_id = ?;", StatusInactive, Id)
	d.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		d.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			d.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
