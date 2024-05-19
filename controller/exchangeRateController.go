package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/service"
)

func GetExchangeRate(c *gin.Context) {
	rate, err := service.GetRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the latest rate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate.Rate})
}
