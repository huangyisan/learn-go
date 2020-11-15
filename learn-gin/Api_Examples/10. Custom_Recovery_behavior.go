package main

// custom recovery 在当前v1.6.3版本还未加入
import "github.com/gin-gonic/gin"

func main() {
	// 定义一个不带中间件的gin
	r := gin.New()
	r.Use(gin.Logger())

	r.Run(":7777")
}
