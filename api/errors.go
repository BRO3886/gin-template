package api

import (
	"errors"
)

//a list of api specific errors
var (
	ErrInvalidToken       = errors.New("Invalid token")
	ErrTokenMissing       = errors.New("Token not provided")
	ErrNotFound           = errors.New("Not found")
	ErrInvalidCredentials = errors.New("Invalid credentials provided")
	ErrNoContent          = errors.New("No content")
	ErrInvalidSlug        = errors.New("Invalid slug")
	ErrUnauthorised       = errors.New("You are not allowed to perform this action")
	ErrForbidden          = errors.New("Access to this resource is forbidden")
	ErrAlreadyExists      = errors.New("already exists")
)
