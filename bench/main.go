package main

import (
	"bench/myredis"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/getkey", myredis.GetKey)
	router.POST("/setkey", myredis.SetKey)
	router.Run()
}
