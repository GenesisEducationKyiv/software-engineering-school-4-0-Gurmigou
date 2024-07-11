package initializer

import (
	"gorm.io/gorm"
	"se-school-case/db"
	"se-school-case/infra/external-api/rate/provider"
	cronjobs "se-school-case/internal/cron-jobs"
	rateshandler "se-school-case/internal/rate/handler"
	ratesrepo "se-school-case/internal/rate/repo"
	ratesservice "se-school-case/internal/rate/service"
	subhandler "se-school-case/internal/subscriber/handler"
	subrepo "se-school-case/internal/subscriber/repo"
	subservice "se-school-case/internal/subscriber/service"
	queue "se-school-case/mailer/pkg"
	"se-school-case/pkg/constants"
)

var ()

type dependencies struct {
	subscriberService subhandler.SubscriberInterface
	rateService       rateshandler.RateInterface
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
	rabbitMq, _ := queue.NewRabbitMQConnection(constants.RABBIT_MQ_URL, "mailer_events")

	rateService := ratesservice.NewService(&rateRepository, &bankFetchService)
	subscriberService := subservice.NewService(&subscriberRepository)
	cronService := cronjobs.NewService(rabbitMq, &subscriberService, &rateService)
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
