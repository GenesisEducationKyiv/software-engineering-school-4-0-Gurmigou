package service

type MailInterface interface {
	SendEmailToAll(subject string, templatePath string) error
}
