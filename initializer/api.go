package initializer

import (
	"github.com/gin-gonic/gin"
	cronjobs "se-school-case/internal/cron-jobs"
	"se-school-case/internal/rates"
	"se-school-case/internal/subscribers"
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
	rates.NewHandler(deps.rateService).Register(engine)
	subscribers.NewHandler(deps.subscriberService).Register(engine)
	cronjobs.NewHandler(deps.mailService).Register(engine)
	return &api{engine}
}

func (a *api) HandleRequests() {
	StartServer(a.router)
}
