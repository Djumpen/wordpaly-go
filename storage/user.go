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

// Create new user
func (s *Storage) CreateUser(user *User) (int64, error) {
	var createUserStmt = `INSERT INTO user ( 
		username,
		hash
	) VALUES (
		:username,
		:hash
	)`

	// user.CreatedAt = timeNowFunc()
	res, err := sqlx.NamedExec(s.db, createUserStmt, user)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Get user by ID
func (s *Storage) UserByID(id int64) (*User, error) {
	var selectUserByIDStmt = `SELECT * FROM user WHERE id = ?`

	var user User
	idconv := strconv.Itoa(int(id))
	err := sqlx.Select(s.db, &user, selectUserByIDStmt, idconv)
	return &user, err
}

// Get user by username
func (s *Storage) UserByUsername(username string) (*User, error) {
	var selectUserByUsernameStmt = `SELECT * FROM user WHERE username = ? LIMIT 1`

	var user User
	err := sqlx.Get(s.db, &user, selectUserByUsernameStmt, username)
	return user.conv(), err
}

// func (s *Storage) UserHash(username) (int64, error) {
// 	var stmt = `SELECT * FROM user WHERE username = ? AND hash = ?`

// 	var user *User
// 	err := sqlx.Select(s.db, user, stmt, username, hash)
// 	return user.ID, err
// }
