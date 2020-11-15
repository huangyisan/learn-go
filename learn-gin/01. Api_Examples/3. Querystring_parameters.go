package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	//请求/welcome>firstname=h&lastname=yisan
	//第一个参数只需要填写uri地址即可
	route.GET("/welcome", func(c *gin.Context) {
		// 获取querystring中的key，如果没获取到，则使用guest替代。
		firstname := c.DefaultQuery("firstname","guest")
		lastname := c.Query("lastname")
		c.String(200, "your name is %s%s", firstname,lastname)
	})

	route.Run(":7777")
}