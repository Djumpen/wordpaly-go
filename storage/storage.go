package storage

import (
	"database/sql"

	"github.com/djumpen/wordplay-go/apierrors"
	"github.com/djumpen/wordplay-go/mysqldb"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

// const (
// 	ErrDuplicate = "ERR_DUPLICATE"
// )

// type StorageError struct {
// 	code string
// 	err  error
// }

// func NewError(err error, code string) *StorageError {
// 	return &StorageError{
// 		code: code,
// 		err:  err,
// 	}
// }

// func (e *StorageError) Error() string {
// 	return e.err.Error()
// }

// func CheckError(err error, code string) bool {
// 	if e, ok := err.(*mysql.); ok {
// 		if e.code == code {
// 			return true
// 		}
// 	}
// 	return false
// }

// TODO: add errrors
func wrapError(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return apierrors.NewNoRows(err)
	}
	if mysqldb.CheckError(err, mysqldb.ER_DUP_ENTRY) {
		return apierrors.NewDuplicateEntry(err)
	}
	return err
}

func txx(tx *sql.Tx) sqlx.Ext {
	txx := sqlx.Tx{Tx: tx, Mapper: reflectx.NewMapperFunc("db", sqlx.NameMapper)}
	return sqlx.Ext(&txx)
}
