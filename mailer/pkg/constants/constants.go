package constants

import "os"

var (
	DB_URL          string
	DB_FULL_URL     string
	GOOGLE_USERNAME string
	GOOGLE_PASSWORD string
	TEMPLATE_PATH   string
	RABBITMQ_URL    string
	QUEUE_NAME      string
)

func InitEnvValues() {
	DB_URL = os.Getenv("DB_URL")
	DB_FULL_URL = os.Getenv("DB_FULL_URL")
	GOOGLE_USERNAME = os.Getenv("GOOGLE_USERNAME")
	GOOGLE_PASSWORD = os.Getenv("GOOGLE_PASSWORD")
	TEMPLATE_PATH = os.Getenv("TEMPLATE_PATH")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	QUEUE_NAME = os.Getenv("QUEUE_NAME")
}
