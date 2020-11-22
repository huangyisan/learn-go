package router

import "github.com/gin-gonic/gin"
import "goapi/handler/user"


func Router(r *gin.Engine) *gin.Engine {
	r.POST("/v1/register", user.Register)
	r.POST("/v1/auth", user.Auth)
	return r
}