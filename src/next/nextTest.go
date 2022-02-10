package main

import "github.com/gin-gonic/gin"

type HandlerFunc func(ctx *gin.Context)

type HandlersChain []HandlerFunc

func main() {
	var handlers = make(HandlersChain, 10)
	handlers = append(handlers, func(ctx *gin.Context) {

	})
}
