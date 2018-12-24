package storage

import (
	"database/sql"
	"strconv"

	"github.com/djumpen/wordplay-go/mysqldb"
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

	res, err := sqlx.NamedExec(s.db, createUserStmt, user)
	if err != nil {
		if mysqldb.CheckError(err, mysqldb.ER_DUP_ENTRY) {
			return 0, NewErr(err, ErrDuplicate)
		}
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// Get user by username
func (s *Storage) UserByUsername(username string) (*User, error) {
	var selectUserByUsernameStmt = `SELECT * FROM user WHERE username = ? LIMIT 1`

	var user User
	err := sqlx.Get(s.db, &user, selectUserByUsernameStmt, username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user.conv(), err
}
