package cron_jobs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"se-school-case/pkg/domain/mail"
)

// Controller defines the structure for the controller
type Controller struct {
	mailService mail.Service
}

// NewController initializes the routes and returns a controller instance
func NewController(router *gin.Engine, mailService mail.Service) *Controller {
	ctrl := &Controller{mailService: mailService}
	router.POST("/api/notify", ctrl.PostExplicitlyNotify)
	return ctrl
}

func (c *Controller) PostExplicitlyNotify(context *gin.Context) {
	c.mailService.SendEmailToAll("Exchange rate notification", os.Getenv("TEMPLATE_PATH"))
	context.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
