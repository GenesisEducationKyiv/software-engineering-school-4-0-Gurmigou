package cron_jobs

import (
	"fmt"
	"log"
	"mailer/internal/event"
	"mailer/internal/mail"
	"mailer/pkg/constants"
)

type MailerCronJob struct {
	userRepo    event.Repository
	mailService mail.MailService
}

func NewMailerCronJob(userRepository event.Repository, mailService mail.MailService) MailerCronJob {
	return MailerCronJob{
		userRepo:    userRepository,
		mailService: mailService,
	}
}

func (m *MailerCronJob) SendEmails() {
	allSubscribers, err := m.userRepo.FindAll()
	if err != nil {
		log.Printf("Failed to fetch subscribers: %v", err)
	}

	for _, user := range allSubscribers {
		err := m.mailService.SendEmail("Exchange rate notification",
			constants.TEMPLATE_PATH, user.Email,
			fmt.Sprintf("%.2f", event.CurrentRate))
		if err != nil {
			log.Printf("Failed to send email: %v", err)
		}
	}
}
