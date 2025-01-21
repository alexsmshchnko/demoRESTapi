package entity

import (
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
		return ErrEmptyFirstname
	}
	if u.Lastname == "" {
		return ErrEmptyLastname
	}
	if u.Email == "" {
		return ErrEmptyEmail
	}
	if u.Age == 0 {
		return ErrEmptyAge
	}

	return nil
}
