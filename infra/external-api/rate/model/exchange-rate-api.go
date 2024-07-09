package model

type ExchangeRateAPI struct {
	Date    string             `json:"date"`
	RateMap map[string]float64 `json:"usd"`
}
