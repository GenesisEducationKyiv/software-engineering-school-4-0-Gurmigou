package main

import (
	"github.com/gin-gonic/gin"
	"se-school-case/controller"
	"se-school-case/initializer"
	"se-school-case/service"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
	initializer.AutoMigrateDatabase()
	service.StartRateUpdater()
}

func main() {
	r := gin.Default()

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/api/subscribe", controller.PostAddUserEmail)
	r.GET("/api/rate", controller.GetExchangeRate)

	initializer.StartServer(r)
}
