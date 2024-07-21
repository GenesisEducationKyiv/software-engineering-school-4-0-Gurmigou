package initializer

import (
	"log"
	"mailer/internal/mail"
	"mailer/pkg/constants"
	"mailer/pkg/queue"
)

type dependencies struct {
	rabbitMQConnection *queue.RabbitMQ
	mailService        mail.MailService
}

func wireDependencies() *dependencies {
	InitEnv()

	// Connection to RabbitMQ
	rabbitMQConn, err := queue.NewRabbitMQConnection(constants.RABBITMQ_URL, constants.QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	//defer rabbitMQConn.Close()

	// Create mail service
	mailService := mail.NewService(rabbitMQConn)
	return &dependencies{
		rabbitMQConnection: rabbitMQConn,
		mailService:        mailService,
	}
}
