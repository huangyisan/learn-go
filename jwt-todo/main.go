package main

import (
	"github.com/gin-gonic/gin"
	. "jwt-todo/handler"
	. "jwt-todo/middleware"
	"log"
)

var (
	routers = gin.Default()
)

func main() {
	routers.POST("/login", Login)
	routers.POST("/todo", TokenAuthMiddleware(), CreateTodo)
	routers.POST("/logout", TokenAuthMiddleware(), Logout)
	routers.POST("/token/refresh", Refresh)
	log.Fatal(routers.Run(":8080"))
}
