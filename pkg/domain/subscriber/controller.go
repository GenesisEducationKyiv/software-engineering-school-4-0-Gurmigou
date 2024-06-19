package subscriber

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/util/app-error"
)

type Controller struct {
	subscriberService Service
}

func NewController(router *gin.Engine, subscriberService Service) *Controller {
	ctrl := &Controller{subscriberService}
	router.POST("/api/subscribe", ctrl.AddUserEmail)
	return ctrl
}

func (c *Controller) AddUserEmail(context *gin.Context) {
	var input mail.EmailDto

	// Bind input
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"app-error": "Email request body is not correct."})
		return
	}

	// Handle email subscription in service layer
	if err := c.subscriberService.Add(input.Email); err != nil {
		if errors.Is(err, app_errors.ErrEmailAlreadyExists) {
			context.JSON(http.StatusConflict, gin.H{"app-error": "Email already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"app-error": "Failed to add email"})
		}
		return
	}

	// Email successfully added
	context.JSON(http.StatusOK, gin.H{"message": "Email added successfully"})
}
