package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/smtp"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/util"
	"text/template"
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

type EmailSendDto struct {
	Email       string
	CurrentDate string
	Rate        string
}

type RabbitMQInterface interface {
	Consume() (<-chan amqp.Delivery, error)
}

type MailService struct {
	rabbitMQConn RabbitMQInterface
}

func NewService(rabbitMQConn RabbitMQInterface) MailService {
	return MailService{rabbitMQConn: rabbitMQConn}
}

func (s *MailService) StartConsumer() {
	msgs, err := s.rabbitMQConn.Consume()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event Event
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			for _, email := range event.Data.Subscribers {
				err := s.SendEmail("Exchange rate notification",
					constants.TEMPLATE_PATH, email,
					fmt.Sprintf("%.2f", event.Data.ExchangeRate))
				if err != nil {
					log.Printf("Failed to send email: %v", err)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (s *MailService) SendEmail(subject string, templatePath string, sendTo string, rate string) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	err = t.Execute(&body, EmailSendDto{
		Email:       sendTo,
		CurrentDate: util.GetCurrentDateString(),
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
