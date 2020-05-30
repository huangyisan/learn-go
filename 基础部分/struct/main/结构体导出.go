package main
//结构体中，结构体名称首字母大写，则可以被导出。
//结构体属性，属性名词首字母大写，则可以被访问。
import (
	"../dump"
	"fmt"
)

func main()  {
	human := dump.Human{
		"yisan",
		29,
	}

	var student dump.Student
	student.Grade = "one"

	fmt.Println(human)
	fmt.Println(student)
}
