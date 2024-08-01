package initializer

import (
	"log"
	"mailer/internal"
	cronjobs "mailer/internal/cron-jobs"
	"mailer/internal/event"
	"mailer/internal/mail"
	"mailer/pkg/constants"
	"mailer/pkg/queue"
)

type dependencies struct {
	rabbitMQConnection *queue.RabbitMQ
	eventConsumer      event.EventConsumerService
	cronJobsMailer     cronjobs.MailerCronJob
}

func wireDependencies() *dependencies {
	InitEnv()
	ConnectToDatabase()

	// Connection to RabbitMQ
	rabbitMQConn, err := queue.NewRabbitMQConnection(constants.RABBITMQ_URL, constants.QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Create mail service
	mailService := mail.NewService(rabbitMQConn)

	// Create repository to manage subscribers
	repo := internal.NewRepository(DB)

	// Create cron-job mailer service
	cronJobsMailer := cronjobs.NewMailerCronJob(repo, mailService)

	// Create event consumer service
	eventService := event.NewEventConsumerService(repo, mailService, cronJobsMailer, *rabbitMQConn)

	return &dependencies{
		rabbitMQConnection: rabbitMQConn,
		eventConsumer:      eventService,
		cronJobsMailer:     cronJobsMailer,
	}
}
