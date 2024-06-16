package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain/rate"
)

func GetExchangeRate(c *gin.Context) {
	rate_, err := rate.GetRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the latest rate"})
		return
	}

	c.String(http.StatusOK, "%f", rate_.Rate)
}
