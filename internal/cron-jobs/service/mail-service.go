package service

type MailInterface interface {
	SendEmail(subject string, templatePath string, sendTo string, rate float64) error
	SendEmailToAll(subject string, templatePath string) error
}
