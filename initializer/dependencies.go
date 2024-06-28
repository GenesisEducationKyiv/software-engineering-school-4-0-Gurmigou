package initializer

import (
	"gorm.io/gorm"
	"se-school-case/db"
	cronjobs "se-school-case/internal/cron-jobs"
	"se-school-case/internal/external-api/rate"
	"se-school-case/internal/mails"
	"se-school-case/internal/rates"
	"se-school-case/internal/subscribers"
	"se-school-case/pkg/constants"
)

type dependencies struct {
	subscriberService subscribers.SubscriberInterface
	rateService       rates.RateInterface
	mailService       mails.MailInterface
	cronService       cronjobs.CronJobsInterface
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
	cronService := cronjobs.NewService(&mailService)
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
