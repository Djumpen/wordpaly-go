package services

import (
	"database/sql"

	"github.com/djumpen/wordplay-go/storage"
)

type usersStorage interface {
	Create(*sql.Tx, *storage.User) (int64, error)
	ByUsername(tx *sql.Tx, username string) (*storage.User, error)
}

type usersSerice struct {
	st usersStorage
	db *sql.DB
}

func NewUsersService(st usersStorage, db *sql.DB) *usersSerice {
	return &usersSerice{
		st: st,
		db: db,
	}
}

func (s *usersSerice) Create(u *storage.User) (id int64, err error) {
	withTransaction(s.db, func(tx *sql.Tx) error {
		id, err = s.st.Create(tx, u)
		return err
	})
	return
}

func (s *usersSerice) ByUsername(u string) (user *storage.User, err error) {
	withTransaction(s.db, func(tx *sql.Tx) error {
		user, err = s.st.ByUsername(tx, u)
		return err
	})
	return
}
