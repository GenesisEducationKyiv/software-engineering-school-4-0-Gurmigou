package cron_jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain"
	"se-school-case/pkg/util/constants"
)

type CronJobsInterface interface {
	StartScheduler()
}

type CronJobsHandler struct {
	mailService MailInterface
}

func NewHandler(mailService MailInterface) domain.Registrable {
	return &CronJobsHandler{mailService: mailService}
}

func (h *CronJobsHandler) Register(engine *gin.Engine) {
	engine.POST("/api/notify", h.ExplicitlyNotify)
}

// swagger:route POST /api/notify CronJobs postExplicitlyNotify
// Explicitly notify all subscribers
//
// Sends an email notification to all subscribers about the exchange rates.
//
// responses:
//
//	200: body:gin.H{"message": "Successfully notified all users."}
func (h *CronJobsHandler) ExplicitlyNotify(context *gin.Context) {
	err := h.mailService.SendEmailToAll("Exchange rates notification", constants.TEMPLATE_PATH)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to notify users - " + err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
	}
}
