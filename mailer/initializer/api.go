package initializer

type Api struct {
	deps dependencies
}

func NewApi() *Api {
	deps := wireDependencies()
	return &Api{deps: *deps}
}

func (a *Api) StartApplication() {
	a.deps.cronJobsMailer.StartScheduler()
	a.deps.eventConsumer.ConsumeEvents()
}
