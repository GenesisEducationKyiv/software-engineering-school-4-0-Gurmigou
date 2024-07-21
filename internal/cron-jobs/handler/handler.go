package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/internal"
	"se-school-case/internal/cron-jobs/service"
	"se-school-case/pkg/constants"
)

type CronJobsInterface interface {
	StartScheduler()
}

type CronJobsHandler struct {
	mailService service.MailInterface
}

func NewHandler(mailService service.MailInterface) internal.Registrable {
	return &CronJobsHandler{mailService: mailService}
}

func (h *CronJobsHandler) Register(engine *gin.Engine) {
	engine.POST("/api/notify", h.ExplicitlyNotify)
}

// swagger:route POST /api/notify CronJobs postExplicitlyNotify
// Explicitly notify all subscriber
//
// Sends an email notification to all subscriber about the exchange rate.
//
// responses:
//
//	200: body:gin.H{"message": "Successfully notified all users."}
func (h *CronJobsHandler) ExplicitlyNotify(context *gin.Context) {
	err := h.mailService.SendEmailToAll("Exchange rate notification", constants.TEMPLATE_PATH)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to notify users - " + err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
	}
}
