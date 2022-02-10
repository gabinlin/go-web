package routes

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init(engine *gin.Engine) {
	for _, option := range options {
		option(engine)
	}
}
