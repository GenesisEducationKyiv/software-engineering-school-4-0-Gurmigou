package provider

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"se-school-case/infra/external-api/rate"
	"se-school-case/infra/external-api/rate/model"
)

const (
	ExchangeApiUrl = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.json"
)

type ExchangeApiRateFetch struct {
	rate.DefaultCurrencyFetcher
}

func NewExchangeApiRateFetch() ExchangeApiRateFetch {
	return ExchangeApiRateFetch{}
}

// Fetch Provider: Exchange Rate API service
func (s *ExchangeApiRateFetch) Fetch() (float64, error) {
	resp, err := http.Get(ExchangeApiUrl)
	if err != nil {
		logrus.WithError(err).Error("error fetching exchange rate")
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
	logrus.Info("Exchange Rate API Provider response: ", string(body))

	var exchangeRate model.ExchangeRateAPI
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		logrus.WithError(err).Error("error unmarshaling response")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	uahRate := exchangeRate.RateMap["uah"]
	if uahRate == 0 {
		logrus.Error("rate not found")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	return uahRate, nil
}
