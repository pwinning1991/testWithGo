package psql

import (
	"database/sql"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("psql: resource could note be located")
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserStore struct {
	sql interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}
}

func (us *UserStore) Find(id int) (*User, error) {
	const query = `SELECT id, name, email FROM users WHERE id=$1;`
	row := us.sql.QueryRow(query, id)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, errors.Wrap(err, "psql: error querying for user by id")
	}
}

// Create will create a new user in the DB using the provided user and
// will update the ID of the provided user. If there is an error it will
// be wrapped and returned.
func (us *UserStore) Create(user *User) error {
	const query = `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := us.sql.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return errors.Wrap(err, "psql: error creating new user")
	}
	return nil
}

// Delete will delete a user form the DB. If there is an error it will
// be wrapped and returned.
func (us *UserStore) Delete(id int) error {
	const query = `DELETE FROM users WHERE id=$1;`
	_, err := us.sql.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "psql: error deleting user")
	}
	return nil
}
