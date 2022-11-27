package mysql

type sqlGetUpdateInPersonal struct {
	DocumentType int `db:"psi_document_type_id"`
	TypeUser     int `db:"psi_type_user"`
}
