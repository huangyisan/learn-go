package main

import "github.com/gin-gonic/gin"

// 使用tag form进行绑定
// 将query内容自定义进行输出
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	// 嵌套了StructA
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	// 使用bind绑定
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.Run(":7777")
}

/*
127.0.0.1:7777/getb?field_a=hello&field_b=world
响应：
{
    "a": {
        "FieldA": "hello"
    },
    "b": "world"
}

对照B的结构体
type StructB struct {
	// 嵌套了StructA
	//type StructA struct {
	//	FieldA string `form:"field_a"`
	//}
	NestedStruct StructA
	FieldB string `form:"field_b"`
}
 */