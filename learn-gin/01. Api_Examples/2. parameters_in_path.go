package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func printName(c *gin.Context) {
	// 指定name为param
	name := c.Param("name")
	// c.String方法返回字符串
	c.String(200, "hello %s", name)

}

func printAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	str := fmt.Sprintf("hello %s, do %s", name, action)
	fmt.Println(str)
	c.String(200, str)
}

func main() {
	route := gin.Default()

	//当且仅当请求路径为/user/xxxx的时候，才触发test方法，否则返回404
	route.GET("/user/:name",printName)

	// 请求路径可以存在action或者不存在action，如果不写action，则重定向到name
	route.GET("/user/:name/*action", printAction)

	route.POST("/user/:name/*action", func(c *gin.Context) {
		c.FullPath() == "/user/:name/*action" // true
	})



	route.Run(":7777")
}
