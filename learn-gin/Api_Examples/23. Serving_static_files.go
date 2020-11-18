package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 第一个参数为请求前缀，比如请求http://127.0.0.1:7777/asstes/LICENSE 就能胡渠道LICENSE的信息了， LICENSE是存放在./tmpdir路径下的
	router.Static("/asstes", "./tmpdir")
	// staticFS类似Static，但其请求前缀能打印出对应目录下的所有文件，而static方式如果只请求前缀是不会打印出来的。该功能类似nginx的autoindex on。
	router.StaticFS("/more_static", http.Dir("./tmpdir"))
	// 给单个文件指定其文件相对路径。
	router.StaticFile("/favicon.ico", "./tmpdir/shinobi.ico")
	router.Run(":7777")

}