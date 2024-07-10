package service

import (
	"errors"
	"log"
	"se-school-case/infra/external-api/rate"
	"se-school-case/internal/rate/repo"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/model"
	"time"
)

type RateService struct {
	repository   repo.RateRepositoryInterface
	fetchService rate.CurrencyFetcher
}

func NewService(repository repo.RateRepositoryInterface, currencyFetcher rate.CurrencyFetcher) RateService {
	return RateService{repository: repository, fetchService: currencyFetcher}
}

var ErrNoRateFound = errors.New("no rate found")

func (s *RateService) GetRate() (model.Rate, error) {
	rate, err := s.repository.GetLatestRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo)
	if err != nil {
		if errors.Is(err, ErrNoRateFound) || err.Error() == "record not found" {
			// Fetch exchange rate if no rate found
			exchangeRate, fetchErr := s.fetchService.Fetch()
			if fetchErr != nil {
				return model.Rate{}, fetchErr
			}
			s.SaveRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo, exchangeRate)
			rate, err = s.repository.GetLatestRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo)
			if err != nil {
				return model.Rate{}, err
			}
		} else {
			return model.Rate{}, err
		}
	}

	// Check if the rate is more than 1 hour old
	if time.Since(rate.CreatedAt) > constants.UpdateInterval {
		exchangeRate, fetchErr := s.fetchService.Fetch()
		if fetchErr != nil {
			return model.Rate{}, fetchErr
		}
		s.SaveRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo, exchangeRate)
		rate, err = s.repository.GetLatestRate(constants.DefaultCurrentFrom, constants.DefaultCurrentTo)
		if err != nil {
			return model.Rate{}, err
		}
	}

	return rate, nil
}

func (s *RateService) SaveRate(currencyFrom string, currencyTo string, exchangeRate float64) {
	// Delete existing rate records where CurrencyFrom and CurrencyTo match
	if err := s.repository.DeleteRates(currencyFrom, currencyTo); err != nil {
		log.Printf("Error deleting old exchange rate: %v", err)
		return
	}

	rate := model.Rate{CurrencyFrom: currencyFrom, CurrencyTo: currencyTo, Rate: exchangeRate}
	if err := s.repository.SaveRate(rate); err != nil {
		log.Printf("Error writing exchange rate to database: %v", err)
		return
	}
}
