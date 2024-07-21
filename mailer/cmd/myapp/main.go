package main

import "mailer/initializer"

func main() {
	consumer := initializer.NewConsumer()
	consumer.ConsumeEvents()
}
