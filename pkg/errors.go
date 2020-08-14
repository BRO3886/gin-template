package pkg

import (
	"errors"
)

//a list of errors
var (
	ErrNotFound      = errors.New("Not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrDatabase      = errors.New("Database error")
	ErrEmail         = errors.New("Email not valid")
	ErrPassword      = errors.New("Password must be greater than 6 chars")
	ErrEmailExists   = errors.New("Email already exists")
)
