package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func DummyMiddleware() gin.HandlerFunc {
	// 这种方式定义的中间件，只在加载的时候执行return前面的内容， 所以t为一个固定的值，即便你请求多次。
	// 这种方法常用来初始化实例的时候使用。
	t:= time.Now()
	//fmt.Println("time is",t)

	return func(c *gin.Context) {
		fmt.Println("time is",t)
		c.Next()
	}
}

func main() {

	r := gin.Default()
	
	r.Use(DummyMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
	// ...
	r.Run(":7777")
}
