package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("http2-push").Parse(`
    <html>
    <head>
      <title>Https Test</title>
      <script src="/assets/app.js"></script>
    </head>
    <body>
      <h1 style="color:red;">Welcome, Ginner!</h1>
    </body>
    </html>
    `))

func main() {
	r := gin.Default()
	// 设定静态文件路径， 让后面请求app.js能知道去哪里寻找该文件
	r.Static("/assets", "./tmpdir")
	// 渲染html到浏览器
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		// 实例化一个pusher对象
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			// 尝试将app.js文件push到客户端
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "http2-push", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in http://127.0.0.1:7777
	r.Run(":7777")
}

/*
当浏览器请求http://127.0.0.1:7777的时候，打开console控制台，能看到推送下来的hello world
 */