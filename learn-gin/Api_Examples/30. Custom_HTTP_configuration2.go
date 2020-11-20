package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 对http服务自定义。一些超时时间，等参数
/*
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
ListenAndServe 方法其实是定义一个server，该server定义一些http的参数，然后最终返回server.ListenAndServe()
 */

func main() {
	router := gin.Default()
	// 所以可以自定义server
	server := &http.Server{
		Addr: ":7777",
		Handler: router,
		ReadTimeout: 10 * time.Second,
	}
	// 最后调用ListenAndServe()
	server.ListenAndServe()
}
