package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ProfileForm struct {
	Name string `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/profile", func(c *gin.Context) {
		var form ProfileForm
		if err := c.ShouldBind(&form); err != nil {
			c.String(400,err.Error())
			return
		}
		err := c.SaveUploadedFile(form.Avatar, "./tmpdir/myAvatar")
		if err != nil {
			c.String(444,err.Error())
		}
		c.String(200,"ok")
	})
	r.Run(":7777")
}