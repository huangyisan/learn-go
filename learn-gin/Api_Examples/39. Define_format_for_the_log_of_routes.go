package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 定义终端打印的日志
/*
默认形式：
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
 */
func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		// 修改打印的样式
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// Listen and Server in http://0.0.0.0:8080
	r.Run(":7777")
}

/*
修改后的终端日志输出
2020/11/21 00:19:28 endpoint POST /foo main.main.func2 3
2020/11/21 00:19:28 endpoint GET /bar main.main.func3 3
2020/11/21 00:19:28 endpoint GET /status main.main.func4 3
 */