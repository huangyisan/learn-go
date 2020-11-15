package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 自定义log，所以不使用Default()来创建实例， Default()创建的实例是包含了Logger这个middleware的。

	router := gin.New()
	// 自定义日志输出格式， 使用LoggerWithFormatter
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// 该匿名函数返回为string，所以构造出string进行返回。
		// params有哪些属性，可以点击进入查看其结构体定义。
		return fmt.Sprintf("IP: %s, Method: %s \n", params.ClientIP, params.Method)
	}))

	router.Use(gin.Recovery())

	router.GET("/logformat", func(c *gin.Context) {
		c.String(200, "LogFormat")
	})

	router.Run(":7777")
}
