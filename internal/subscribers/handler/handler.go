package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	suberrors "se-school-case/internal/subscribers/errors"
	submodel "se-school-case/internal/subscribers/model"
	"se-school-case/pkg/model"
)

type SubscriberInterface interface {
	Add(email string) error
	GetAll() ([]model.User, error)
}

type Handler struct {
	subscriberService SubscriberInterface
}

func NewHandler(subscriberService SubscriberInterface) *Handler {
	return &Handler{subscriberService}
}

func (h *Handler) Register(engine *gin.Engine) {
	engine.POST("/api/subscribe", h.AddUserEmail)
}

// swagger:route POST /api/subscribe Subscriber addUserEmail
// Add a new subscribers email
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
	var input submodel.EmailDto

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": "Email request body is not correct."})
		return
	}

	if err := h.subscriberService.Add(input.Email); err != nil {
		if errors.Is(err, suberrors.ErrEmailAlreadyExists) {
			context.JSON(http.StatusConflict, gin.H{"errors": "Email already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"errors": "Failed to add email"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Email added successfully"})
}
