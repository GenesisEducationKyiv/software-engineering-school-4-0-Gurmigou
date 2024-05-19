package service

import (
	"se-school-case/initializer"
	"se-school-case/model"
)

const (
	defaultCurrentFrom = "USD"
	defaultCurrentTo   = "UAH"
)

func GetLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := initializer.DB.Where("currency_from = ? AND currency_to = ?",
		defaultCurrentFrom, defaultCurrentTo).First(&rate).Error
	return rate, err
}
