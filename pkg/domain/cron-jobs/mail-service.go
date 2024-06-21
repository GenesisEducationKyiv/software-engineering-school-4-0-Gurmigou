package cron_jobs

type MailInterface interface {
	SendEmailToAll(subject string, templatePath string)
}
