package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/dto"
	"se-school-case/service"
)

func PostAddUserEmail(c *gin.Context) {
	var input dto.EmailDto

	// Bind input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Handle email subscription in service layer
	if err := service.HandleEmailSubscription(input.Email); err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
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
	service.SendEmailNotificationsToAll()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
