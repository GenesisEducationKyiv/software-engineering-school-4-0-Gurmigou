package initializer

import (
	"log"
	"mailer/internal/event"
	"mailer/internal/mail"
	"mailer/pkg/constants"
	"mailer/pkg/queue"
)

type dependencies struct {
	rabbitMQConnection *queue.RabbitMQ
	mailService        mail.MailService
	cronJobsConsumer   event.CronJobConsumerService
}

func wireDependencies() *dependencies {
	InitEnv()

	// Connection to RabbitMQ
	rabbitMQConn, err := queue.NewRabbitMQConnection(constants.RABBITMQ_URL, constants.QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Create mail service
	mailService := mail.NewService(rabbitMQConn)

	// Create repository to manage subscribers
	repo := event.NewRepository(DB)

	// Create Cron jobs consumer service
	cronJobConsumerService := event.NewCronJobConsumerService(repo, mailService, *rabbitMQConn)

	return &dependencies{
		rabbitMQConnection: rabbitMQConn,
		mailService:        mailService,
		cronJobsConsumer:   cronJobConsumerService,
	}
}
