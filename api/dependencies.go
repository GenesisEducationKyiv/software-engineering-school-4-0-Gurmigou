package api

import (
	"gorm.io/gorm"
	cron_jobs "se-school-case/pkg/domain/cron-jobs"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/domain/rate"
	"se-school-case/pkg/domain/subscriber"
	"se-school-case/pkg/initializer"
	"se-school-case/pkg/util/constants"
)

type dependencies struct {
	subscriberService subscriber.Service
	rateService       rate.Service
	mailService       mail.Service
	cronService       cron_jobs.Service
}

func wireDependencies() *dependencies {
	initEnv()
	db := connectToDb()
	rateService := rate.NewService(db)
	subscriberService := subscriber.NewService(db)
	mailService := mail.NewService(subscriberService, rateService)
	cronService := cron_jobs.NewService(mailService)
	return &dependencies{
		subscriberService: subscriberService,
		rateService:       rateService,
		mailService:       mailService,
		cronService:       cronService,
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
