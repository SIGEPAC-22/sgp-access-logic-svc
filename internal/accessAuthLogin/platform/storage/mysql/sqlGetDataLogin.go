package mysql

type SqlGetDataLogin struct {
	User     string `db:"psi_user"`
	Password string `db:"psi_password"`
}
