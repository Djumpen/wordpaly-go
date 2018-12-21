package storage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

var timeNowFunc = func() uint {
	return uint(time.Now().Unix())
}
