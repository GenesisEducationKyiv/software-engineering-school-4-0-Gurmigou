package rate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"se-school-case/pkg/util"
	"se-school-case/pkg/util/constants"
)

type FetchService struct{}

func NewRateFetchService() FetchService {
	return FetchService{}
}

func (s *FetchService) FetchExchangeRate() (float64, error) {
	resp, err := http.Get(constants.RATE_API_URL)
	if err != nil {
		return 0, fmt.Errorf("error fetching exchange rate: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	var rates []RateAPIDto
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling response: %w", err)
	}

	for _, rateResp := range rates {
		if rateResp.CCY == constants.DefaultCurrentFrom && rateResp.BaseCCY == constants.DefaultCurrentTo {
			exchangeRate := util.ParseFloat(rateResp.Sale)
			return exchangeRate, nil
		}
	}

	return 0, errors.New("rate not found")
}
