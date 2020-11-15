package main

import "github.com/gin-gonic/gin"

// 在11. how to write log file已经使用DisableConsoleColor()方法关闭了日志在控制台颜色的显示。

// ForceConsoleColor()函数可以让颜色强制显示

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "Pong",
		})
	})

	router.Run(":7777")
}