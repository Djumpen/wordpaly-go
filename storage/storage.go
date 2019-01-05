package storage

import (
	"database/sql"

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

// func CheckError(err error, code string) bool {
// 	if e, ok := err.(*StorageError); ok {
// 		if e.code == code {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (e *StorageError) Error() string {
// 	return e.err.Error()
// }

type ErrNoRows string

func (e ErrNoRows) Error() string {
	return string(e)
}
func (e ErrNoRows) IsNoRows() bool {
	return true
}

func wrapErr(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return ErrNoRows(err.Error())
	}
	return err
	// TODO: add errrors
	// if mysqldb.CheckError(err, mysqldb.ER_DUP_ENTRY) {
	// 	return 0, NewErr(err, ErrDuplicate)
	// }
}

func txx(tx *sql.Tx) sqlx.Ext {
	txx := sqlx.Tx{Tx: tx, Mapper: reflectx.NewMapperFunc("db", sqlx.NameMapper)}
	return sqlx.Ext(&txx)
}
