package main

import "github.com/gin-gonic/gin"

// 路由组定义，当多个路由有同一个父级路由的时候，则可以使用。
// 比如restful api分隔的会给/v1/  /v2/之类的路由

func main() {
	router := gin.Default()
	// 标记父类路由为/v1
	v1 := router.Group("v1")
	// 定义了v1下面的相关内容
	// 这边无端端的大括号会产生疑问，大括号里的内容作为单独的语句块，其里面的变量是单独作用域。
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, "/login")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(200, "/login")
		})
	}

	v2 := router.Group("v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(200, "Get /login")
		})
	}

	router.Run(":7777")
}