package service

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"se-school-case/dto"
	"se-school-case/initializer"
	"se-school-case/model"
	"text/template"
	"time"
)

func SendEmailNotificationsToAll() {
	sendEmailToAll("Exchange rate notification", "./resource/email.html")
}

// SendEmailToAll sends emails to all users in the database with the current exchange rate.
func sendEmailToAll(subject string, templatePath string) {
	users, err := getAllUsers()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
		return
	}

	rate, err := getLatestRate()
	if err != nil {
		log.Fatalf("Failed to get latest rate: %v", err)
		return
	}

	for _, user := range users {
		err := sendEmail(subject, templatePath, user.Email, rate.Rate)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", user.Email, err)
		}
	}
}

func getAllUsers() ([]model.User, error) {
	var users []model.User
	err := initializer.DB.Find(&users).Error
	return users, err
}

func getLatestRate() (model.Rate, error) {
	var rate model.Rate
	err := initializer.DB.Where("currency_from = ? AND currency_to = ?",
		DefaultCurrentFrom, DefaultCurrentTo).First(&rate).Error
	return rate, err
}

func sendEmail(subject string, templatePath string, sendTo string, rate float64) error {
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	err = t.Execute(&body, dto.EmailSendDto{
		Email:       sendTo,
		CurrentDate: getCurrentDateString(),
		Rate:        fmt.Sprintf("%.2f", rate),
	})
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth(
		"",
		"se.school.case.2024.notification@gmail.com",
		"tyctaulbmtblferm",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"se.school.case.2024.notification@gmail.com",
		[]string{sendTo},
		[]byte(msg),
	)
	return err
}

func getCurrentDateString() string {
	currentDate := time.Now().Format("2006-01-02 15:04")
	return currentDate
}
