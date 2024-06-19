package rate

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"se-school-case/pkg/model"
	"se-school-case/pkg/util"
	"se-school-case/pkg/util/constants"
	"time"
)

type Service interface {
	GetRate() (model.Rate, error)
	SaveRate(currencyFrom string, currencyTo string, exchangeRate float64)
}

type service struct {
	repository *gorm.DB
}

func NewService(repository *gorm.DB) Service {
	return &service{repository}
}

var ErrNoRateFound = errors.New("no rate found")

func (s *service) GetRate() (model.Rate, error) {
	rate, err := s.getLatestRate()
	if err != nil {
		if errors.Is(err, ErrNoRateFound) || err.Error() == "record not found" {
			// Fetch exchange rate if no rate found
			s.FetchExchangeRate()
			rate, err = s.getLatestRate()
			if err != nil {
				return model.Rate{}, err
			}
		} else {
			return model.Rate{}, err
		}
	}

	// Check if the rate is more than 1 hour old
	if time.Since(rate.CreatedAt) > constants.UpdateInterval {
		s.FetchExchangeRate()
		rate, err = s.getLatestRate()
		if err != nil {
			return model.Rate{}, err
		}
	}

	return rate, nil
}

func (s *service) SaveRate(currencyFrom string, currencyTo string, exchangeRate float64) {
	// Delete existing rate records where CurrencyFrom and CurrencyTo match
	if err := s.repository.Where("currency_from = ? AND currency_to = ?",
		currencyFrom, currencyTo).Delete(&model.Rate{}).Error; err != nil {
		log.Printf("Error deleting old exchange rates: %v", err)
		return
	}

	// Add new rate record
	rate := model.Rate{CurrencyFrom: currencyFrom, CurrencyTo: currencyTo, Rate: exchangeRate}
	if err := s.repository.Create(&rate).Error; err != nil {
		log.Printf("Error writing exchange rate to database: %v", err)
		return
	}
}

func (s *service) FetchExchangeRate() {
	resp, err := http.Get(os.Getenv("RATE_API_URL"))
	if err != nil {
		fmt.Println("Error fetching exchange rate")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var rates []RateAPIDto

	err = json.Unmarshal(body, &rates)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return
	}

	for _, rateResp := range rates {
		if rateResp.CCY == constants.DefaultCurrentFrom && rateResp.BaseCCY == constants.DefaultCurrentTo {
			exchangeRate := util.ParseFloat(rateResp.Sale)
			s.SaveRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo, exchangeRate)
			break
		}
	}
}

func (s *service) getLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := s.repository.Where("currency_from = ? AND currency_to = ?",
		constants.DefaultCurrentFrom, constants.DefaultCurrentTo).First(&rate).Error
	if err != nil {
		return model.Rate{}, err
	}
	return rate, nil
}
