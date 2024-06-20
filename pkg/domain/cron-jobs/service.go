package cron_jobs

import (
	"github.com/go-co-op/gocron"
	"log"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/util/constants"
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
	scheduler := gocron.NewScheduler(time.Local)

	_, err := scheduler.Every(1).Day().At(
		constants.EMAIL_SEND_TIME).Do(func() {
		s.mailService.SendEmailToAll(
			"Exchange rate notification", constants.TEMPLATE_PATH)
	})
	if err != nil {
		log.Fatalf("Error scheduling email notifications: %v", err)
	}

	scheduler.StartAsync()
}
