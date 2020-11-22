package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/router"
)

func main() {
	r := gin.Default()

	router.Router(r)

	if err := r.Run(":7777"); err != nil {
		panic(err.Error())
	}

}