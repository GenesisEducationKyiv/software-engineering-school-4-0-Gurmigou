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
	PORT            string
	DB_URL          string
	DB_FULL_URL     string
	GOOGLE_USERNAME string
	GOOGLE_PASSWORD string
	TEMPLATE_PATH   string
	EMAIL_SEND_TIME string
	RABBITMQ_URL    string
	QUEUE_NAME      string
)

func InitEnvValues() {
	PORT = os.Getenv("PORT")
	DB_URL = os.Getenv("DB_URL")
	DB_FULL_URL = os.Getenv("DB_FULL_URL")
	GOOGLE_USERNAME = os.Getenv("GOOGLE_USERNAME")
	GOOGLE_PASSWORD = os.Getenv("GOOGLE_PASSWORD")
	TEMPLATE_PATH = os.Getenv("TEMPLATE_PATH")
	EMAIL_SEND_TIME = os.Getenv("EMAIL_SEND_TIME")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	QUEUE_NAME = os.Getenv("QUEUE_NAME")
}
