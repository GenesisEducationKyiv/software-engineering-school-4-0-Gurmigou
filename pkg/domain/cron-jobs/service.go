package cron_jobs

import (
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"se-school-case/pkg/domain/mail"
	"time"
)

type Service interface {
	StartScheduler()
}

type service struct {
	mailService mail.Service
}

func NewService(mailService mail.Service) Service {
	return &service{mailService}
}

func (s *service) StartScheduler() {
	emailTime := os.Getenv("EMAIL_SEND_TIME") // expected format "15:04"
	if emailTime == "" {
		log.Fatalf("EMAIL_SEND_TIME environment variable not set")
	}

	scheduler := gocron.NewScheduler(time.Local)

	// Schedule the email job
	_, err := scheduler.Every(1).Day().At(emailTime).Do(func() {
		s.mailService.SendEmailToAll("Exchange rate notification",
			os.Getenv("TEMPLATE_PATH"))
	})
	if err != nil {
		log.Fatalf("Error scheduling email notifications: %v", err)
	}

	// Start the scheduler
	scheduler.StartAsync()
}
