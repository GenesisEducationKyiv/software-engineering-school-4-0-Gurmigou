package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain/mail"
	"se-school-case/pkg/domain/user"
)

func PostAddUserEmail(c *gin.Context) {
	var input mail.EmailDto

	// Bind input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email request body is not correct."})
		return
	}

	// Handle email subscription in service layer
	if err := user.AddUserSubscription(input.Email); err != nil {
		if errors.Is(err, user.ErrEmailAlreadyExists) {
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
	mail.SendEmailNotificationsToAll()
	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified all users."})
}
