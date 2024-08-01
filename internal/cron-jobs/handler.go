package cron_jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/internal"
)

type CronJobsInterface interface {
	StartScheduler()
	NotifyAboutExchangeRate() error
	ExplicitlyNotify() error
}

type CronJobsHandler struct {
	cronJobsService CronJobsInterface
}

func NewHandler(cronJobsService CronJobsInterface) internal.Registrable {
	return &CronJobsHandler{cronJobsService: cronJobsService}
}

func (h *CronJobsHandler) Register(engine *gin.Engine) {
	engine.POST("/api/notify", h.ExplicitlyNotify)
}

// swagger:route POST /api/notify CronJobs postExplicitlyNotify
// Explicitly notify all subscribers
//
// Sends an email notification to all subscribers about the exchange rate.
//
// responses:
//
//	200: body:gin.H{"message": "Successfully notified all users."}
func (h *CronJobsHandler) ExplicitlyNotify(context *gin.Context) {
	err := h.cronJobsService.ExplicitlyNotify()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to notify users - " + err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
	}
}
