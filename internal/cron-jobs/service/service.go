package cron_jobs

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"log"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/model"
	"time"
)

type Event struct {
	EventID     string    `json:"eventId"`
	EventType   string    `json:"eventType"`
	AggregateID string    `json:"aggregateId"`
	Timestamp   string    `json:"timestamp"`
	Data        EventData `json:"data"`
}

type EventData struct {
	ExchangeRate float64  `json:"exchangeRate"`
	Subscribers  []string `json:"subscribers"`
}

type MailInterface interface {
	SendEmail(subject string, templatePath string, sendTo string, rate float64) error
}

type RateInterface interface {
	GetRate() (model.Rate, error)
	SaveRate(currencyFrom string, currencyTo string, exchangeRate float64)
}

type SubscriberInterface interface {
	Add(email string) error
	GetAll() ([]model.User, error)
}

type CronJobsService struct {
	rabbitMQService   RabbitMQInterface
	subscriberService SubscriberInterface
	rateService       RateInterface
}

type RabbitMQInterface interface {
	Publish(message string) error
}

func NewService(
	rabbitMQService RabbitMQInterface,
	subscriberService SubscriberInterface,
	rateService RateInterface) CronJobsService {
	return CronJobsService{
		rabbitMQService:   rabbitMQService,
		subscriberService: subscriberService,
		rateService:       rateService,
	}
}

func (s *CronJobsService) StartScheduler() {
	scheduler := gocron.NewScheduler(time.Local)

	_, err := scheduler.Every(1).Day().At(constants.EMAIL_SEND_TIME).Do(s.NotifySubscribers)
	if err != nil {
		log.Fatalf("Error scheduling email notifications: %v", err)
	}

	scheduler.StartAsync()
}

func (s *CronJobsService) NotifySubscribers() error {
	users, err := s.subscriberService.GetAll()
	if err != nil {
		return err
	}

	rate, err := s.rateService.GetRate()
	if err != nil {
		return err
	}

	subscribers := make([]string, len(users))
	for i, user := range users {
		subscribers[i] = user.Email
	}

	event := Event{
		EventID:     "1",
		EventType:   "RateNotification",
		AggregateID: "rate-1",
		Timestamp:   time.Now().Format(time.RFC3339),
		Data: EventData{
			ExchangeRate: rate.Rate,
			Subscribers:  subscribers,
		},
	}

	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = s.rabbitMQService.Publish(string(message))
	if err != nil {
		return err
	}
	return nil
}
