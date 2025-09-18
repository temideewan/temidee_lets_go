package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the "users" table.
func (u *UserModel) Insert(name, email, password string) error {
	// TODO: implement the insert
	return nil
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (u *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}


// exists will return a boolean if the user exists
func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
