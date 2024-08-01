package constants

import (
	"os"
	"time"
)

const (
	DefaultCurrentFrom = "USD"
	DefaultCurrentTo   = "UAH"
	UpdateInterval     = 1 * time.Hour
)

var (
	PORT         string
	DB_URL       string
	DB_FULL_URL  string
	RABBITMQ_URL string
	QUEUE_NAME   string
)

func InitEnvValues() {
	PORT = os.Getenv("PORT")
	DB_URL = os.Getenv("DB_URL")
	DB_FULL_URL = os.Getenv("DB_FULL_URL")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	QUEUE_NAME = os.Getenv("QUEUE_NAME")
}
