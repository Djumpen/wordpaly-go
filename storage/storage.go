package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	ErrDuplicate = "ERR_DUPLICATE"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

type StorageError struct {
	code string
	err  error
}

func NewErr(err error, code string) *StorageError {
	return &StorageError{
		code: code,
		err:  err,
	}
}

func CheckError(err error, code string) bool {
	if e, ok := err.(*StorageError); ok {
		if e.code == code {
			return true
		}
	}
	return false
}

func (e *StorageError) Error() string {
	return e.err.Error()
}

var timeNowFunc = func() uint {
	return uint(time.Now().Unix())
}
