package rate

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"se-school-case/pkg/util/constants"
)

type ExchangeApiRateFetch struct {
	DefaultCurrencyFetcher
}

func NewExchangeApiRateFetch() ExchangeApiRateFetch {
	return ExchangeApiRateFetch{}
}

// Provider: Exchange rate api service
func (s *ExchangeApiRateFetch) Fetch() (float64, error) {
	resp, err := http.Get(constants.EXCHANGE_API_URL)
	if err != nil {
		logrus.WithError(err).Error("error fetching exchange rates")
		return s.DefaultCurrencyFetcher.Fetch()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.WithError(err).Error("error closing response body")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("error reading response body")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	var exchangeRate ExchangeRateAPI
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		logrus.WithError(err).Error("error unmarshaling response")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	uahRate := exchangeRate.RateMap["uah"]
	if uahRate == 0 {
		logrus.Error("rates not found")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	return uahRate, nil
}
