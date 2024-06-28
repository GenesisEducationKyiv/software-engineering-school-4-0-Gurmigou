package initializer

import (
	"github.com/gin-gonic/gin"
	"se-school-case/pkg/constants"
)

func StartServer(r *gin.Engine) {
	port := constants.PORT
	if port == "" {
		port = "8080" // Default port if not specified
	}
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
