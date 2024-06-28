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
	db := connectToDb()

	// Initialize chain of Rate fetchers
	bankFetchService := rate.NewBankRateFetchService()
	exchangeFetchService := rate.NewExchangeApiRateFetch()
	bankFetchService.SetNext(&exchangeFetchService)

	rateService := rates.NewService(db, &bankFetchService)
	subscriberService := subscribers.NewService(db)
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

func connectToDb() *gorm.DB {
	initializer.RunMigrations()
	db := initializer.ConnectToDatabase()
	return db
}
