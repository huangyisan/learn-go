package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/model/user"
)

func Auth(c *gin.Context) {
	db := model.GetDB()
	username := c.PostForm("username")
	//password := c.PostForm("password")

	dbPassword, _ := db.Where("username = ?", username).Take(&user.User{}).Get("password")

	//dbPassword.Get("password")
	fmt.Println(dbPassword)

	c.JSON(200, gin.H{
		"message":"auth",
		"dbPassword":dbPassword,
	})
}
