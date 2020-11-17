package main

import "github.com/gin-gonic/gin"

// 绑定url，直接在tag上标记uri即可。

type Person struct {
	// tag中的uri表示绑定uri验证
	Id string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}


func main(){
	router := gin.Default()
	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		// 使用ShouldBindUri方法绑定uri
		if err:=c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{
				"error message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"name": person.Name,
			"id": person.Id,
		})
	})
	router.Run(":7777")
}

/*
curl --location --request GET '127.0.0.1:7777/yisan/987fbc97-4bed-5078-9f07-9141ba07c9f3'
 */

