package initializer

import (
	"gorm.io/gorm"
	"se-school-case/db"
	"se-school-case/infra/external-api/rate/provider"
	cronjobsservice "se-school-case/internal/cron-jobs"
	ratesrepo "se-school-case/internal/rate"
	rateshandler "se-school-case/internal/rate/handler"
	subservice "se-school-case/internal/subscriber"
	"se-school-case/pkg/constants"
	"se-school-case/pkg/queue"
)

type dependencies struct {
	subscriberService subservice.SubscriberInterface
	rateService       rateshandler.RateInterface
	cronService       cronjobsservice.CronJobsInterface
	rabbitMQ          queue.RabbitMQ
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

	rateService := ratesrepo.NewService(&rateRepository, &bankFetchService)
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
