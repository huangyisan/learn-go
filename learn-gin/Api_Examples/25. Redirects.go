package main

import "github.com/gin-gonic/gin"

// gin支持内部、外部重定向

func main () {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		// 外部重定向
		c.Redirect(301, "/foo")
	})
	router.GET("/foo", func(c *gin.Context) {
		c.String(200, "foo")
	})



	router.GET("/test1", func(c *gin.Context) {
		// 改写了url
		c.Request.URL.Path = "/test2"
		//HandleContext 对改写后的c进行处理, 内部重定向
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})


	router.Run(":7777")
}