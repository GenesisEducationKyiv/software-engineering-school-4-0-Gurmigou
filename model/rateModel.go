package model

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	CurrencyFrom string  `gorm:"not null"`
	CurrencyTo   string  `gorm:"not null"`
	Rate         float64 `gorm:"not null"`
}
