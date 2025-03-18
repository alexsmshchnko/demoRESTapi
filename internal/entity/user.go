package entity

import (
	"errors"
	"time"
)

// User model info
// @Description User information
type User struct {
	ID        string    `json:"id" example:"d1e1a2ca-9e08-4fe6-8fd8-bc71e499cb63" format:"uuid"`
	Firstname string    `json:"firstname" example:"Doe"`
	Lastname  string    `json:"lastname" example:"John"`
	Email     string    `json:"email" example:"my@mail.com"`
	Age       uint      `json:"age" example:"25"`
	Created   time.Time `json:"created"`
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
