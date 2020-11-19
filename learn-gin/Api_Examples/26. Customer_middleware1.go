package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 这种方式定义的中间件，请求一次会执行一次，每次得到的时间都不同。
func DummyMiddleware(c *gin.Context) {
	t := time.Now()
	fmt.Println("time is",t)

// Pass on to the next-in-chain

	c.Next()

}

func main() {

// Insert this middleware definition before any routes
	router := gin.Default()

	router.Use(DummyMiddleware)

	router.GET("/test", func(c *gin.Context) {
		log.Println("get /test")
		c.String(200, "test")
	})

	router.Run(":7777")

}