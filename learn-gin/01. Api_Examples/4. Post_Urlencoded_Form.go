package main

import "github.com/gin-gonic/gin"

// Post 提交数据常见为x-www-form-urlencoded或者data-form方式
// x-www-form-urlencoded为post默认方式

func main()  {
	route := gin.Default()

	route.POST("/urlencoded_post", func(c *gin.Context) {
		// PostForm方法从urlencoded form 或者 multipart form获取对应key的value
		message := c.PostForm("message")
		// 从urlencoded form 或者 multipart form获取对应key的value, 没获取到则返回默认值
		nick := c.DefaultPostForm("nick", "anonymous")

		// 将获取到的form信息，以Json方式回写到前台页面。
		c.JSON(200, gin.H{
			"status": "post",
			"message": message,
			"nick": nick,
		})
	})

	route.Run(":7777")
}

// 使用curl模拟
/*
curl --location --request POST '127.0.0.1:7777/urlencoded_post' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'nick=huangdou' \
--data-urlencode 'message=go to school'


response:
{
    "message": "go to school",
    "nick": "huangdou",
    "status": "post"
}
 */