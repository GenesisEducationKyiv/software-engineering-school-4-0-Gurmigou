package subscriber

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"se-school-case/pkg/domain"
	"se-school-case/pkg/model"
	"se-school-case/pkg/util/app-error"
)

type SubscriberInterface interface {
	Add(email string) error
	GetAll() ([]model.User, error)
}

type Handler struct {
	subscriberService SubscriberInterface
}

func NewHandler(subscriberService SubscriberInterface) domain.Registrable {
	return &Handler{subscriberService}
}

func (h *Handler) Register(engine *gin.Engine) {
	engine.POST("/api/subscribe", h.AddUserEmail)
}

func (h *Handler) AddUserEmail(context *gin.Context) {
	var input EmailDto

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"app-error": "Email request body is not correct."})
		return
	}

	if err := h.subscriberService.Add(input.Email); err != nil {
		if errors.Is(err, app_errors.ErrEmailAlreadyExists) {
			context.JSON(http.StatusConflict, gin.H{"app-error": "Email already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"app-error": "Failed to add email"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Email added successfully"})
}
