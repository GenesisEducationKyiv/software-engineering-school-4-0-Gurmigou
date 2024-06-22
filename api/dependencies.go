package api

import (
	"gorm.io/gorm"
	cron_jobs "se-school-case/pkg/domain/cron-jobs"
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
	initEnv()
	db := connectToDb()
	rateRepo := rates.NewRateRepository(db)
	fetchService := rates.NewRateFetchService()
	rateService := rates.NewService(rateRepo, &fetchService)
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

func initEnv() {
	initializer.LoadEnvVariables()
	constants.InitEnvValues()
}

func connectToDb() *gorm.DB {
	initializer.RunMigrations()
	db := initializer.ConnectToDatabase()
	return db
}
