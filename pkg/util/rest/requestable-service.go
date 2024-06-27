package rest

import "github.com/gin-gonic/gin"

type Requestable interface {
	Request(c *gin.Context, e *gin.Engine)
}
