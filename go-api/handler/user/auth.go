package user

import (
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/model/user"
	"goapi/response"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c *gin.Context) {
	db := model.GetDB()


	username := c.PostForm("username")
	password := c.PostForm("password")

	var user user.User
	db.Where("username = ?", username).Take(&user)
	if user.ID == 0 {
		response.Fail(c,"","不存在的用户")
		return
	}

	//比较密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c,"","验证失败")
		return
	}

	// 对比密码成功
	response.Success(c, "","认证成功")

}
