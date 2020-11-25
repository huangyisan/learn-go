package user

import (
	"github.com/gin-gonic/gin"
	"goapi/model"
	"goapi/model/user"
	"goapi/response"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")

	// 验证判断
	if len(password1) < 6 {
		response.Fail(c, "", "密码过短")
		return
	}

	if password1 != password2 {
		response.Fail(c, "", "两次密码输入不一致")
		return
	}

	// 获取db对象
	db := model.GetDB()

	//数据库是否存在该用户
	a := db.Where("username = ?", username).Take(&user.User{})
	if a.RowsAffected == 0 {
		// 加密password
		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
		if err != nil {
			response.Fail(c, "", "密码加密失败")
			return
		}
		newUser := user.User{
			Username: username,
			Password: string(encryptPassword),
		}
		db.Create(&newUser)
		response.Success(c, "", "创建用户成功")



	}else{
		response.Fail(c, "","已经存在该用户")
		return
	}

	//db.First("username")
	//userIsExist(username)
}
