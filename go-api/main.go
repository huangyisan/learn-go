package main

import (
	"github.com/gin-gonic/gin"
	"goapi/router"
)

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run(":7777")

}