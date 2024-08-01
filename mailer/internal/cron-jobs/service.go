package cron_jobs

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"mailer/internal"
	"mailer/internal/mail"
	"mailer/pkg/constants"
	"time"
)

type MailerCronJob struct {
	userRepo    internal.Repository
	mailService mail.MailService
}

func NewMailerCronJob(userRepository internal.Repository, mailService mail.MailService) MailerCronJob {
	return MailerCronJob{
		userRepo:    userRepository,
		mailService: mailService,
	}
}

func (m *MailerCronJob) StartScheduler() {
	scheduler := gocron.NewScheduler(time.Local)

	_, err := scheduler.Every(1).Day().At(constants.EMAIL_SEND_TIME).Do(m.SendEmails)
	if err != nil {
		log.Fatalf("Error scheduling email notifications: %v", err)
	}

	scheduler.StartAsync()
}

func (m *MailerCronJob) SendEmails() {
	allSubscribers, err := m.userRepo.FindAll()
	if err != nil {
		log.Printf("Failed to fetch subscribers: %v", err)
	}

	for _, user := range allSubscribers {
		err := m.mailService.SendEmail("Exchange rate notification",
			constants.TEMPLATE_PATH, user.Email,
			fmt.Sprintf("%.2f", internal.CurrentRate))
		if err != nil {
			log.Printf("Failed to send email: %v", err)
		}
	}
}
