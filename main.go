package main

import (
	"github.com/gin-gonic/gin"
	controller2 "se-school-case/pkg/controller"
	initializer2 "se-school-case/pkg/initializer"
	"se-school-case/pkg/service"
)

func init() {
	initializer2.LoadEnvVariables()
	initializer2.ConnectToDatabase()
	initializer2.RunMigrations()
	service.StartScheduledEmail()
}

func main() {
	r := gin.Default()

	/*
		Util end-points:
		/api/ping - ping-pong server
		/api/notify - explicitly notifies all users using email without schedules interval
	*/
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/api/notify", controller2.PostExplicitlyNotify)

	// Required end-points
	r.POST("/api/subscribe", controller2.PostAddUserEmail)
	r.GET("/api/rate", controller2.GetExchangeRate)

	initializer2.StartServer(r)
}
