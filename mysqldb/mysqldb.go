package mysqldb

import (
	"fmt"

	"github.com/djumpen/wordplay-go/config"
	"github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"
)

//New creates new database connection for mysql database
func New(creds config.DB) *sqlx.DB {
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

func CheckError(err error, code uint16) bool {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == code {
			return true
		}
	}
	return false
}
