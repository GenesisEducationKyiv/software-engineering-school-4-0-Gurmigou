package rate

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	rateService Service
}

func NewController(router *gin.Engine, rateService Service) *Controller {
	ctrl := &Controller{rateService}
	router.GET("/api/rate", ctrl.GetExchangeRate)
	return ctrl
}

func (c *Controller) GetExchangeRate(context *gin.Context) {
	rateResp, err := c.rateService.GetRate()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"app-error": "Failed to get the latest rate"})
		return
	}

	context.String(http.StatusOK, "%f", rateResp.Rate)
}
