package user

import (
	"github.com/gin-gonic/gin"
	. "go-api/handler"
	"go-api/pkg/errno"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	infos,count, err := 
}
