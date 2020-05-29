package main

import (
	fa "./工厂模式"
	"fmt"
)

// 如果在一个其他非main包里面， 写了一个Student的结构体，因为是大写的S， 所以可以在main包中引入后使用。但如果Student修改成了小写的s， 此时就无法
// 在main中使用了，因为是私有的结构体， 此时如果main函数还是想继续使用student结构体，此时就要用到工厂模式来解决。


func main() {
	// 因为student首字母小写没法使用
	//stu := fa.student{
	//	Name:"lilei",
	//}

	// 调工厂函数
	stu := fa.NewStudent("lilei", 22)
	fmt.Println(stu.GetAge())
}