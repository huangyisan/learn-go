package main

import "fmt"

// map是引用类型, 遵守引用类型传递的机制, 在一个函数接受map, 修改后, 会直接修改原来的map
// map的容量达到后, 增加元素会自动扩容.
// map的value也经常使用struct类型



// 定义学生结构体
type student struct {
	Name string
	Age int
	Address string
}

func main() {
	stuInfo := make(map[int]student)

	// 实例化结构体
	stu1 := student{
		Name:"yisan",
		Age:22,
		Address:"china",
	}

	stu2 := student{
		Name:"yisan1",
		Age:32,
		Address:"cc",
	}

	// 将内容放入stuInfo中
	stuInfo[0] = stu1
	stuInfo[1] = stu2

	fmt.Println(stuInfo)
	fmt.Println(stuInfo[0].Age)

//	遍历年龄
for k,v := range stuInfo {
	fmt.Println("学生编号为", k)
	fmt.Println("学生的年龄", v.Age)
}
}