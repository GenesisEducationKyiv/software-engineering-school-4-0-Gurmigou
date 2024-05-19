package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"se-school-case/dto"
	"se-school-case/initializer"
	"se-school-case/model"
	"strconv"
	"time"
)

const (
	updateInterval = 1 * time.Hour
)

func fetchExchangeRate() {
	resp, err := http.Get(os.Getenv("RATE_API_URL"))
	if err != nil {
		fmt.Println("Error fetching exchange rate")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var rates []dto.RateApiDto

	err = json.Unmarshal(body, &rates)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return
	}

	for _, rate := range rates {
		if rate.CCY == DefaultCurrentFrom && rate.BaseCCY == DefaultCurrentTo {
			exchangeRate := parseFloat(rate.Sale)
			writeResultToDatabase(DefaultCurrentFrom, DefaultCurrentTo, exchangeRate)
			break
		}
	}
}

func writeResultToDatabase(currencyFrom string, currencyTo string, exchangeRate float64) {
	// Delete existing rate records where CurrencyFrom and CurrencyTo match
	if err := initializer.DB.Where("currency_from = ? AND currency_to = ?",
		currencyFrom, currencyTo).Delete(&model.Rate{}).Error; err != nil {
		log.Printf("Error deleting old exchange rates: %v", err)
		return
	}

	// Add new rate record
	rate := model.Rate{CurrencyFrom: currencyFrom, CurrencyTo: currencyTo, Rate: exchangeRate}
	if err := initializer.DB.Create(&rate).Error; err != nil {
		log.Printf("Error to write exchange rate to database: %v", err)
		return
	}
}

func parseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Error parsing float: %v", err)
		return 0.0
	}
	return result
}

func StartRateUpdater() {
	go func() {
		for {
			fetchExchangeRate()
			time.Sleep(updateInterval)
		}
	}()
}
