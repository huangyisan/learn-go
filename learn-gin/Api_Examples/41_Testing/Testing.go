package main

import "github.com/gin-gonic/gin"

// 官方推荐net/http/httptest包进行测试


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":7777")
}
