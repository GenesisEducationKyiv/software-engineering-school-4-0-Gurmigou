package cron_jobs

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"log"
	"math/rand"
	"se-school-case/pkg/model"
	"strconv"
	"time"
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

type RateInterface interface {
	GetRate() (model.Rate, error)
	SaveRate(currencyFrom string, currencyTo string, exchangeRate float64)
}

type SubscriberInterface interface {
	Exists(email string) (bool, error)
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

	// Notify immediately on startup
	err := s.NotifyAboutExchangeRate()
	if err != nil {
		return
	}

	_, err = scheduler.Every(1).Hour().Do(s.NotifyAboutExchangeRate)
	if err != nil {
		log.Fatalf("Error scheduling exchange rate notifications: %v", err)
	}

	scheduler.StartAsync()
}

func (s *CronJobsService) NotifyAboutExchangeRate() error {
	rate, err := s.rateService.GetRate()
	if err != nil {
		return err
	}

	event := Event{
		EventID:     strconv.Itoa(rand.Intn(9999)),
		EventType:   CurrencyRateNotification,
		AggregateID: "rate-1",
		Timestamp:   time.Now().Format(time.RFC3339),
		Data: EventData{
			ExchangeRate: rate.Rate,
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

func (s *CronJobsService) ExplicitlyNotify() error {
	event := Event{
		EventID:     strconv.Itoa(rand.Intn(9999)),
		EventType:   ExplicitlyNotify,
		AggregateID: "rate-1",
		Timestamp:   time.Now().Format(time.RFC3339),
		Data:        EventData{},
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
