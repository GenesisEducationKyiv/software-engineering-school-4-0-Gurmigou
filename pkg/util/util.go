package util

import (
	"log"
	"strconv"
	"time"
)

func GetCurrentDateString() string {
	currentDate := time.Now().Format("2006-01-02 15:04")
	return currentDate
}

func ParseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Error parsing float: %v", err)
		return 0.0
	}
	return result
}
