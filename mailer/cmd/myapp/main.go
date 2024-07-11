package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"se-school-case/mailer/internal/mail"
	"se-school-case/pkg/queue"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("QUEUE_NAME")
	templatePath := os.Getenv("TEMPLATE_PATH")

	rabbitMQConn, err := queue.NewRabbitMQConnection(rabbitMQURL, queueName)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	mailService := mail.NewService(rabbitMQConn)

	msgs, err := rabbitMQConn.Consume()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event mail.Event
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			for _, email := range event.Data.Subscribers {
				err := mailService.SendEmail("Exchange rate notification",
					templatePath, email,
					fmt.Sprintf("%.2f", event.Data.ExchangeRate))
				if err != nil {
					log.Printf("Failed to send email: %v", err)
				}
			}
		}
	}()

	<-forever
}
