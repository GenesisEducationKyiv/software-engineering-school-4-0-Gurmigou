package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/dto"
	service2 "se-school-case/pkg/service"
)

func PostAddUserEmail(c *gin.Context) {
	var input dto.EmailDto

	// Bind input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email request body is not correct."})
		return
	}

	// Handle email subscription in service layer
	if err := service2.HandleEmailSubscription(input.Email); err != nil {
		if errors.Is(err, service2.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add email"})
		}
		return
	}

	// Email successfully added
	c.JSON(http.StatusOK, gin.H{"message": "Email added successfully"})
}

func PostExplicitlyNotify(c *gin.Context) {
	service2.SendEmailNotificationsToAll()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
