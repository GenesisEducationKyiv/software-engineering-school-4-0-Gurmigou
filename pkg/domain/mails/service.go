package mails

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	cron_jobs "se-school-case/pkg/domain/cron-jobs"
	"se-school-case/pkg/domain/rates"
	"se-school-case/pkg/domain/subscribers"
	"se-school-case/pkg/util"
	"se-school-case/pkg/util/constants"
	"text/template"
)

type MailInterface interface {
	cron_jobs.MailInterface
	SendEmail(subject string, templatePath string, sendTo string, rate float64) error
}

type MailService struct {
	subscriberService subscribers.SubscriberInterface
	rateService       rates.RateInterface
}

func NewService(subscriberService subscribers.SubscriberInterface,
	rateService rates.RateInterface) MailService {
	return MailService{
		subscriberService: subscriberService,
		rateService:       rateService,
	}
}

func (s *MailService) SendEmailToAll(subject string, templatePath string) {
	users, err := s.subscriberService.GetAll()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
		return
	}

	rateResp, err := s.rateService.GetRate()
	if err != nil {
		log.Fatalf("Failed to get latest rates: %v", err)
		return
	}

	for _, userResp := range users {
		err := s.SendEmail(subject, templatePath, userResp.Email, rateResp.Rate)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", userResp.Email, err)
		}
	}
}

func (s *MailService) SendEmail(subject string, templatePath string, sendTo string, rate float64) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	err = t.Execute(&body, EmailSendDto{
		Email:       sendTo,
		CurrentDate: util.GetCurrentDateString(),
		Rate:        fmt.Sprintf("%.2f", rate),
	})
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		constants.GOOGLE_USERNAME,
		constants.GOOGLE_PASSWORD,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		constants.GOOGLE_USERNAME,
		[]string{sendTo},
		[]byte(msg),
	)
	return err
}
