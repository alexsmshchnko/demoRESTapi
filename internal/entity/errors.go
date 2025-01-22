package entity

import (
	e "demorestapi/internal/common/errors"
)

var (
	ErrEmptyFirstname = e.NewIncorrectInputError("empty-firstname", "empty Firstname") // errors.New("validation error: empty Firstname")
	ErrEmptyLastname  = e.NewIncorrectInputError("empty-lastname", "empty Lastname")
	ErrEmptyEmail     = e.NewIncorrectInputError("email-is-missing", "Email is missing")
	ErrEmptyAge       = e.NewIncorrectInputError("age-is-not-set", "Age is not set")
)
