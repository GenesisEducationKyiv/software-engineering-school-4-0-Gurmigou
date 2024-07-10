package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/util"
	"text/template"
)

type EmailSendDto struct {
	Email       string
	CurrentDate string
	Rate        string
}

type MailService struct{}

func NewService() MailService {
	return MailService{}
}

func (s *MailService) SendEmail(subject string, templatePath string, sendTo string, rate float64) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	err = t.Execute(&body, EmailSendDto{
		Email:       sendTo,
		CurrentDate: util.GetCurrentDateString(),
		Rate:        fmt.Sprintf("%.2f", rate),
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
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
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
