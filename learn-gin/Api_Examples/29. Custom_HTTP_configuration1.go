package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// Run函数其实就是对http.ListenAndServe封装了一层，并做了一些预处理
	//router.Run(":7777")
	http.ListenAndServe(":7777", router)
}
