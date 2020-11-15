package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	route :=gin.Default()

	// 限定内存大小，这边我的理解是会先把文件内容读入到内存中。默认是32MB
	route.MaxMultipartMemory = 8 << 20 // 8388608 8MB

	route.POST("/upload", func(c *gin.Context) {
		// 如果使用postman，则指定form-data的key为file方式，选择对应的文件至value。
		// FormFIle里面的name参数表示form中的key，如果采用html方式编写了表单，那么这个name就是<input type="file" name="fname">
		file, _:= c.FormFile("fname")
		log.Println(file.Filename)

		// 将文件传输到服务端的路径
		c.SaveUploadedFile(file, fmt.Sprintf("./tmpdir/%s", file.Filename))
		c.String(200, "上传成功 %s", file.Filename)
	})

	route.Run(":7777")
}