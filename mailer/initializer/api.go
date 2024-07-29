package initializer

type Api struct {
	deps dependencies
}

func NewApi() *Api {
	deps := wireDependencies()
	return &Api{deps: *deps}
}

func (a *Api) StartApplication() {
	a.deps.eventConsumer.ConsumeEvents()
	a.deps.cronJobsMailer.StartScheduler()
}
