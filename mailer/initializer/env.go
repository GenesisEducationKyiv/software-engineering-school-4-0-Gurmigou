package initializer

import (
	"github.com/joho/godotenv"
	"log"
	"se-school-case/mailer/pkg/constants"
)

func InitEnv() {
	loadEnvVariables()
	constants.InitEnvValues()
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
