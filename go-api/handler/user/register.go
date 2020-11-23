package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// 数据库是否存在该用户
	//db.First("username")
	//userIsExist(username)


}
