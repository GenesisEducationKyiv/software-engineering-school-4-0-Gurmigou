package event

import (
	"encoding/json"
	"log"
	"mailer/internal"
	cronjobs "mailer/internal/cron-jobs"
	"mailer/internal/mail"
	"mailer/pkg/queue"
)

type EventType string

const (
	Subscribe                EventType = "Subscribe"
	Unsubscribe              EventType = "Unsubscribe"
	CurrencyRateNotification EventType = "CurrencyRateNotification"
	ExplicitlyNotify         EventType = "ExplicitlyNotify"
)

type Event struct {
	EventID     string    `json:"eventId"`
	EventType   EventType `json:"eventType"`
	AggregateID string    `json:"aggregateId"`
	Timestamp   string    `json:"timestamp"`
	Data        EventData `json:"data"`
}

type EventData struct {
	CurrentDate  string  `json:"currentDate"`
	ExchangeRate float64 `json:"exchangeRate"`
	Email        string  `json:"email"`
}

type EventConsumerService struct {
	repo          internal.Repository
	mailService   mail.MailService
	mailerCronJob cronjobs.MailerCronJob
	rabbitMQ      queue.RabbitMQ
}

func NewEventConsumerService(repo internal.Repository,
	mailService mail.MailService,
	mailerCronJob cronjobs.MailerCronJob,
	rabbitMQ queue.RabbitMQ) EventConsumerService {
	return EventConsumerService{
		repo:          repo,
		mailService:   mailService,
		mailerCronJob: mailerCronJob,
		rabbitMQ:      rabbitMQ,
	}
}

func (c *EventConsumerService) ConsumeEvents() {
	msgs, err := c.rabbitMQ.Consume()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event Event
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			var eventType = event.EventType
			switch eventType {
			case Subscribe:
				c.handleSubscribeEvent(event)
			case Unsubscribe:
				c.handleUnsubscribeEvent(event)
			case CurrencyRateNotification:
				c.handleCurrencyRateNotificationEvent(event)
			case ExplicitlyNotify:
				c.handleExplicitlyNotify()
			}
		}
	}()
	<-forever
}

func (c *EventConsumerService) handleSubscribeEvent(event Event) {
	user := &internal.User{
		Email: event.Data.Email,
	}

	existingUser, err := c.repo.FindByEmail(event.Data.Email)
	if err != nil {
		log.Printf("Error checking if user exists: %v", err)
		return
	}

	if existingUser != nil {
		log.Printf("User with email %s already exists", event.Data.Email)
		return
	}

	if err := c.repo.Create(user); err != nil {
		log.Printf("Error creating user: %v", err)
		return
	}

	log.Printf("User with email %s subscribed successfully", event.Data.Email)
}

func (c *EventConsumerService) handleUnsubscribeEvent(event Event) {
	user, err := c.repo.FindByEmail(event.Data.Email)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return
	}

	if user == nil {
		log.Printf("User with email %s not found", event.Data.Email)
		return
	}

	if err := c.repo.DeleteOne(user); err != nil {
		log.Printf("Error deleting user: %v", err)
		return
	}

	log.Printf("User with email %s unsubscribed successfully", event.Data.Email)
}

func (c *EventConsumerService) handleCurrencyRateNotificationEvent(event Event) {
	internal.CurrentRate = event.Data.ExchangeRate
}

func (c *EventConsumerService) handleExplicitlyNotify() {
	c.mailerCronJob.SendEmails()
}
