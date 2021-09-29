package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "加密错误",
		})
		return
	}

	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashedPassword),
	}

	DB.Create(&newUser)

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func UserLogin(c *gin.Context) {
	DB := common.GetDB()
	// args
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// validate
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

	// 手机号码验证
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "用户不存在",
		})
		return
	}
	// 判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码错误",
			"errorMsg": err,
			"info": user,
			//"info2": user.Telephone,
		})
		return
	}

	// 发放token
	token := "111"
	c.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
		"data": gin.H{"code": token},
	})
	// 返回结果
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


