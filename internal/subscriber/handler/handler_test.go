package handler_test

//
//import (
//	"bytes"
//	"errors"
//	"net/http"
//	"net/http/httptest"
//	"se-school-case/tests/cron-jobs/mocks"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestAddUserEmail(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockSubscriberService := mocks.NewMockSubscriberInterface(ctrl)
//	mockRabbitMQ := mock_queue.NewMockRabbitMQ(ctrl)
//
//	h := handler.NewHandler(mockSubscriberService, mockRabbitMQ)
//	router := gin.Default()
//	h.Register(router)
//
//	t.Run("successful subscription", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(false, nil)
//		mockRabbitMQ.EXPECT().Publish(gomock.Any()).Return(nil)
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("POST", "/api/subscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusOK, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email added successfully")
//	})
//
//	t.Run("invalid email format", func(t *testing.T) {
//		reqBody := `{"email": "invalid-email"}`
//		req, _ := http.NewRequest("POST", "/api/subscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusBadRequest, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email request body is not correct.")
//	})
//
//	t.Run("email already exists", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(true, nil)
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("POST", "/api/subscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusConflict, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email already exists")
//	})
//
//	t.Run("error checking email existence", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(false, errors.New("error"))
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("POST", "/api/subscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusInternalServerError, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Failed to check email")
//	})
//
//	t.Run("error publishing message to queue", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(false, nil)
//		mockRabbitMQ.EXPECT().Publish(gomock.Any()).Return(errors.New("error"))
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("POST", "/api/subscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusInternalServerError, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Failed to publish message")
//	})
//}
//
//func TestDeleteUserEmail(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockSubscriberService := mock_subscriber.NewMockSubscriberInterface(ctrl)
//	mockRabbitMQ := mock_queue.NewMockRabbitMQ(ctrl)
//
//	h := handler.NewHandler(mockSubscriberService, mockRabbitMQ)
//	router := gin.Default()
//	h.Register(router)
//
//	t.Run("successful unsubscription", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(true, nil)
//		mockRabbitMQ.EXPECT().Publish(gomock.Any()).Return(nil)
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("DELETE", "/api/unsubscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusOK, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email delete request received successfully")
//	})
//
//	t.Run("invalid email format", func(t *testing.T) {
//		reqBody := `{"email": "invalid-email"}`
//		req, _ := http.NewRequest("DELETE", "/api/unsubscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusBadRequest, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email request body is not correct.")
//	})
//
//	t.Run("email not found", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(false, nil)
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("DELETE", "/api/unsubscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusNotFound, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Email not found")
//	})
//
//	t.Run("error checking email existence", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(false, errors.New("error"))
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("DELETE", "/api/unsubscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusInternalServerError, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Failed to check email")
//	})
//
//	t.Run("error publishing message to queue", func(t *testing.T) {
//		mockSubscriberService.EXPECT().Exists("test@example.com").Return(true, nil)
//		mockRabbitMQ.EXPECT().Publish(gomock.Any()).Return(errors.New("error"))
//
//		reqBody := `{"email": "test@example.com"}`
//		req, _ := http.NewRequest("DELETE", "/api/unsubscribe", bytes.NewBufferString(reqBody))
//		req.Header.Set("Content-Type", "application/json")
//		resp := httptest.NewRecorder()
//
//		router.ServeHTTP(resp, req)
//
//		assert.Equal(t, http.StatusInternalServerError, resp.Code)
//		assert.Contains(t, resp.Body.String(), "Failed to publish message")
//	})
//}
