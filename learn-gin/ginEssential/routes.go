package main

import (
	"ginEssential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/auth/login", controller.UserLogin)
	return r
}