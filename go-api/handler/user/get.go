package user

import (
	"github.com/gin-gonic/gin"
	. "go-api/handler"
	"go-api/model"
	"go-api/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
