package mail

import (
	"bytes"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"mailer/pkg/constants"
	"net/smtp"
	"text/template"
	"time"
)

type EmailSendDto struct {
	Email       string
	CurrentDate string
	Rate        string
}

type RabbitMQInterface interface {
	Consume() (<-chan amqp.Delivery, error)
}

type MailServiceInterface interface {
	SendEmail(subject string, templatePath string, sendTo string, rate string) error
}

type MailService struct {
	rabbitMQConn RabbitMQInterface
}

func NewService(rabbitMQConn RabbitMQInterface) MailService {
	return MailService{rabbitMQConn: rabbitMQConn}
}

func getCurrentDateString() string {
	currentDate := time.Now().Format("2006-01-02 15:04")
	return currentDate
}

func (s *MailService) SendEmail(subject string, templatePath string, sendTo string, rate string) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	err = t.Execute(&body, EmailSendDto{
		Email:       sendTo,
		CurrentDate: getCurrentDateString(),
		Rate:        rate,
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
