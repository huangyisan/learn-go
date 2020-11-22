package user

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"auth",
	})
}
