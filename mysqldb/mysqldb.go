package mysqldb

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//New creates new database connection for mysql database
func New(login, pwd, host, port, dbname string) *sqlx.DB {
	db := sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s",
		login,
		pwd,
		host,
		port,
		dbname,
	))
	db.MustExec("SET NAMES utf8")
	db.MustExec("SET CHARACTER SET utf8")
	db.MustExec("SET collation_connection = utf8_unicode_ci")
	return db
}
