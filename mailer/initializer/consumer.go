package initializer

import (
	"encoding/json"
	"fmt"
	"log"
	"mailer/internal/mail"
	"mailer/pkg/constants"
)

type ConsumerApplication interface {
	ConsumeEvents()
}

type consumerApplication struct {
	deps *dependencies
}

func NewConsumer() ConsumerApplication {
	deps := wireDependencies()
	return &consumerApplication{deps}
}

func (c *consumerApplication) ConsumeEvents() {
	msgs, err := c.deps.rabbitMQConnection.Consume()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event mail.Event
			err := json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			for _, email := range event.Data.Subscribers {
				err := c.deps.mailService.SendEmail("Exchange rate notification",
					constants.TEMPLATE_PATH, email,
					fmt.Sprintf("%.2f", event.Data.ExchangeRate))
				if err != nil {
					log.Printf("Failed to send email: %v", err)
				}
			}
		}
	}()

	<-forever
}
