package rates

import (
	"errors"
	"log"
	"se-school-case/pkg/model"
	"se-school-case/pkg/util/constants"
	"time"
)

type RateService struct {
	repository   RateRepository
	fetchService RateFetchInterface
}

func NewService(repository RateRepository, rateFetch RateFetchInterface) RateService {
	return RateService{repository: repository, fetchService: rateFetch}
}

var ErrNoRateFound = errors.New("no rates found")

func (s *RateService) GetRate() (model.Rate, error) {
	rate, err := s.getLatestRate()
	if err != nil {
		if errors.Is(err, ErrNoRateFound) || err.Error() == "record not found" {
			// Fetch exchange rates if no rates found
			exchangeRate, fetchErr := s.fetchService.FetchExchangeRate()
			if fetchErr != nil {
				return model.Rate{}, fetchErr
			}
			s.SaveRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo, exchangeRate)
			rate, err = s.getLatestRate()
			if err != nil {
				return model.Rate{}, err
			}
		} else {
			return model.Rate{}, err
		}
	}

	// Check if the rates is more than 1 hour old
	if time.Since(rate.CreatedAt) > constants.UpdateInterval {
		exchangeRate, fetchErr := s.fetchService.FetchExchangeRate()
		if fetchErr != nil {
			return model.Rate{}, fetchErr
		}
		s.SaveRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo, exchangeRate)
		rate, err = s.getLatestRate()
		if err != nil {
			return model.Rate{}, err
		}
	}

	return rate, nil
}

func (s *RateService) SaveRate(currencyFrom string, currencyTo string, exchangeRate float64) {
	// Delete existing rates records where CurrencyFrom and CurrencyTo match
	if err := s.repository.Where("currency_from = ? AND currency_to = ?",
		currencyFrom, currencyTo).Delete(&model.Rate{}); err != nil {
		log.Printf("Error deleting old exchange rates: %v", err)
		return
	}

	rate := model.Rate{CurrencyFrom: currencyFrom, CurrencyTo: currencyTo, Rate: exchangeRate}
	if err := s.repository.Create(&rate); err != nil {
		log.Printf("Error writing exchange rates to database: %v", err)
		return
	}
}

func (s *RateService) getLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := s.repository.Where("currency_from = ? AND currency_to = ?",
		constants.DefaultCurrentFrom, constants.DefaultCurrentTo).First(&rate)
	if err != nil {
		return model.Rate{}, err
	}
	return rate, nil
}
