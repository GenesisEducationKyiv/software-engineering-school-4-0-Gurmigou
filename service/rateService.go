package service

import (
	"se-school-case/initializer"
	"se-school-case/model"
)

const (
	DefaultCurrentFrom = "USD"
	DefaultCurrentTo   = "UAH"
)

func GetLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := initializer.DB.Where("currency_from = ? AND currency_to = ?",
		DefaultCurrentFrom, DefaultCurrentTo).First(&rate).Error
	return rate, err
}
