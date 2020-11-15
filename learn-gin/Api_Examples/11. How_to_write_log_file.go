package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 因为不再输出到控制台， 讲将输出写入文件，所以关闭控制台彩色输出
	gin.DisableConsoleColor()

	// 创建一个file
	f, _ := os.Create("./tmpdir/customer_log.txt")

	// 修改默认输出至文件中
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果要同时在屏幕以及文件中打印日志，可以如下，因为MultiWriter接受多个io对象
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 接下来照旧， 初始化实例后开始使用
	router := gin.Default()
	router.GET("/printFile", func(c *gin.Context) {
		c.String(200, "printFile")
	})

	router.Run(":7777")
}
