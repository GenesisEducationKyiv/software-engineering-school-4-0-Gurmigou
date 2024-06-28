package app_errors

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrRateFetch          = errors.New("no rate fetchers can fetch the exchange rate")
)
