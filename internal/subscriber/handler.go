package subscriber

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	cronjobs "se-school-case/internal/cron-jobs"
	"se-school-case/pkg/queue"
	"strconv"
	"time"
)

type SubscriberInterface interface {
	Exists(email string) (bool, error)
}

type Handler struct {
	subscriberService SubscriberInterface
	rabbitMQ          queue.RabbitMQ
}

func NewHandler(subscriberService SubscriberInterface, mq queue.RabbitMQ) *Handler {
	return &Handler{subscriberService, mq}
}

func (h *Handler) Register(engine *gin.Engine) {
	engine.POST("/api/subscribe", h.AddUserEmail)
	engine.POST("/api/unsubscribe", h.DeleteUserEmail)
}

// swagger:route POST /api/subscribe Subscriber addUserEmail
// Add a new subscriber email
//
// Adds a new email to the list of subscribers.
//
// responses:
//
//	200: body:gin.H{"message": "Email added successfully"}
//	400: body:gin.H{"errors": "Email request body is not correct."}
//	409: body:gin.H{"errors": "Email already exists"}
//	500: body:gin.H{"errors": "Failed to add email"}
func (h *Handler) AddUserEmail(context *gin.Context) {
	var input EmailDto

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": "Email request body is not correct."})
		return
	}

	exists, err := h.subscriberService.Exists(input.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to check email"})
		return
	}
	if exists {
		context.JSON(http.StatusConflict, gin.H{"errors": "Email already exists"})
		return
	}

	event := cronjobs.Event{
		EventID:     strconv.Itoa(rand.Intn(9999)),
		EventType:   cronjobs.Subscribe,
		AggregateID: "sub-1",
		Timestamp:   time.Now().Format(time.RFC3339),
		Data: cronjobs.EventData{
			Email: input.Email,
		},
	}

	message, err := json.Marshal(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to marshal event"})
		return
	}

	err = h.rabbitMQ.Publish(string(message))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to publish message"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Email added successfully"})
}

// swagger:route DELETE /api/unsubscribe Subscriber deleteUserEmail
// Delete a subscriber email
//
// Removes an email from the list of subscribers.
//
// responses:
//
//	200: body:gin.H{"message": "Email delete request received successfully"}
//	400: body:gin.H{"errors": "Email request body is not correct."}
//	404: body:gin.H{"errors": "Email not found"}
//	500: body:gin.H{"errors": "Failed to process request"}
func (h *Handler) DeleteUserEmail(context *gin.Context) {
	var input EmailDto

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": "Email request body is not correct."})
		return
	}

	exists, err := h.subscriberService.Exists(input.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to check email"})
		return
	}
	if !exists {
		context.JSON(http.StatusNotFound, gin.H{"errors": "Email not found"})
		return
	}

	event := cronjobs.Event{
		EventID:     strconv.Itoa(rand.Intn(9999)),
		EventType:   cronjobs.Unsubscribe,
		AggregateID: "sub-1",
		Timestamp:   time.Now().Format(time.RFC3339),
		Data: cronjobs.EventData{
			Email: input.Email,
		},
	}

	message, err := json.Marshal(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to marshal event"})
		return
	}

	err = h.rabbitMQ.Publish(string(message))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to publish message"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Email delete request received successfully"})
}
