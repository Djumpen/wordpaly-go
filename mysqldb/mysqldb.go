package mysqldb

import (
	"fmt"

	"github.com/djumpen/wordplay-go/cfg"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//New creates new database connection for mysql database
func New(creds cfg.DB) *sqlx.DB {
	db := sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=5s",
		creds.User,
		creds.Password,
		creds.Host,
		creds.Port,
		creds.DB,
	))
	db.MustExec("SET NAMES utf8")
	db.MustExec("SET CHARACTER SET utf8")
	db.MustExec("SET collation_connection = utf8_unicode_ci")
	return db
}
