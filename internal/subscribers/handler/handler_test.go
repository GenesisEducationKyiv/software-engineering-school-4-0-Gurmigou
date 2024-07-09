package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	suberrors "se-school-case/internal/subscribers/errors"
	"se-school-case/internal/subscribers/handler/mock"
	"se-school-case/internal/subscribers/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddUserEmail_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock.NewMockSubscriberInterface(mockCtrl)
	mockService.EXPECT().Add("test@example.com").Return(nil).Times(1)

	handler := &Handler{subscriberService: mockService}
	router := gin.Default()
	router.POST("/api/subscribe", handler.AddUserEmail)

	requestBody, err := json.Marshal(map[string]string{
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/subscribe", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expectedResponse := `{"message": "Email added successfully"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())
}

func TestAddUserEmail_EmailExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock.NewMockSubscriberInterface(mockCtrl)
	mockService.EXPECT().Add("test@example.com").Return(suberrors.ErrEmailAlreadyExists).Times(1)

	controller := Handler{subscriberService: mockService}
	router := gin.Default()
	router.POST("/api/subscribe", controller.AddUserEmail)

	// Corrected request body to match the expected format
	requestBody, err := json.Marshal(map[string]string{
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/subscribe", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)
	expectedResponse := `{"errors": "Email already exists"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())
}

func TestAddUserEmail_InvalidEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	controller := &Handler{}
	router := gin.Default()
	router.POST("/api/subscribe", controller.AddUserEmail)

	requestBody, err := json.Marshal(model.EmailDto{Email: "invalid-email"})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/subscribe", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	expectedResponse := `{"errors": "Email request body is not correct."}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())
}

func TestAddUserEmail_InternalServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock.NewMockSubscriberInterface(mockCtrl)
	mockService.EXPECT().Add("test@example.com").Return(errors.New("internal error")).Times(1)

	handler := &Handler{subscriberService: mockService}
	router := gin.Default()
	router.POST("/api/subscribe", handler.AddUserEmail)

	requestBody, err := json.Marshal(map[string]string{
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/subscribe", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	expectedResponse := `{"errors": "Failed to add email"}`
	assert.JSONEq(t, expectedResponse, rr.Body.String())
}
