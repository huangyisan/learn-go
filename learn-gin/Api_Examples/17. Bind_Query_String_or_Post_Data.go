package main
// 同时绑定query以及json方式

import (
"log"
"time"

"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":7777")
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	// 如果是get请求，则直接绑定query， 如果是post, 则先验证content-type是否为json或者xml， 然后用form-data方式进行提交数据
	if err := c.ShouldBind(&person); err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}else {
		log.Println("err",err.Error())
	}

	c.String(200, "Success")
}

/*
Curl请求样例
Get:
 curl -X GET "localhost:7777/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"

Post:
curl --location --request POST '127.0.0.1:7777/testing' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"testname"
}'
 */