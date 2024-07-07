package errors

import "errors"

var (
	ErrRateFetch = errors.New("no rate fetchers can fetch the exchange rate")
)
