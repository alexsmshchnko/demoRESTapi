package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

func NewUser() *User {
	return &User{}
}

func (u *User) Validate() error {
	if u.Firstname == "" {
		return errors.New("empty Firstname")
	}
	if u.Lastname == "" {
		return errors.New("empty Lasttname")
	}
	if u.Email == "" {
		return errors.New("empty email")
	}
	if u.Age == 0 {
		return errors.New("age is not set")
	}

	return nil
}
