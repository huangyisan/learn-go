package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个不带middleware的实例
	r := gin.New()

	// 注册middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 对于每个请求，都可以为其添加中间件,一个或者多个
	//r.GET("/benchmark",middleware1,middleware2 )
	r.Run(":7777")


}