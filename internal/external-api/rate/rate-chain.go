package rate

import (
	apperrors "se-school-case/pkg/errors"
)

type CurrencyFetcher interface {
	Fetch() (float64, error)
	SetNext(fetcher CurrencyFetcher)
}

type DefaultCurrencyFetcher struct {
	next CurrencyFetcher
}

func (c *DefaultCurrencyFetcher) SetNext(fetcher CurrencyFetcher) {
	c.next = fetcher
}

func (c *DefaultCurrencyFetcher) Fetch() (float64, error) {
	if c.next == nil {
		return 0, apperrors.ErrRateFetch
	}
	return c.next.Fetch()
}
