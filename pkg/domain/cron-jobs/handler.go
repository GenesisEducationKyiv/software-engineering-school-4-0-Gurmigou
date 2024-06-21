package cron_jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/util/constants"
)

type CronJobsInterface interface {
	StartScheduler()
}

type CronJobsHandler struct {
	mailService MailInterface
}

func NewHandler(router *gin.Engine, mailService MailInterface) *CronJobsHandler {
	ctrl := &CronJobsHandler{mailService: mailService}
	router.POST("/api/notify", ctrl.PostExplicitlyNotify)
	return ctrl
}

func (c *CronJobsHandler) PostExplicitlyNotify(context *gin.Context) {
	c.mailService.SendEmailToAll("Exchange rate notification", constants.TEMPLATE_PATH)
	context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
