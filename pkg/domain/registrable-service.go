package domain

import "github.com/gin-gonic/gin"

type Registrable interface {
	Register(engine *gin.Engine)
}
