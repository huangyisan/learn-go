package middleware

import (
	"github.com/gin-gonic/gin"
	. "jwt-todo/handler"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			// 不通过则结束,不在处理
			c.Abort()
			return
		}
		c.Next()
	}
}
