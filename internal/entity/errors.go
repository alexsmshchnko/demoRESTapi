package entity

import "errors"

var (
	ErrEmptyFirstname = errors.New("validation error: empty Firstname")
	ErrEmptyLastname  = errors.New("validation error: empty Lastname")
	ErrEmptyEmail     = errors.New("validation error: Email is missing")
	ErrEmptyAge       = errors.New("validation error: Age is not set")
)
