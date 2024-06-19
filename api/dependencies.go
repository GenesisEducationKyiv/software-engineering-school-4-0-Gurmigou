package api

import (
	cron_jobs "se-school-case/pkg/domain/cron-jobs"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/domain/rate"
	"se-school-case/pkg/domain/subscriber"
	"se-school-case/pkg/initializer"
)

type dependencies struct {
	subscriberService subscriber.Service
	rateService       rate.Service
	mailService       mail.Service
	cronService       cron_jobs.Service
}

func wireDependencies() *dependencies {
	initializer.RunMigrations()
	db := initializer.ConnectToDatabase()
	rateService := rate.NewService(db)
	subscriberService := subscriber.NewService(db)
	mailService := mail.NewService(subscriberService, rateService)
	cron_jobs.NewService(mailService)
	return &dependencies{
		subscriberService: subscriberService,
		rateService:       rateService,
		mailService:       mailService,
	}
}
