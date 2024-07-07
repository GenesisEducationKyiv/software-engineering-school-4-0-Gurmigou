package handler

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

func NewHandler(rateService RateInterface) *Handler {
	return &Handler{rateService}
}

func (h *Handler) Register(engine *gin.Engine) {
	engine.GET("/api/rates", h.GetExchangeRate)
}

// swagger:route GET /api/rates Rate getExchangeRate
// Get the latest exchange rates
//
// responses:
//
//	200: body:string The latest exchange rates
//	500: body:gin.H{"errors": "Failed to get the latest rates"}
func (h *Handler) GetExchangeRate(context *gin.Context) {
	rateResp, err := h.rateService.GetRate()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to get the latest rates"})
		return
	}
	context.String(http.StatusOK, "%f", rateResp.Rate)
}
