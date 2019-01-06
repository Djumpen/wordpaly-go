package storage

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID         int64  `db:"id"`
	Username   string `db:"username"`
	Email      string
	Name       string `db:"name"`
	Hash       string `db:"hash"`
	CreatedAt  string `db:"created_at"`
	ModifiedAt string `db:"modified_at"`

	EmailNull sql.NullString `db:"email" json:"-"`
}

func (u *User) conv() *User {
	u.Email = u.EmailNull.String
	return u
}

type userStorage struct{}

func NewUserStorage() *userStorage {
	return &userStorage{}
}

// Create new user
func (s *userStorage) Create(tx *sql.Tx, user *User) (int64, error) {
	var createUserStmt = `INSERT INTO user ( 
		username,
		hash
	) VALUES (
		:username,
		:hash
	)`

	res, err := sqlx.NamedExec(txx(tx), createUserStmt, user)
	if err != nil {
		return 0, wrapError(err)
	}
	return res.LastInsertId()
}

// Get user by ID
func (s *userStorage) ByID(tx *sql.Tx, id int64) (*User, error) {
	var selectUserByIDStmt = `SELECT * FROM user WHERE id = ?`

	var user User
	idconv := strconv.Itoa(int(id))
	err := sqlx.Select(txx(tx), &user, selectUserByIDStmt, idconv)

	return &user, wrapError(err)
}

// Get user by username
func (s *userStorage) ByUsername(tx *sql.Tx, username string) (*User, error) {
	var selectUserByUsernameStmt = `SELECT * FROM user WHERE username = ? LIMIT 1`

	var user User
	err := sqlx.Get(txx(tx), &user, selectUserByUsernameStmt, username)
	if err != nil {
		return nil, wrapError(err)
	}
	return user.conv(), nil
}
