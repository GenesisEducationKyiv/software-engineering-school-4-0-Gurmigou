package main

import "se-school-case/mailer/initializer"

func main() {
	consumer := initializer.NewConsumer()
	consumer.ConsumeEvents()
}
