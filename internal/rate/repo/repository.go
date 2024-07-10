package repo

import (
	"gorm.io/gorm"
	"se-school-case/pkg/model"
)

type RateRepositoryInterface interface {
	GetLatestRate(currencyFrom string, currencyTo string) (model.Rate, error)
	SaveRate(rate model.Rate) error
	DeleteRates(currencyFrom string, currencyTo string) error
}

type RateRepository struct {
	db *gorm.DB
}

func NewRateRepository(db *gorm.DB) RateRepository {
	return RateRepository{db: db}
}

func (r *RateRepository) GetLatestRate(currencyFrom string, currencyTo string) (model.Rate, error) {
	var rate model.Rate
	err := r.db.Where("currency_from = ? AND currency_to = ?",
		currencyFrom, currencyTo).First(&rate).Error
	if err != nil {
		return model.Rate{}, err
	}
	return rate, nil
}

func (r *RateRepository) SaveRate(rate model.Rate) error {
	if err := r.db.Create(&rate).Error; err != nil {
		return err
	}
	return nil
}

func (r *RateRepository) DeleteRates(currencyFrom string, currencyTo string) error {
	if err := r.db.Where("currency_from = ? AND currency_to = ?",
		currencyFrom, currencyTo).Delete(&model.Rate{}).Error; err != nil {
		return err
	}
	return nil
}
