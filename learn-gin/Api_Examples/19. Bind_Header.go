package main

import (
	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate int `header:"Rate" binding:"required"`
	Domain string `header:"Domain"`
}


func main() {
	router := gin.Default()
	router.GET("/header", func(c *gin.Context) {
		var header testHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"Rate": header.Rate,
			"Domain": header.Domain,
		})
	})
	router.Run(":7777")
}