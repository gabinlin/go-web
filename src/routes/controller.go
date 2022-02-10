package routes

import "github.com/gin-gonic/gin"

type Controller interface {
	Registry(e *gin.Engine)
}
