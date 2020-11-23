package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpcode int, code int, content string, msg string) {
	c.JSON(httpcode, gin.H{
		"code" : code,
		"content": content,
		"message": msg,
	})
}

func Success(c *gin.Context, content string, msg string) {
	Response(c, 200, 200, content, msg)
}

// 400
func Fail(c *gin.Context, content string, msg string) {
	Response(c, 400, 400, content, msg)
}

	
