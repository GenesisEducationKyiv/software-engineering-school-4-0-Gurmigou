package initializer

import (
	"github.com/gin-gonic/gin"
	"os"
)

func StartServer(r *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	// Start the server on the specified port
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
