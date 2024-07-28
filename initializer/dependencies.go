package initializer

import (
	"gorm.io/gorm"
	"se-school-case/db"
	"se-school-case/infra/external-api/rate/provider"
	cronjobsservice "se-school-case/internal/cron-jobs"
	rateshandler "se-school-case/internal/rate/handler"
	ratesrepo "se-school-case/internal/rate/repo"
	ratesservice "se-school-case/internal/rate/service"
	subservice "se-school-case/internal/subscriber"
	subhandler "se-school-case/internal/subscriber/handler"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/queue"
)

type dependencies struct {
	subscriberService subhandler.SubscriberInterface
	rateService       rateshandler.RateInterface
	cronService       cronjobsservice.CronJobsInterface
}

func wireDependencies() *dependencies {
	InitEnv()
	db := setUpDatabase()

	// Initialize repositories
	rateRepository := ratesrepo.NewRateRepository(db)
	subscriberRepository := subservice.NewSubscriberRepository(db)

	// Initialize chain of Rate fetchers
	bankFetchService := provider.NewBankRateFetchService()
	exchangeFetchService := provider.NewExchangeApiRateFetch()
	bankFetchService.SetNext(&exchangeFetchService)
	rabbitMq, _ := queue.NewRabbitMQConnection(constants.RABBITMQ_URL, constants.QUEUE_NAME)

	rateService := ratesservice.NewService(&rateRepository, &bankFetchService)
	subscriberService := subservice.NewService(&subscriberRepository)
	cronService := cronjobsservice.NewService(rabbitMq, &subscriberService, &rateService)
	return &dependencies{
		subscriberService: &subscriberService,
		rateService:       &rateService,
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
