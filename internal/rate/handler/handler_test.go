package handler_test

import (
	"net/http"
	"net/http/httptest"
	"se-school-case/internal/rate/handler"
	"se-school-case/internal/rate/handler/mock"
	"se-school-case/pkg/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetExchangeRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRateService := mock.NewMockRateInterface(ctrl)
	rate := model.Rate{Rate: 27.5}
	mockRateService.EXPECT().GetRate().Return(rate, nil)

	h := handler.NewHandler(mockRateService)
	router := gin.Default()
	h.Register(router)

	req, err := http.NewRequest(http.MethodGet, "/api/rates", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "27.500000", rr.Body.String())
}

func TestHandler_GetExchangeRate_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRateService := mock.NewMockRateInterface(ctrl)
	mockRateService.EXPECT().GetRate().Return(model.Rate{}, assert.AnError)

	h := handler.NewHandler(mockRateService)
	router := gin.Default()
	h.Register(router)

	req, err := http.NewRequest(http.MethodGet, "/api/rates", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"errors":"Failed to get the latest rate"}`, rr.Body.String())
}
