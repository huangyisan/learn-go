package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	req := c.Request
	c.JSON(200, gin.H{
		"Method": req.Method,
		"URL":    req.URL,
	})
	fmt.Println("一个请求成功后的测试打印")
}

func main() {
	// 默认可以使用gin.Default()和gin.New()创建。区别在于gin.Default()也适用gin.New()创建engine实例，但是会默认使用Logger和Recover中间件。
	route := gin.Default()

	// 第一个参数为请求路径， 第二个参数是方法，其定义为type HandlerFunc func(*Context), 可以鼠标点击查看Context的结构体内容
	route.GET("/get", test)
	route.POST("/post", test)
	// Delete Put Head Option Patch都类似

	panic(route.Run(":7777"))

}
