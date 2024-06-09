package service

import (
	"bytes"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"net/smtp"
	"os"
	"se-school-case/pkg/dto"
	"se-school-case/pkg/initializer"
	"se-school-case/pkg/model"
	"text/template"
	"time"
)

// StartScheduledEmail starts a goroutine to send email notifications
// at a specified time every day using gocron.
func StartScheduledEmail() {
	emailTime := os.Getenv("EMAIL_SEND_TIME") // expected format "15:04"
	if emailTime == "" {
		log.Fatalf("EMAIL_SEND_TIME environment variable not set")
	}

	scheduler := gocron.NewScheduler(time.Local)

	// Schedule the email job
	_, err := scheduler.Every(1).Day().At(emailTime).Do(SendEmailNotificationsToAll)
	if err != nil {
		log.Fatalf("Error scheduling email notifications: %v", err)
	}

	// Start the scheduler
	scheduler.StartAsync()
}

func SendEmailNotificationsToAll() {
	sendEmailToAll("Exchange rate notification", os.Getenv("TEMPLATE_PATH"))
}

// SendEmailToAll sends emails to all users in the database with the current exchange rate.
func sendEmailToAll(subject string, templatePath string) {
	users, err := getAllUsers()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
		return
	}

	rate, err := GetRate()
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
		os.Getenv("GOOGLE_USERNAME"),
		os.Getenv("GOOGLE_PASSWORD"),
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("GOOGLE_USERNAME"),
		[]string{sendTo},
		[]byte(msg),
	)
	return err
}

func getCurrentDateString() string {
	currentDate := time.Now().Format("2006-01-02 15:04")
	return currentDate
}
