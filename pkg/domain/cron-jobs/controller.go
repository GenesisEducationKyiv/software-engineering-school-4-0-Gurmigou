package cron_jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/util/constants"
)

type Controller struct {
	mailService mail.Service
}

func NewController(router *gin.Engine, mailService mail.Service) *Controller {
	ctrl := &Controller{mailService: mailService}
	router.POST("/api/notify", ctrl.PostExplicitlyNotify)
	return ctrl
}

func (c *Controller) PostExplicitlyNotify(context *gin.Context) {
	c.mailService.SendEmailToAll("Exchange rate notification", constants.TEMPLATE_PATH)
	context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
