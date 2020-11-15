package main

import "github.com/gin-gonic/gin"

// gin使用go-playground/validator/v10对请求体进行验证。
// 当前支持JSON，XML，YAML以及标准的form方式(foo=bar&boo=baz)进行验证。
// 我们只需要对结构体中的属性使用``标记tag即可。比如对json的绑定写成 `json:filedName`

// Gin中提供了两种方式进行绑定
// Must Bind，使用该方法，如果验证失败，会直接返回400状态码，并直接抛出错误。
// Should Bind, 使用该方法， 如果验证失败， 失败可以被开发者捕获，然后进一步处理，比如自定义状态码或者自定义错误。


type Login struct {
	// 一些注意点
	// 属性首字母大写。 这边绑定了User和Password，两者必须为json的方式传入，且带了required，表示必须存在
	// Email的属性，json提交的时候为mail，所以提交过来时什么key，是和tag中定义的相关，和属性名称无关。
	// binding后面的属性如果有多个，用逗号分割，紧挨着写，不然会报错。 binding:"required,email" 而非 binding:"required, email"
	User string `json: "user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"mail" binding:"required,email"`
}

func main() {
	router := gin.Default()

	// 编写一个POST请求，提交user和password信息
	router.POST("/login", func(c *gin.Context) {
		// 定义json为Login类型
		var json Login
		// 绑定校验是否异常, 使用ShouldBindJSON方法绑定json类型
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400,gin.H{
				"err" : err.Error(),
				//异常则退出

			})
			return
		}
		c.JSON(200, gin.H{
			"message":"validate ok",
		})
	})
	router.Run(":7777")
}

// postman -> body -> raw
/*
{
    "user":"fakename",
    "password":"123123"
}
 */