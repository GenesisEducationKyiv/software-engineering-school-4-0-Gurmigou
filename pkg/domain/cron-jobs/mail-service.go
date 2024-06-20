package cron_jobs

type MailService interface {
	SendEmailToAll(subject string, templatePath string)
}
