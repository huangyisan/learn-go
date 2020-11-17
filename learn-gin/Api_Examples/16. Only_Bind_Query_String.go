package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Persion struct {
	Name string `form:"name"`
	Address string  `form:"address"`
}

func startPage(c *gin.Context) {
	var person Persion
	// ShouldBindQuery方法只绑定query
	if c.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		c.String(200, "success")
	}
}

func main() {
	router := gin.Default()
	// Any为任何请求方式
	router.Any("/person", startPage)
	router.Run(":7777")
}