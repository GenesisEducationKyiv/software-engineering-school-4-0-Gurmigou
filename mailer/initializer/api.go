package initializer

type Api struct {
	deps dependencies
}

func NewApi() *Api {
	deps := wireDependencies()
	return &Api{deps: *deps}
}

func (a *Api) StartConsumer() {
	a.deps.cronJobsConsumer.ConsumeEvents()
}
