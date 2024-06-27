package app_errors

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)
