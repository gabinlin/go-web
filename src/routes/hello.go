package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Include(func(engine *gin.Engine) {
		engine.GET("/test", func(context *gin.Context) {
			context.String(http.StatusOK, "test")
		})
	})
}
