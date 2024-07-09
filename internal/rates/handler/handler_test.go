package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"se-school-case/internal/rates/handler/mock"
	"se-school-case/pkg/model"
	"testing"
)

func TestHandler_GetExchangeRate_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRateService := mock.NewMockRateInterface(mockCtrl)
	mockRateService.EXPECT().GetRate().Return(model.Rate{Rate: 37.07}, nil).Times(1)

	handler := NewHandler(mockRateService)
	router := gin.Default()
	router.GET("/api/rates", handler.GetExchangeRate)

	// Act
	req, err := http.NewRequest("GET", "/api/rates", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "37.070000", rr.Body.String())
}

func TestHandler_GetExchangeRate_Failure(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRateService := mock.NewMockRateInterface(mockCtrl)
	mockRateService.EXPECT().GetRate().Return(model.Rate{}, errors.New("service error")).Times(1)

	handler := NewHandler(mockRateService)
	router := gin.Default()
	router.GET("/api/rates", handler.GetExchangeRate)

	req, err := http.NewRequest("GET", "/api/rates", nil)
	assert.NoError(t, err)

	// Act
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"errors": "Failed to get the latest rates"}`, rr.Body.String())
}
