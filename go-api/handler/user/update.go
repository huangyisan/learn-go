package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
	. "go-api/handler"
	"go-api/model"
	"go-api/pkg/errno"
	"go-api/util"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))
	// binding the user data
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
