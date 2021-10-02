package main

import (
	"ginEssential/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/auth/login", controller.UserLogin)
	// 每次请求接口的时候，都先调用中间件进行处理
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
