package user

import (
	"github.com/gin-gonic/gin"
	. "go-api/handler"
	"go-api/model"
	"go-api/pkg/auth"
	"go-api/pkg/errno"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 从数据库中获取用户信息
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 对比密码
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}
}
