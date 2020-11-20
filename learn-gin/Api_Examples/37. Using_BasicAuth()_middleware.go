package main

// basic认证是gin框架自己提供的
// 在访问的时候需要添加Authorization头
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	// Group上启用basicAuth认证中间件
	// Accounts左边为user, 右边为password, 设定了部分用户
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		// gin.AuthUserKey 为字符串常量"user"， 从context中获取到user
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println("user:",user)
		// 将得到的user信息去secrets中获取其对应的secret信息(email和phone，并不是密码)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Listen and serve on 0.0.0.0:7777
	r.Run(":7777")
}

/*
curl --location --request GET '127.0.0.1:7777/admin/secrets?user=%22yisan%22' \
--header 'Authorization: Basic Zm9vOmJhcg=='


Zm9vOmJhcg==
$ echo "Zm9vOmJhcg==" | base64.exe -d
foo:bar

$ echo -n "austin:1234" | base64
YXVzdGluOjEyMzQ=

如果尝试用manu:4321 则返回"NO SECRET"
 */
