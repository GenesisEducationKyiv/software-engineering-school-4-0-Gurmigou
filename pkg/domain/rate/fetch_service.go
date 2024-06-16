package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	var rates []RateAPIDto

	err = json.Unmarshal(body, &rates)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return
	}

	for _, rate := range rates {
		if rate.CCY == DefaultCurrentFrom && rate.BaseCCY == DefaultCurrentTo {
			exchangeRate := parseFloat(rate.Sale)
			SaveRate(DefaultCurrentFrom, DefaultCurrentTo, exchangeRate)
			break
		}

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
