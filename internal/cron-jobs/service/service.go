package service

import (
	"github.com/go-co-op/gocron"
	"log"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/model"
	"time"
)

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
	mailService       MailInterface
	subscriberService SubscriberInterface
	rateService       RateInterface
}

func NewService(mailService MailInterface,
	subscriberService SubscriberInterface,
	rateService RateInterface) CronJobsService {
	return CronJobsService{
		mailService:       mailService,
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

	for _, user := range users {
		err := s.mailService.SendEmail("Exchange rate notification", constants.TEMPLATE_PATH, user.Email, rate.Rate)
		if err != nil {
			return err
		}
	}
	return nil
}
