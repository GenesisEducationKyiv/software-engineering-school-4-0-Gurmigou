package constants

import "os"

var (
	DB_URL          string
	GOOGLE_USERNAME string
	GOOGLE_PASSWORD string
	TEMPLATE_PATH   string
	RABBITMQ_URL    string
	QUEUE_NAME      string
	EMAIL_SEND_TIME string
)

func InitEnvValues() {
	DB_URL = os.Getenv("DB_URL")
	GOOGLE_USERNAME = os.Getenv("GOOGLE_USERNAME")
	GOOGLE_PASSWORD = os.Getenv("GOOGLE_PASSWORD")
	TEMPLATE_PATH = os.Getenv("TEMPLATE_PATH")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	QUEUE_NAME = os.Getenv("QUEUE_NAME")
	EMAIL_SEND_TIME = os.Getenv("EMAIL_SEND_TIME")
}
