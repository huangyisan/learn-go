package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		// 获取通过name获取对应的cookie
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			// 如果没有获取到，则set cookie
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "127.0.0.1", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run(":7777")
}

/*
通过浏览器控制台可以查看到cookie信息
 */