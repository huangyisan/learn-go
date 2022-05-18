package main

import (
	"github.com/gin-gonic/gin"
	. "jwt-todo/handler"
	"log"
)

var (
	routers = gin.Default()
)

func main() {
	routers.POST("/login", Login)
	routers.POST("/todo", CreateTodo)
	log.Fatal(routers.Run(":8080"))
}
