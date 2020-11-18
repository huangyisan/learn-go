package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

/*
XML: gin.Context.XML
YAML: gin.Context.YAML
JSON: gin.Context.JSON
ProtoBuf: gin.Context.ProtoBuf
 */

func main() {
	router := gin.Default()
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"type": "JSON",
		})
	})

	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"type": "YAML",
		})
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"type": "XML",
		})
	})

	router.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// pb写需要单独学习下
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	router.Run(":7777")
}