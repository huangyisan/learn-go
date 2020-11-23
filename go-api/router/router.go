package router

import (
	"github.com/gin-gonic/gin"
)
import "goapi/handler/user"


func Router(r *gin.Engine) *gin.Engine {
	r.POST("/api/v1/register", user.Register)
	r.POST("/api/v1/auth", user.Auth)
	return r
}