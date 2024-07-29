package initializer

import (
	"github.com/gin-gonic/gin"
	cronjobs "se-school-case/internal/cron-jobs"
	rateshandler "se-school-case/internal/rate/handler"
	subhandler "se-school-case/internal/subscriber"
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
	rateshandler.NewHandler(deps.rateService).Register(engine)
	subhandler.NewHandler(deps.subscriberService, deps.rabbitMQ).Register(engine)
	cronjobs.NewHandler(deps.cronService).Register(engine)
	return &api{engine}
}

func (a *api) HandleRequests() {
	StartServer(a.router)
}
