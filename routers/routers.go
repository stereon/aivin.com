package routers

import (

	"github.com/gin-gonic/gin"
	"github.com/stereon/aivin.com/controller"
	"github.com/stereon/aivin.com/middleware"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)
	return r
}