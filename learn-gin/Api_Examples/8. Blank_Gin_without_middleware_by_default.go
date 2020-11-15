package main

import "github.com/gin-gonic/gin"

// 使用gin.New()，可以生成一个不带任何middleware的实例。
// 使用gin.Default(), 默认生成一个带Logger and Recovery的实例。

func main() {
	router := gin.New()
	router.Run(":7777")
}

