package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// 1、静态文件夹
	engine.Static("/static", "./src/static")
	// 2、静态文件
	engine.StaticFile("/file/1", "./src/static/1.jpeg")
	// 3、类似于静态文件夹，但是更定制化
	engine.StaticFS("/fs", http.Dir("./src/static"))
	group := engine.Group("/api/v1")
	{
		group.GET("/:arg1", func(context *gin.Context) {
			context.String(http.StatusOK, context.Param("arg1"))
		})
		group.GET("/special/*action", func(context *gin.Context) {
			context.String(http.StatusOK, context.Request.URL.RequestURI())
		})
	}
	engine.Run()
}
