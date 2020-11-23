package user

import (
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/model/user"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")

	// 验证判断
	if len(password1) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "密码过短",
		})
		return
	}

	if password1 != password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": "两次密码输入不一致",
		})
		return
	}

	// 获取db对象
	db := model.GetDB()
	var newUser user.User
	newUser = user.User{
		Username: username,
		Password: password1,
	}

	//数据库是否存在该用户
	a := db.Where("username = ?", username).Take(&user.User{})
	if a.RowsAffected == 0 {
		db.Create(&newUser)
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"message": "已经存在该用户",
		})
	}

	//db.First("username")
	//userIsExist(username)
}
