package initializer

import (
	"gorm.io/gorm"
	"se-school-case/db"
	"se-school-case/infra/external-api/rate/provider"
	cronjobs "se-school-case/internal/cron-jobs/handler"
	jobsservice "se-school-case/internal/cron-jobs/service"
	mailsservice "se-school-case/internal/mails/service"
	rateshandler "se-school-case/internal/rates/handler"
	ratesrepo "se-school-case/internal/rates/repo"
	ratesservice "se-school-case/internal/rates/service"
	subhandler "se-school-case/internal/subscribers/handler"
	subrepo "se-school-case/internal/subscribers/repo"
	subservice "se-school-case/internal/subscribers/service"
	"se-school-case/pkg/constants"
)

type dependencies struct {
	subscriberService subhandler.SubscriberInterface
	rateService       rateshandler.RateInterface
	mailService       mailsservice.MailInterface
	cronService       cronjobs.CronJobsInterface
}

func wireDependencies() *dependencies {
	InitEnv()
	db := setUpDatabase()

	// Initialize repositories
	rateRepository := ratesrepo.NewRateRepository(db)
	subscriberRepository := subrepo.NewSubscriberRepository(db)

	// Initialize chain of Rate fetchers
	bankFetchService := provider.NewBankRateFetchService()
	exchangeFetchService := provider.NewExchangeApiRateFetch()
	bankFetchService.SetNext(&exchangeFetchService)

	rateService := ratesservice.NewService(&rateRepository, &bankFetchService)
	subscriberService := subservice.NewService(&subscriberRepository)
	mailService := mailsservice.NewService(&subscriberService, &rateService)
	cronService := jobsservice.NewService(&mailService)
	return &dependencies{
		subscriberService: &subscriberService,
		rateService:       &rateService,
		mailService:       &mailService,
		cronService:       &cronService,
	}
}

func InitEnv() {
	LoadEnvVariables()
	constants.InitEnvValues()
}

func setUpDatabase() *gorm.DB {
	db.RunMigrations()
	return db.ConnectToDatabase()
}
