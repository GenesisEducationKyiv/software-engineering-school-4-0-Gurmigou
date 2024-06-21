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
	engine.POST("/api/notify", h.PostExplicitlyNotify)
}

func (h *CronJobsHandler) PostExplicitlyNotify(context *gin.Context) {
	h.mailService.SendEmailToAll("Exchange rate notification", constants.TEMPLATE_PATH)
	context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}