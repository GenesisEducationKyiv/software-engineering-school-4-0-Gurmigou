package service

import "se-school-case/pkg/model"

type RateInterface interface {
	GetRate() (model.Rate, error)
	SaveRate(currencyFrom string, currencyTo string, exchangeRate float64)
}
