package internal

import "github.com/gin-gonic/gin"

type Registrable interface {
	Register(engine *gin.Engine)
}
