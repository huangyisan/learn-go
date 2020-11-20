package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 采用ShouldBind则会无法重新使用body内的信息, 使用ShouldBindBodyWith可以将body内的信息先存储在context中，然后进行绑定。
// ShouldBindBodyWith会略微降低性能。
// 只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf使用ShouldBindBodyWith，对于其他格式，如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// 如果使用ShouldBind，则不会继续往下匹配，因为body不可以被重新消费。
// 所以，如果body格式固定的情况下，是可以用ShouldBind来处理的。
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}else {
		c.String(200,errA.Error())
	}
}

func main() {
	router := gin.Default()
	router.POST("/getbody", SomeHandler)
	router.Run(":7777")
}

/*
curl --location --request POST '127.0.0.1:7777/getbody' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "bar":"foo22",
    "foo":"foo33"
}'
 */