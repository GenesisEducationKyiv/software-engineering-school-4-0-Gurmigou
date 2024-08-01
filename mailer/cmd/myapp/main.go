package main

import "mailer/initializer"

func main() {
	consumer := initializer.NewApi()
	consumer.StartApplication()
}
