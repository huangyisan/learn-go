package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

// 如果要优雅的关闭或者重启服务，有一些方法可以使用。
// 采用第三方包，或者用内建方法来实现。

// 第三方包的方式endless， 这个包的部分信号win不支持，建议在linux上验证使用

func main() {
	router := gin.Default()
	router.GET("/shoutdown", func(c *gin.Context) {
		c.String(200, "shutdown")
	})


	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "world")
	})


	time.Sleep(5 * time.Second)
	endless.ListenAndServe(":4242", router)
}