package api

import (
	"gorm.io/gorm"
	cron_jobs "se-school-case/pkg/domain/cron-jobs"
	"se-school-case/pkg/domain/external-api/rate"
	"se-school-case/pkg/domain/mails"
	"se-school-case/pkg/domain/rates"
	"se-school-case/pkg/domain/subscribers"
	"se-school-case/pkg/initializer"
	"se-school-case/pkg/util/constants"
)

type dependencies struct {
	subscriberService subscribers.SubscriberInterface
	rateService       rates.RateInterface
	mailService       mails.MailInterface
	cronService       cron_jobs.CronJobsInterface
}

func wireDependencies() *dependencies {
	InitEnv()
	db := setUpDatabase()

	// Initialize repositories
	rateRepository := rates.NewRateRepository(db)
	subscriberRepository := subscribers.NewSubscriberRepository(db)

	// Initialize chain of Rate fetchers
	bankFetchService := rate.NewBankRateFetchService()
	exchangeFetchService := rate.NewExchangeApiRateFetch()
	bankFetchService.SetNext(&exchangeFetchService)

	rateService := rates.NewService(&rateRepository, &bankFetchService)
	subscriberService := subscribers.NewService(&subscriberRepository)
	mailService := mails.NewService(&subscriberService, &rateService)
	cronService := cron_jobs.NewService(&mailService)
	return &dependencies{
		subscriberService: &subscriberService,
		rateService:       &rateService,
		mailService:       &mailService,
		cronService:       &cronService,
	}
}

func InitEnv() {
	initializer.LoadEnvVariables()
	constants.InitEnvValues()
}

func setUpDatabase() *gorm.DB {
	initializer.RunMigrations()
	db := initializer.ConnectToDatabase()
	return db
}
