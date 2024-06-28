package rate

import (
	"encoding/json"
	"io"
	"net/http"
	"se-school-case/pkg/util"
	"se-school-case/pkg/util/constants"

	"github.com/sirupsen/logrus"
)

type BankRateFetch struct {
	DefaultCurrencyFetcher
}

func NewBankRateFetchService() BankRateFetch {
	return BankRateFetch{}
}

// Provider: privatbank-api
func (s *BankRateFetch) Fetch() (float64, error) {
	resp, err := http.Get(constants.RATE_BANK_API_URL + "TODO:REMOVE")
	if err != nil || resp.StatusCode != http.StatusOK {
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

	var rates []RateAPIDto
	err = json.Unmarshal(body, &rates)
	if err != nil {
		logrus.WithError(err).Error("error unmarshaling response")
		return s.DefaultCurrencyFetcher.Fetch()
	}

	for _, rateResp := range rates {
		if rateResp.CCY == constants.DefaultCurrentFrom &&
			rateResp.BaseCCY == constants.DefaultCurrentTo {
			exchangeRate := util.ParseFloat(rateResp.Sale)
			return exchangeRate, nil
		}
	}

	logrus.Error("rates not found")
	return s.DefaultCurrencyFetcher.Fetch()
}