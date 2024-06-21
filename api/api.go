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
	engine := gin.Default()
	deps := wireDependencies()
	deps.cronService.StartScheduler()
	rate.NewHandler(deps.rateService).Register(engine)
	subscriber.NewHandler(deps.subscriberService).Register(engine)
	cronjobs.NewHandler(deps.mailService).Register(engine)
	return &api{engine}
}

func (a *api) HandleRequests() {
	initializer.StartServer(a.router)
}
