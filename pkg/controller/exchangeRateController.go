package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/service"
)

func GetExchangeRate(c *gin.Context) {
	rate, err := service.GetRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the latest rate"})
		return
	}

	c.String(http.StatusOK, "%f", rate.Rate)
}
