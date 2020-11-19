package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func main() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("./tmpdir/LICENSE")
	})

	//
	//var fs http.FileSystem = FileOpen("ss")
	// 可以将这个http.Dir理解为主目录
	var fs http.FileSystem = http.Dir("./tmpdir")// ...
	//fs = fo.Open("./dd")
	//var fs http.FileSystem = http.Dir("./tmpdir")// ...

	router.GET("/fs/file", func(c *gin.Context) {
		//基于http.Dir的目录中的某个文件地址。
		c.FileFromFS("LICENSE", fs)
	})
	router.Run(":7777")
}