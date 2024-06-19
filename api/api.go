package api

import (
	"github.com/gin-gonic/gin"
	cronjobs "se-school-case/pkg/domain/cron-jobs"
	"se-school-case/pkg/domain/rate"
	"se-school-case/pkg/domain/subscriber"
	"se-school-case/pkg/initializer"
)

type Api interface {
	HandleRequests()
}

type api struct {
	router *gin.Engine
}

func NewApi() Api {
	router := gin.Default()
	deps := wireDependencies()
	deps.cronService.StartScheduler()
	rate.NewController(router, deps.rateService)
	subscriber.NewController(router, deps.subscriberService)
	cronjobs.NewController(router, deps.mailService)
	return &api{router}
}

func (a *api) HandleRequests() {
	initializer.StartServer(a.router)
}
