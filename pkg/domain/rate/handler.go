package rate

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/model"
)

type RateInterface interface {
	GetRate() (model.Rate, error)
	SaveRate(currencyFrom string, currencyTo string, exchangeRate float64)
}

type RateFetchInterface interface {
	FetchExchangeRate() (float64, error)
}

type Handler struct {
	rateService RateInterface
}

func NewHandler(router *gin.Engine, rateService RateInterface) *Handler {
	ctrl := &Handler{rateService}
	router.GET("/api/rate", ctrl.GetExchangeRate)
	return ctrl
}

func (c *Handler) GetExchangeRate(context *gin.Context) {
	rateResp, err := c.rateService.GetRate()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"app-error": "Failed to get the latest rate"})
		return
	}

	context.String(http.StatusOK, "%f", rateResp.Rate)
}
