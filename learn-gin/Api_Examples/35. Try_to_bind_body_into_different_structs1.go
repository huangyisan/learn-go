package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 采用ShouldBind则会无法重新使用body内的信息, 使用ShouldBindBodyWith可以将body内的信息先存储在context中，然后进行绑定。
// ShouldBindBodyWith会略微降低性能。

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// 该方法会将body内容先存放，所以可以一直被重用。

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 如下顺序匹配对于body内容的前后顺序无关，条件只会挨个匹配整个body的内容，如果命中了则执行。
	//{
	//    "bar":"bar22",
	//    "foo":"foo22"
	//}
	// 上面为body内容，formA匹配到foo,则直接输出， 而非因为bar是第一条，且和formB相关，所以优先打印出formB
	// This reads c.Request.Body and stores the result into the context.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// At this time, it reuses body stored in the context.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// And it can accepts other formats
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
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