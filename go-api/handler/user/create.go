package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	. "go-api/handler"
	"go-api/pkg/errno"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")

		//这里对err进行判断是否为user not found类型,方便如果上面username没填写触发的error 进行比对.
		if errno.IsErrUserNotFound(err) {
			log.Debug("err type is ErrUserNotFound")
		}

		SendResponse(c, err, nil)
		// 这里打印出了error 给到开发人员
		log.Errorf(err, "Get an error")
		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)

}
