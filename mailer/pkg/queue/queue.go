package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQConnection(url string, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		err := ch.Close()
		if err != nil {
			return nil, err
		}
		err = conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &RabbitMQ{
		conn:      conn,
		channel:   ch,
		queueName: queue.Name,
	}, nil
}

func (r *RabbitMQ) Publish(message string) error {
	err := r.channel.Publish(
		"",          // exchange
		r.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

func (r *RabbitMQ) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *RabbitMQ) Close() {
	err := r.channel.Close()
	if err != nil {
		return
	}
	err = r.conn.Close()
	if err != nil {
		return
	}
}
