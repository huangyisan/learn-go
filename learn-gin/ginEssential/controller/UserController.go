package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {

	DB := common.GetDB()

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	// 参数校验
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "手机号码必须为11位",
		})
		return
	}

	if len(password) <6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "密码长度必须大于6位",
		})
		return
	}

	// 判断手机号码是否存在
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "手机已经被注册",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomName()
	}

	// 注册用户
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	}

	DB.Create(&newUser)

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	fmt.Println(user.ID)
	if user.ID != 0 {
		return true
	}
	return false
}


