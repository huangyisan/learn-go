package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// 多文件form-data上传

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	router.POST("/upload", func(c *gin.Context) {
		// 返回多个form信息
		form, _:= c.MultipartForm()
		log.Printf("=====================%v\n",form)
		// 获取以upload为key name的file，作为map
		files := form.File["upload"]
		// 对上面的file map进行遍历，挨个调用SaveUploadFile方法
		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, fmt.Sprintf("./tmpdir/%s", file.Filename))

		}
		c.String(200, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":7777")
}