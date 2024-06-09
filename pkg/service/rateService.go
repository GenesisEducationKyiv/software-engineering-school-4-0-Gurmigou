package service

import (
	"errors"
	"se-school-case/pkg/initializer"
	"se-school-case/pkg/model"
	"time"
)

const (
	DefaultCurrentFrom = "USD"
	DefaultCurrentTo   = "UAH"
	updateInterval     = 1 * time.Hour
)

var ErrNoRateFound = errors.New("no rate found")

func getLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := initializer.DB.Where("currency_from = ? AND currency_to = ?",
		DefaultCurrentFrom, DefaultCurrentTo).First(&rate).Error
	if err != nil {
		return model.Rate{}, err
	}
	return rate, nil
}

func GetRate() (model.Rate, error) {
	rate, err := getLatestRate()
	if err != nil {
		if errors.Is(err, ErrNoRateFound) || err.Error() == "record not found" {
			// Fetch exchange rate if no rate found
			fetchExchangeRate()
			rate, err = getLatestRate()
			if err != nil {
				return model.Rate{}, err
			}
		} else {
			return model.Rate{}, err
		}
	}

	// Check if the rate is more than 1 hour old
	if time.Since(rate.CreatedAt) > updateInterval {
		fetchExchangeRate()
		rate, err = getLatestRate()
		if err != nil {
			return model.Rate{}, err
		}
	}

	return rate, nil
}
